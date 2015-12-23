var configs = {
    "initDataUrl":  "/?act=initData",
    "bookmarkUrl":  "/bookmark",
    "bookmarksUrl": "/bookmarks"
};

// TODO Move all global variables to this global object
var global = {
    "plugins": {}
};

var utils = {
    "tplCompile": function(id) {
        var html = $("#" + id).html();
        return Handlebars.compile(html);
    },

    "ajax": function(url, method, requestData, successCallback, errorCallback) {
        $.ajax({
            "url":  url,
            "type":  method,
            "data":  JSON.stringify(requestData),
            "contentType":  "application/json",
            "cache": false,
            "dataType": "json",
            "success":  function(data){
                successCallback(data);
            },
            "error":  function(XMLHttpRequest, textStatus, errorThrown) {
                errorCallback(textStatus);
            }
        });
    }
};

var templates = {
    "argsOptions":     utils.tplCompile("args_tpl"),
    "pluginOptions":   utils.tplCompile("plugin_option_tpl"),
    "pluginPanel":     utils.tplCompile("plugin_panel_tpl"),
    "bookmarkOptions": utils.tplCompile("bookmark_option_tpl")
};

var page = {
    "renderData": function(data) {
        // url
        $('#method').bootstrapSwitch('state', data.Method=="GET");
        $('#url').val(data.Url);

        // args
        this.renderArgs(data.Args, true);

        // bm
        $('#bm_switch').bootstrapSwitch('state', data.Bm.Switch);
        $('#bm_n').val(data.Bm.N);
        $('#bm_c').val(data.Bm.C);

        // plugin
        this.renderPlugin(data.Plugin);
    },

    "renderBookmarks": function(bookmarks, bookmarkName) {
        var html = templates.bookmarkOptions({"Bookmarks": bookmarks});
        $("#bookmark").html(html);
        $("#bookmark").selectpicker('val', bookmarkName);
        $('#bookmark').selectpicker("refresh");
    },

    "renderPlugin": function(plugin) {
        var html = templates.pluginPanel(global.plugins[plugin.Key]);
        $("#plugin_panel").html(html);
        $("#plugin").selectpicker('val', plugin.Key);

        // init plugins value
        var pluginData = plugin.Data;
        for (var i in pluginData) {
            $("#plugin_" + i).val(pluginData[i]);
        }
    },

    "renderPlugins": function(plugins, pluginKey) {
        var html = templates.pluginOptions({"Plugins": plugins});
        $("#plugin").html(html);
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
        $('#plugin').selectpicker("refresh");
    },

    "message": function (msg) {
        $("#alerter_content").html(msg);
        $("#alerter").modal('show');
    },

    "inputDialoyMessage": function(msg) {
        $("#bookmark_add_err").html(msg);
        $("#bookmark_add_err").removeClass("hidden");
        return false;
    },

    "inputDialoyMessageHide": function() {
        $("#bookmark_add_err").addClass("hidden");
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

var data = {
    "get": function() {
        var data = {};

        // 获取数据
        data.Method = $("#method").bootstrapSwitch("state") ? "GET" : "POST";
        data.Url = $.trim($("#url").val());
        data.Bm = {};
        data.Bm.Switch = $("#bm_switch").bootstrapSwitch("state");
        data.Bm.N =  parseInt($.trim($("#bm_n").val()));
        data.Bm.C =  parseInt($.trim($("#bm_c").val()));

        data.Args = [];
        $("#args_body tr").each(function() {
            var key = $.trim($(this).find(".arg-key").val());
            if (key == "") {
                return false;
            }
            data.Args.push({
                "Key":     key,
                "Value":   $.trim($(this).find(".arg-value").val()),
                "Method":  $(this).find(".arg-method").bootstrapSwitch("state") ? "GET" : "POST"
            });
        });

        data.Plugin = {};
        data.Plugin.Data = {};
        data.Plugin.Key = $("#plugin").val();
        for (var i in global.plugins[data.Plugin.Key].FieldNames) {
            data.Plugin.Data[i] = $.trim($("#plugin_"+i).val());
        }

        // 校验
        if (isNaN(data.Bm.N) || isNaN(data.Bm.C)) {
            throw "压测数据必须是数字";
        }
        if (data.Bm.N <= 0 || data.Bm.C <= 0) {
            throw "压测数据必须是正整数";
        }

        return data;
    }

};

var plugins = {
    "use": function() {
        var key = $("#plugin").selectpicker('val');
        var plugin = global.plugins[key];
        page.renderPlugin(plugin);
    }
};

var bookmarks = {
    "handleUse": function() { // TODO
        var btn = this;
        $(btn).button('loading');
        var name = $("#bookmark").selectpicker('val');

        utils.ajax(configs.bookmarkUrl, "POST", {"Name":name}, function(data) {
            $(btn).button('reset');

            console.log(data);

            if (data.Status != 200) {
                return page.inputDialoyMessage(data.Message);
            }

            page.renderData(data.Data);

        }, function(textStatus) {
            $(btn).button('reset'); // 恢复按钮状态
            return page.inputDialoyMessage(textStatus);
        });
    },

    "add": function() {
        // validate data
        try {
            data.get();
        } catch (e) {
            return page.message(e);
        }

        $("#bookmark_add_err").html("");
        $("#bookmark_add_err").addClass("hidden");
        $("#bookmark_add_input").val("");
        $("#add_dialog").modal("show");

        return true;
    },

    "handleAdd": function() {
        var name = $.trim($("#bookmark_add_input").val());

        var bookmark = {};
        try {
            bookmark.Data = data.get();
            bookmark.Name = name;
        } catch(e) {
            return page.inputDialoyMessage(e);
        }

        // 安全禁用
        var btn = this;
        $(btn).button('loading');

        console.log("insert bookmark", bookmark);

        utils.ajax(configs.bookmarksUrl, "POST", bookmark, function(data) {
            $(btn).button('reset'); // 恢复按钮状态

            if (data.Status != 200) {
                return page.inputDialoyMessage(data.Message);
            }

            // 成功[
            var renderData = {"Bookmarks":[bookmark.Name]};
            var html = templates.bookmarkOptions(renderData);
            $('#bookmark').append(html);
            $('#add_dialog').modal('hide');
            $("#bookmark").selectpicker('val', bookmark.Name);
            $('#bookmark').selectpicker("refresh");

        }, function(textStatus) {
            $(btn).button('reset'); // 恢复按钮状态
            return page.inputDialoyMessage(textStatus);
        });
    },

    "edit": function() {
        $("#confirm_dialog_title").html("编辑书签");
        $("#confirm_dialog_text").html("您确定要将当前内容替换到选定书签吗？");
        $("#confirm_dialog_submit_btn").click(this.handlEdit);
        $("#confirm_dialog").modal("show");
    },

    "handleEdit": function() {
        
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

    // 【回调地狱】获取初始化数据 initData
    return utils.ajax(configs.initDataUrl, "GET", {}, function(respData) {
        if (respData.Status != 200) {
            page.message(respData.Message);
            throw "init error";
        }

        console.log("initData:", respData);

        var data = respData.Data;
        global.plugins = data.Plugins;

        page.renderBookmarks(data.Bookmarks, data.Bookmark.Name);
        page.renderPlugins(data.Plugins, data.Bookmark.Data.Plugin.Key);
        page.renderData(data.Bookmark.Data);
        page.refresh();

        // event binding
        $("#bookmark_use_btn").click(bookmarks.handleUse);
        $("#bookmark_add_btn").click(bookmarks.add);
        $("#bookmark_add_input").focus(page.inputDialoyMessageHide);
        $("#bookamrkAddBtn").click(bookmarks.handleAdd);
        $("#bookmark_edit_btn").click(bookmarks.edit);
        $("#plugin_use_btn").click(plugins.use);

    }, function(textStatus) {
        page.message("Request error: " + textStatus);
        throw "init error";
    });

    // btnclick
    gOptionTpl = returnArgOptionTpl();
    $("#args_add_btn").click(argsAdd);

    $("#bookmark_drop_btn").click(bookmarkDrop);
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
        renderPlugin(plugin);
        gPluginKey = key;
    }, "json");
}

function handleBookmarkDelete() {
    var key = $("#bookmark").selectpicker("val");
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

function renderData(data) {
    renderBookmarkOptions(data);
    renderPluginOptions(data);

    var bookmark = data.bookmarks[data.selected];
    renderContent(bookmark);
    renderArgsOption(bookmark);

    var plugin = data.plugins[bookmark.plugin.key];
    renderPlugin(plugin);
    renderPlugin(bookmark, plugin);
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

function renderPlugin(plugin) {
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

// function handleSubmit() {
//     var data = null;
//     try {
//         data = getRequestData();
//     } catch(e) {
//         alertMsg(e);
//         return;
//     }

//     var btn = this;
//     $(btn).button('loading');

// 	$.ajax({
//         'url': g.submitUrl,
// 		'data': JSON.stringify(data),
// 		'type': 'POST',
// 		'processData': false,
// 		'contentType': 'application/json',
// 		'dataType':	'json',
// 		'success':	function(data) {
//             $(btn).button('reset');

//             if (data.status != 200) {
//                 alertMsg(data.msg);
//                 return false;
//             }

//             // 成功
//             renderResult(data.data);
// 		},
// 		'error':	function (XMLHttpRequest, textStatus, errorThrown) {
//             $(btn).button('reset');
//             alertMsg(e);
// 		}
// 	});

//     return false;
// }

// function renderResult(result) {
//     gResultTest = result.test;

//     var tpl = $("#result_tpl").html();
//     var html = juicer(tpl, result);
//     $("#result_panel").html(html);

//     try {
//         var div = $('<div></div>');
//         $("#result_test_panel").append(div);
//         var options = {"dom" : div};
//         var jf = new JsonFormater(options); //创建对象
//         jf.doFormat(gResultTest);    //格式化json

//     } catch(e) {
//         var iFrame = $('<iframe style="width: 100%; min-height: 350px;"></iframe>');
//         $("#result_test_panel").append(iFrame);
//         var iFrameDoc = iFrame[0].contentDocument || iFrame[0].contentWindow.document;
//         iFrameDoc.write(gResultTest);
//         iFrameDoc.close();
//     }

//     $("#result_new_tab_btn").click(function() {
//         var newWindowObi=window.open("在新标签中浏览");
//         newWindowObi.document.write(gResultTest);
//     });
// }
