<!doctype html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport"
          content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>{{.PageTitle}}</title>
    <link rel="stylesheet" href="/static/bootstrap/css/bootstrap.min.css"/>
    <link rel="stylesheet" href="/static/css/ger3000.css">
    <script src="//cdn.bootcss.com/jquery/2.2.2/jquery.min.js"></script>
    <script src="//cdn.bootcss.com/bootstrap/3.3.6/js/bootstrap.min.js"></script>
    {{if .IsWordsPage}}
    <script type="text/javascript">
        function ShowMeans(whichObj) {
            $(whichObj).parent().parent().find("p").css('display', 'block');
            whichObj.disabled = true;
        }

        function MarkWord(whichObj) {
            whichObj.disabled = true;
            $.get("/words/mark/" + $(whichObj).attr("value"));
        }

        function DeleteWord(whichObj) {
            whichObj.disabled = true;
            $.get("/words/del/" + $(whichObj).attr("value"));
        }
    </script>
    <style type="text/css">
        #each_word {
            font-size: 30px;
            color:darkblue;
        }
        #means {
            {{ if not .ShowMeans }}
            display: none;
            {{end}}
            font-size: 20px;
        }
    </style>
    {{end}}
</head>
<body>

<div class="wrapper">
    <nav class="navbar navbar-inverse">
        <div class="container">
            <div class="navbar-header">
                <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#navbar"
                        aria-expanded="false" aria-controls="navbar">
                    <span class="sr-only">Toggle navigation</span>
                    <span class="icon-bar"></span>
                    <span class="icon-bar"></span>
                    <span class="icon-bar"></span>
                </button>
                <span><a class="navbar-brand" style="color:#fff;" href="/">GRE3000</a></span>
                <span><a class="navbar-brand" style="color:#22ffff;" href="/words">点我进词表</a></span>
            </div>
            <div id="navbar" class="navbar-collapse collapse header-navbar">
                <ul class="nav navbar-nav navbar-right">
                    <li>
                        <a href="/about">关于</a>
                    </li>
                    {{if .IsLogin}}
                    <li>
                        <a href="/user/{{.UserInfo.Username}}">
                            {{.UserInfo.Username}}
                        </a>
                    </li>
                    <li>
                        <a href="javascript:;" class="dropdown-toggle" data-toggle="dropdown"
                           data-hover="dropdown">
                            设置
                            <span class="caret"></span>
                        </a>
                        <span class="dropdown-arrow"></span>
                        <ul class="dropdown-menu">
                            <li><a href="/user/setting">个人资料</a></li>
                            {{if FuncHasPermission .UserInfo.Id "user:list"}}
                            <li><a href="/user/list">用户管理</a></li>
                            {{end}}
                            {{if FuncHasPermission .UserInfo.Id "role:list"}}
                            <li><a href="/role/list">角色管理</a></li>
                            {{end}}
                            {{if FuncHasPermission .UserInfo.Id "permission:list"}}
                            <li><a href="/permission/list">权限管理</a></li>
                            {{end}}
                            <li><a href="/logout">退出</a></li>
                        </ul>
                    </li>
                    {{else}}
                    <li><a href="/login">登录</a></li>
                    <li><a href="/register">注册</a></li>
                    {{end}}
                </ul>
            </div>
        </div>
    </nav>
    <div class="container">
        {{.LayoutContent}}
    </div>
</div>
<div class="container">
    <br>
    <div class="text-center">
        ©2017 Powered by <a href="https://jolly23.com" target="_blank">Jolly23</a><br>
        <script type="text/javascript">var cnzz_protocol = (("https:" == document.location.protocol) ? " https://" : " http://");document.write(unescape("%3Cspan id='cnzz_stat_icon_1265971333'%3E%3C/span%3E%3Cscript src='" + cnzz_protocol + "s13.cnzz.com/z_stat.php%3Fid%3D1265971333%26show%3Dpic' type='text/javascript'%3E%3C/script%3E"));</script>
    </div>
    <br>
</div>

</body>

</html>
