$(function () {
    captcha();
    $('#imgcode').on('click', function () {
        captcha();
    });
})

// 获取验证码
function captcha() {
    var url = ctx + "captchaImage?s=" + Math.random();
    $.ajax({
        type: "get",
        url: url,
        success: function (r) {
            if (r.code == "0") {
                $("#imgcode").attr("src", r.data);
                $("#idkey").val(r.idkey);
            }
        }
    });
}

// 检查登陆
$("#checkLogin").on("click", function checkLogin() {
    if (!check()) {
        return
    }
    // 发送请求
    $.ajax({
        type: 'post',
        url: '/checklogin',
        data: $("#loginForm").serialize(),
        dataType: 'json',
        cache: false,
        success: function (result) {
            resultData(result);
        }
    });
})

// 处理返回数据
function resultData(result) {
    switch (result.code) {
        case 0: // 成功
            saveCookie();
            window.location.href = "/index";
            break;
        case -1: // 验证码错误
            $("#code").tips({
                side: 1,
                msg: result.msg,
                bg: '#FF5080',
                time: 3
            });
            $("#code").focus();
            break;
        default:
            $("#loginName").tips({
                side: 1,
                msg: result.msg,
                bg: '#FF5080',
                time: 15
            });
            $("#loginName").focus();
    }
}

// 检查参数
function check() {
    if ($("#loginName").val() === "") {
        $("#loginName").tips({
            side: 1,
            msg: '用户名不得为空',
            bg: '#AE81FF',
            time: 3
        });
        $("#loginName").focus();
        return false;
    } else {
        $("#loginName").val(jQuery.trim($("#loginName").val()));
    }

    let password = $("#password").val();
    if (password === "") {
        $("#password").tips({
            side: 1,
            msg: '密码不得为空',
            bg: '#AE81FF',
            time: 3
        });
        $("#password").focus();
        return false;
    }

    // 验证cookie中密码
    calcMd5(password);

    if ($("#code").val() === "") {
        $("#code").tips({
            side: 1,
            msg: '验证码不得为空',
            bg: '#AE81FF',
            time: 3
        });
        $("#code").focus();
        return false;
    }

    $("#loginBox").tips({
        side: 1,
        msg: '正在登录 , 请稍后 ...',
        bg: '#68B500',
        time: 10
    });
    return true
}

// 计算md5
function calcMd5(pwd) {
    let cookiePwdMd5 = $.cookie('password');
    if (pwd !== cookiePwdMd5) {
        let loginName = $("#loginName").val();
        $("#password").val(md5(loginName + pwd))
    }
}

// 保存到Cookie中
function saveCookie() {
    if ($("#saveId").is(":checked")) {
        $.cookie('loginName', $("#loginName").val(), {
            path: '/',
            expires: 7
        });
        $.cookie('password', $("#password").val(), {
            path: '/',
            expires: 7
        })
    }
}

$("#saveId").on('click', function () {
    if (!$("#saveId").is(":checked")) {
        $.cookie('loginName', '', {
            expires: -1
        });
        $.cookie('password', '', {
            expires: -1
        });
        $("#loginName").val('');
        $("#password").val('');
    }
})