//var newWindowObi=window.open("出错信息");
//newWindowObi.document.write(json);

var gData = null;
var gOptionTpl = null;
var gConfirmDialogSubmitBtnAction = "";
var gPluginKey = "";

$(function() {
    // style
    $('.switch[type="checkbox"]').bootstrapSwitch();
    $('.selectpicker').selectpicker();

    // btnclick
    gOptionTpl = returnArgOptionTpl();
    $("#args_add_btn").click(argsAdd);
    $("#bookmark_use_btn").click(handleBookmarkUse);
    $("#bookmark_add_btn").click(bookmarkAdd);
    $("#bookmark_edit_btn").click(bookmarkEdit);
    $("#bookmark_drop_btn").click(bookmarkDrop);
    $("#bookamrkAddBtn").click(handleBookamrkAdd)
    $("#bookmark_add_input").focus(hiddenErrAlert);
    $("#confirm_dialog_submit_btn").click(clickConfirmDialogSubmitBtn)
    $("#plugin_use_btn").click(pluginUse);
    $("#submit_btn").click(handleSubmit);

    // render data
    gData = returnConfigData();
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

function returnConfigData() {
    var text = $("#config_json").html();
    var json = $.parseJSON(text);
    return json;
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

function argsRemove(btn) {
    $(btn).parent().parent().remove();
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
		},
		'error':	function (XMLHttpRequest, textStatus, errorThrown) {
            $(btn).button('reset');
            alertMsg(e);
		}
	});

    return false;
}

var g_num = 0;
var submit_url = '/submit';

var arg_tpl = ' <tr class="tr-args"> '+
'<td> <label>参数名：</label> <input type="text" class="form-control arg-keys" style="display: inline; width: 12em;" /> </td> '+
'<td> <label>参数值：</label> <input type="text" class="form-control arg-values" style="display: inline; width: 12em;" /> </td> '+
'<td> <label for="method">参数类型：</label> <input type="radio" name="arg_method_%NUM%" id="arg_method_get_%NUM%" value="get" /><label for="arg_method_get_%NUM%">GET</label>&nbsp; '+
'<input type="radio" name="arg_method_%NUM%" id="arg_method_post_%NUM%" class="arg-method-post" value="post" checked="checked" /><label for="arg_method_post_%NUM%">POST</label> </td> <td> '+
'<button type="button" class="btn btn-default" onclick="removeArg(this)">删除参数</button> </td> </tr> ';

var result_tpl = ' <p><span class="text-danger">执行时间：</span><span class="text-success"> %TIMES% 秒</span></p>'+
'<p><span class="text-danger">状态码：</span><span class="text-success"> %STATUS% </span></p>'+
'<p><span class="text-danger">返回值如下：</span></p>'+
'<hr />'+
'<div id="dataContainer"></div> '+
'<hr />'+
'<p class="text-danger">压测数据：</p>'+
'<pre>%BOOM%</pre>';

var err_tpl = ' <h4 class="text-danger">请求出错！</h4><hr /> <p><span class="text-warning">原因：</span> <span class="text-muted">%MSG%</span> </p> ';

$(function() {

    $("#addArgBtn").click(function() {
        var tpl = arg_tpl.replace(/%NUM%/g, g_num);
        $("#table_args").append(tpl);
        g_num++;
    });

});


function onSubmit() {
    var data = {
        "url":      $.trim($("input[name=url]").val()),
        "n":		parseInt($("#boom_n").val()),
        "c":		parseInt($("#boom_c").val()),
        "key":      $.trim($("#api_sign_key").val()),
        "secret":   $.trim($("#api_secret").val()),
        "args":     []
    };

    if (data.url == "") {
        alert("请指定URL地址");
        return false;
    }

	if ($("#method_post").is(":checked")) {
		data.method = "POST";
	} else {
		data.method = "GET";
	}

    $("tr.tr-args").each(function() {
        var key = $(this).find(".arg-keys").val();
        if (key == "") {
            return false;
        }
        var value = $(this).find(".arg-values").val();
        var method = "";
        if ($(this).find(".arg-method-post").is(":checked")) {
            method = "post";
        } else {
            method = "get";
        }
        data.args.push({
            "key":      key,
            "value":    value,
            "method":   method
        });
    });

     console.log("reqeustData", data);

     $("#resultPanel").html("Please wait...");

	$.ajax(submit_url, {
		'data': JSON.stringify(data),
		'type': 'POST',
		'processData': false,
		'contentType': 'application/json',
		'dataType':	'json',
		'success':	function(response) {
			console.log("resultData", response);

			var tpl = '';

			if (response.Status != 200) {
				tpl = err_tpl.replace("%MSG%", response.Msg);
				$("#resultPanel").html(tpl);
				return;
			}

			tpl = result_tpl.replace("%STATUS%", response.Data.Status);
			tpl = tpl.replace("%TIMES%", response.Data.Times);
			tpl = tpl.replace("%BOOM%", response.Data.Boom);
			$("#resultPanel").html(tpl);

			var options = {
				dom : '#dataContainer' //对应容器的css选择器
			};
			var jf = new JsonFormater(options); //创建对象
			jf.doFormat(response.Data.Data); //格式化json

		},
		'error':	function (XMLHttpRequest, textStatus, errorThrown) {
			console.log("errorData", errorThrown);
			alert(textStatus);
		}
	});

     return false;
}

