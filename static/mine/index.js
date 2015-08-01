//var newWindowObi=window.open("出错信息");
//newWindowObi.document.write(json);

$(function() {
    // style
    $('.switch[type="checkbox"]').bootstrapSwitch();
    $('.selectpicker').selectpicker();

    // btnclick
    $("#args_add_btn").click(argsAdd);

    // render data
    var data = returnConfigData();
    renderBookmarkOptions(data);
});


function renderBookmarkOptions(data) {
    var tpl = $("#bookmark_option_tpl").html();
    var html = juicer(tpl, data);
    $("#bookmark").html(html);
    $('#bookmark').selectpicker("refresh");
}

function returnConfigData() {
    var text = $("#config_json").html();
    var json = $.parseJSON(text);
    return json;
}

function argsAdd() {
    var tpl = $("#args_tpl").html();
    $("#args_body").append(tpl);
    $('.switch[type="checkbox"]').bootstrapSwitch();
}

function argsRemove(btn) {
    $(btn).parent().parent().remove();
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

