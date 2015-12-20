var gData = null;
var gOptionTpl = null;
var gConfirmDialogSubmitBtnAction = "";
var gPluginKey = "";
var gResultTest = "";

var configs = {
    "initDataUrl": "/?act=initData"
};

var utils = {
    "tplCompile": function(id) {
        var html = $("#" + id).html();
        return Handlebars.compile(html);
    }
};

var templates = {
    "argsOptions":   utils.tplCompile("args_tpl"),
    "pluginOptions": utils.tplCompile("plugin_option_tpl"),
    "pluginPanel": utils.tplCompile("plugin_panel_tpl")
};

var page = {
    "renderBookmark": function(bookmark) {
        var data = bookmark.Data;

        $('#method').bootstrapSwitch('state', data.Method=="GET");
        $('#url').val(data.Url);

        this.renderArgs(data.Args, true);

        $('#bm_switch').bootstrapSwitch('state', data.Bm.Switch);
        $('#bm_n').val(data.Bm.N);
        $('#bm_c').val(data.Bm.C);
    },

    "renderBookmarks": function(bookmarks) {
    },

    "renderPlugins": function(plugins, pluginKey) {
        var html = templates.pluginOptions({"Plugins": plugins});
        $("#plugin").html(html);
        $('#plugin').selectpicker("refresh");

        var html0 = templates.pluginPanel(plugins[pluginKey]);
        $("#plugin_panel").html(html0);
    },

    "renderArgs": function(args, isReset) {
        var html = templates.argsOptions({"Args": args});
        if (isReset) {
            return $("#args_body").html(html);
        }
        return $("#args_body").append(html);
    },

    "refresh": function() {
        $('.switch[type="checkbox"]').bootstrapSwitch();
        $('.selectpicker').selectpicker();
    },

    "message": function (msg) {
        $("#alerter_content").html(msg);
        $("#alerter").modal('show');
    }
};

var args = {
    "add": function() {
        page.renderArgs([{
            "Key":    "",
            "Value":  "",
            "Method": "GET"
        }], false);

        page.refresh();
    },

    "remove": function (btn) {
        $(btn).parent().parent().remove();
    }
};

// window.onload
$(function() {
    // init libaray
    Handlebars.registerHelper('eq', function(v1, v2, options) {
        if(v1 == v2) {
            return options.fn(this);
        }
        return options.inverse(this);
    });

    // set style
    page.refresh();

    // 【回调地狱】获取初始化数据
    $.ajax(configs.initDataUrl, {
        "type": "GET",
        // "async": false,
        "dataType": "json",
        "success": function(respData, textStatus, jqXHR) {
            if (respData.Status != 200) {
                page.message(respData.Message);
                throw "init error";
            }

            console.log(respData);
            page.renderBookmark(respData.Data.Bookmark);
            page.renderPlugins(respData.Data.Plugins,
                               respData.Data.Bookmark.Data.Plugin.Key);
            page.refresh();
        },
        "error":  function(XMLHttpRequest, textStatus, errorThrown) {
            page.message("Request error: " + textStatus);
            throw "init error";
        }
    });

    return;

    // btnclick
    gOptionTpl = returnArgOptionTpl();
    $("#args_add_btn").click(argsAdd);
    $("#bookmark_use_btn").click(handleBookmarkUse);
    $("#bookmark_add_btn").click(bookmarkAdd);
    $("#bookmark_edit_btn").click(bookmarkEdit);
    $("#bookmark_drop_btn").click(bookmarkDrop);
    $("#bookamrkAddBtn").click(handleBookamrkAdd);
    $("#bookmark_add_input").focus(hiddenErrAlert);
    $("#confirm_dialog_submit_btn").click(clickConfirmDialogSubmitBtn);
    $("#plugin_use_btn").click(pluginUse);
    $("#submit_btn").click(handleSubmit);

    // render data
    gData = getInitData();
    gPluginKey = gData.bookmarks[gData.selected].plugin.key;
    renderData(gData);
});

function pluginUse() {
    $.get(g.pluginUrl, {}, function(data) {
        var key = $("#plugin").selectpicker('val');
        var plugin = data.data.plugins[key];
        renderPluginPanel(plugin);
        gPluginKey = key;
    }, "json");
}

function handleBookmarkDelete() {
    var key = $("#bookmark").selectpicker('val');
    var inputs = {"key": key};

    // 安全禁用
    var btn = this;
    $(btn).button('loading');

    // 联网
    $.ajax({
        "url":  g.bookmarkUrl + "?key=" + key,
        "type":  "DELETE",
        "data":  JSON.stringify(bookmark),
        "contentType":  "application/json",
        "cache": false,
        "dataType": "json",
        "success":  function(data){
            // 恢复按钮状态
            $(btn).button('reset');
            $("#confirm_dialog").modal('hide');

            if (data.status != 200) {
                alertMsg(data.msg);
                return false;
            }

            // 成功
            $("#bookmark option[value="+key+"]").remove();
            $("#bookmark").selectpicker('refresh');
        },
       "error":  function(XMLHttpRequest, textStatus, errorThrown) {
            // 恢复按钮状态
            $(btn).button('reset');
            $("#confirm_dialog").modal('hide');

            alertMsg(textStatus);
        }
    });
}

function handleBookmarkUse() {
    var btn = this;
    $(btn).button('loading');

    var key = $("#bookmark").selectpicker('val');
    var inputs = {"key": key};
    $.get(g.bookmarkUrl, inputs, function(data) {
        $(btn).button('reset');

        if (data.status != 200) {
            alertMsg(data.msg);
            return;
        }

        renderData(data.data);

    }, "json");
}

function renderData(data) {
    renderBookmarkOptions(data);
    renderPluginOptions(data);

    var bookmark = data.bookmarks[data.selected];
    renderContent(bookmark);
    renderArgsOption(bookmark);

    var plugin = data.plugins[bookmark.plugin.key];
    renderPluginPanel(plugin);
    renderPlugin(bookmark, plugin);
}

function handleBookamrkAdd() {
    var name = $.trim($("#bookmark_add_input").val());

    // 校验
    var reg = /^\S{1,20}$/;
    if (!reg.test(name)) {
        $("#bookmark_add_err").html("书签名称必须为1~20位非空字符");
        $("#bookmark_add_err").removeClass("hidden");
        return false;
    }

    try {
        var bookmark = getRequestData();
        bookmark.name = name;
    } catch(e) {
        $("#bookmark_add_err").html(e);
        $("#bookmark_add_err").removeClass("hidden");
        return false;
    }

    // 安全禁用
    var btn = this;
    $(btn).button('loading');

    // 联网
    $.ajax({
        "url":  g.bookmarkUrl,
        "type":  "POST",
        "data":  JSON.stringify(bookmark),
        "contentType":  "application/json",
        "cache": false,
        "dataType": "json",
        "success":  function(data){
            // 恢复按钮状态
            $(btn).button('reset');

            if (data.status != 200) {
                $("#bookmark_add_err").html(data.msg);
                $("#bookmark_add_err").removeClass("hidden");
                return false;
            }

            // 成功
            var option = {
                "bookmarks": {}
            };
            option.bookmarks[data.data.insertKey] = {"name": data.data.insertName};

            $('#add_dialog').modal('hide');
            var tpl = $("#bookmark_option_tpl").html();
            var html = juicer(tpl, option);
            $('#bookmark').append(html);

            $("#bookmark").selectpicker('val', data.data.insertKey);
            $('#bookmark').selectpicker("refresh");

        },
       "error":  function(XMLHttpRequest, textStatus, errorThrown) {
            // 恢复按钮状态
            $(btn).button('reset');

            $("#bookmark_add_err").html(textStatus);
            $("#bookmark_add_err").removeClass("hidden");
        }
    });
}

function handleBookmarkEdit() {
    try {
        var bookmark = getRequestData();
        var key = $("#bookmark").selectpicker('val');
        var name = $("#bookmark option[value="+key+"]").html();
        bookmark.name = name;
    } catch(e) {
        $("#confirm_dialog").modal('hide');
        alertMsg(e);
        return false;
    }

    // 安全禁用
    var btn = this;
    $(btn).button('loading');

    // 联网
    $.ajax({
        "url":  g.bookmarkUrl,
        "type":  "PUT",
        "data":  JSON.stringify(bookmark),
        "contentType":  "application/json",
        "cache": false,
        "dataType": "json",
        "success":  function(data){
            // 恢复按钮状态
            $(btn).button('reset');
            $("#confirm_dialog").modal('hide');

            if (data.status != 200) {
                alertMsg(data.msg);
                return false;
            }

            // 成功，无动作
        },
       "error":  function(XMLHttpRequest, textStatus, errorThrown) {
            // 恢复按钮状态
            $(btn).button('reset');
            $("#confirm_dialog").modal('hide');

            alertMsg(textStatus);
        }
    });
}

function hiddenErrAlert() {
    $("#bookmark_add_err").addClass("hidden");
}

function renderPlugin(bookmark, plugin) {
    for (var i in plugin.fields) {
        $("#plugin_"+i).val(bookmark.plugin.data[i]);
    }
}

function renderPluginPanel(plugin) {
    var tpl = $("#plugin_panel_tpl").html();
    var html = juicer(tpl, plugin);
    $("#plugin_panel").html(html);
}

function renderArgsOption(bookmark) {
    var tpl = $("#args_tpl").html();
    var html = juicer(tpl, bookmark);
    $("#args_body").html(html);
    $('.switch[type="checkbox"]').bootstrapSwitch();
}

function renderContent(bookmark) {
    $('#method').bootstrapSwitch('state', bookmark.method=="GET");
    $('#url').val(bookmark.url);

    $('#bm_switch').bootstrapSwitch('state', bookmark.bm.switch);
    $('#bm_n').val(bookmark.bm.n);
    $('#bm_c').val(bookmark.bm.c);
}

function renderPluginOptions(data) {
    var tpl = $("#plugin_option_tpl").html();
    var html = juicer(tpl, data);
    $("#plugin").html(html);
    $("#plugin").selectpicker('val', data.bookmarks[data.selected].plugin.key);
    $('#plugin').selectpicker("refresh");
}

function renderBookmarkOptions(data) {
    var tpl = $("#bookmark_option_tpl").html();
    var html = juicer(tpl, data);
    $("#bookmark").html(html);
    $("#bookmark").selectpicker('val', data.selected);
    $('#bookmark').selectpicker("refresh");
}

function bookmarkEdit() {
    $("#confirm_dialog_title").html("编辑书签");
    $("#confirm_dialog_text").html("您确定要将当前内容替换到选定书签吗？");
    gConfirmDialogSubmitBtnAction = "edit";
    $("#confirm_dialog").modal("show");
}

function bookmarkDrop() {
    $("#confirm_dialog_title").html("删除书签");
    $("#confirm_dialog_text").html("您确定删除选定书签吗？");
    gConfirmDialogSubmitBtnAction = "drop";
    $("#confirm_dialog").modal("show");
}

function clickConfirmDialogSubmitBtn () {
    switch (gConfirmDialogSubmitBtnAction) {
    case 'edit':
        handleBookmarkEdit();
        break;

    case 'drop':
        handleBookmarkDelete();
        break;
    }
}

function bookmarkAdd() {
    try {
        getRequestData();
    } catch(e) {
        alertMsg(e);
        return;
    }
    $("#bookmark_add_err").html("");
    $("#bookmark_add_err").addClass("hidden");
    $("#bookmark_add_input").val("");
    $("#add_dialog").modal("show");
}

function alertMsg(msg) {
    $("#alerter_content").html(msg);
    $("#alerter").modal('show');
}

function argsAdd() {
    $("#args_body").append(gOptionTpl);
    $('.switch[type="checkbox"]').bootstrapSwitch();
}


function returnArgOptionTpl() {
    var tpl = $("#args_tpl").html();
    var empty = {"args":[{
        "key":    "",
        "value":  "",
        "method": "GET",
    }]};
    return juicer(tpl, empty);
}

function getRequestData() {
    var data ={};
    data.bm = {};

    // 获取数据
    data.method = $("#method").bootstrapSwitch("state") ? "GET" : "POST";
    data.url = $.trim($("#url").val());
    data.bm.switch = $("#bm_switch").bootstrapSwitch("state");
    data.bm.n =  parseInt($.trim($("#bm_n").val()));
    data.bm.c =  parseInt($.trim($("#bm_c").val()));

    data.args = [];
    $("#args_body tr").each(function() {
        var key = $.trim($(this).find(".arg-key").val());
        if (key == "") {
            return false;
        }
        data.args.push({
            "key":     key,
            "value":   $.trim($(this).find(".arg-value").val()),
            "method":  $(this).find(".arg-method").bootstrapSwitch("state") ? "GET" : "POST"
        });
    });

    data.plugin = {};
    data.plugin.data = {};
    data.plugin.key = gPluginKey;
    for (var i in gData.plugins[gPluginKey].fields) {
        data.plugin.data[i] = $.trim($("#plugin_"+i).val());
    }

    // 校验
    if (isNaN(data.bm.n) || isNaN(data.bm.c)) {
        throw '压测数据必须是数字';
    }
    if (data.bm.n <= 0 || data.bm.c <= 0) {
        throw '压测数据必须是正整数';
    }

    return data;
}

function handleSubmit() {
    var data = null;
    try {
        data = getRequestData();
    } catch(e) {
        alertMsg(e);
        return;
    }

    var btn = this;
    $(btn).button('loading');

	$.ajax({
        'url': g.submitUrl,
		'data': JSON.stringify(data),
		'type': 'POST',
		'processData': false,
		'contentType': 'application/json',
		'dataType':	'json',
		'success':	function(data) {
            $(btn).button('reset');

            if (data.status != 200) {
                alertMsg(data.msg);
                return false;
            }

            // 成功
            renderResult(data.data);
		},
		'error':	function (XMLHttpRequest, textStatus, errorThrown) {
            $(btn).button('reset');
            alertMsg(e);
		}
	});

    return false;
}

function renderResult(result) {
    gResultTest = result.test;

    var tpl = $("#result_tpl").html();
    var html = juicer(tpl, result);
    $("#result_panel").html(html);

    try {
        var div = $('<div></div>');
        $("#result_test_panel").append(div);
        var options = {"dom" : div};
        var jf = new JsonFormater(options); //创建对象
        jf.doFormat(gResultTest);    //格式化json

    } catch(e) {
        var iFrame = $('<iframe style="width: 100%; min-height: 350px;"></iframe>');
        $("#result_test_panel").append(iFrame);
        var iFrameDoc = iFrame[0].contentDocument || iFrame[0].contentWindow.document;
        iFrameDoc.write(gResultTest);
        iFrameDoc.close();
    }

    $("#result_new_tab_btn").click(function() {
        var newWindowObi=window.open("在新标签中浏览");
        newWindowObi.document.write(gResultTest);
    });
}
