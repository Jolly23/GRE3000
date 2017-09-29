<!DOCTYPE html>
<html lang="en">
<head>
    <title>照片上传</title>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width,initial-scale=1,user-scalable=0">
    <link rel="stylesheet" href="//cdn.bootcss.com/weui/0.4.3/style/weui.min.css">
    <link rel="stylesheet" href="//cdn.bootcss.com/jquery-weui/0.8.0/css/jquery-weui.min.css">
</head>

<body>
<div class='container'>
    <form class="form-horizontal" role="form" action='' method='post'>
        {{ .xsrfdata }}
        <div class="weui_msg">
            <div class="weui_opr_area" id="upload_btn_div">
                <p class="weui_btn_area">
                    <a href="javascript:" class="weui_btn weui_btn_primary" id="upload_btn">上传</a>
                </p>
            </div>
        </div>
    </form>
</div>


<script src="//cdn.bootcss.com/jquery/1.11.0/jquery.min.js"></script>
<script src="//cdn.bootcss.com/jquery-weui/0.8.0/js/jquery-weui.min.js"></script>
<script src="https://res.wx.qq.com/open/js/jweixin-1.2.0.js"></script>

<script>
    $.ajax({
        type: "post",
        data: {
//            id_code: 'p4BoE',
            id_code: 'B8BB1',
            _xsrf: $("[name='_xsrf']").val(),
            uri: window.location.href
        },
        url: 'https://go-wx.jolly23.com/create_js',
        dataType: "json",
        success: function (data) {
            wx.config({
                debug: true,
                appId: data['AppID'],
                timestamp: data['TimeStamp'],
                nonceStr: data['NonceStr'],
                signature: data['Signature'],
                jsApiList: ['chooseImage', 'uploadImage', 'downloadImage', 'hideOptionMenu']
            });

            wx.ready(function () {
                wx.hideOptionMenu();
                wx.chooseImage({
                    count: 1, // 默认9
                    sizeType: ['original'], // 可以指定是原图还是压缩图，默认二者都有
                    sourceType: ['album', 'camera'], // 可以指定来源是相册还是相机，默认二者都有
                    success: function (res) {
                        var localIds = res.localIds; // 返回选定照片的本地ID列表，localId可以作为img标签的src属性显示图片
                        $("#upload_btn").click(function () {
                            wx.uploadImage({
                                localId: localIds[0], // 需要上传的图片的本地ID，由chooseImage接口获得
                                isShowProgressTips: 1, // 默认为1，显示进度提示
                                success: function (res) {
                                    $.post(
                                        "/wx_upload_photo",
                                        {
                                            id_code: 'p4BoE',
                                            media_id: res.serverId,
                                            _xsrf: $("[name='_xsrf']").val()
                                        },
                                        function (data) {
                                            $.showLoading();
                                            if (data['errcode'] === 0) {
                                                $.toast("OK");
                                            }
                                            else {
                                                $.hideLoading();
                                                $.alert("上传失败");
                                            }
                                        });
                                }
                            });
                        });
                    },
                    cancel: function () {
                        $.alert("失败");
                    }
                });
            });
        },
        error: function (xhr, type) {
            location.reload();
        }
    });
</script>

</body>
</html>
