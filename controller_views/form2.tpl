<!DOCTYPE html>

<html>
<head>
    <title>DEBUG</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
</head>

<body>

<form id="AccountFormInfo" method="post">
    {{ .xsrfdata }}

    uid：<input name="uid" type="text" />
    pwd：<input name="pwd" type="text" />
    openid：<input name="openid" type="text" />
    <input type="submit" value="提交" />
</form>

</body>
</html>
