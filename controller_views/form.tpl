<!DOCTYPE html>

<html>
<head>
    <title>添加平台</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
</head>

<body>

<form id="PostFormInfo" method="post">
    {{ .xsrfdata }}

    School：<input name="school" type="text"/><br>
    PlatformName：<input name="platform_name" type="text"/><br>
    AppID：<input name="appid" type="text"/><br>
    AppSecret：<input name="appsecret" type="text"/><br>
    PlatformID：<input name="platform_id" type="text"/><br>
    Token：<input name="token" type="text"/><br>
    EncodingAESKey：<input name="encodingaeskey" type="text"/><br>
    <input type="submit" value="提交"/>
</form>

</body>
</html>
