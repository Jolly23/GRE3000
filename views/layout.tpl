<!doctype html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport"
          content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>{{.PageTitle}}</title>
    <link rel="stylesheet" href="/static/bootstrap/css/bootstrap.min.css"/>
    <script src="//cdn.bootcss.com/jquery/2.2.2/jquery.min.js"></script>
    <script src="//cdn.bootcss.com/bootstrap/3.3.6/js/bootstrap.min.js"></script>

    <link rel="stylesheet" href="/static/css/gre3000.css">
    <script type="text/javascript" src="/static/js/gre3000.js"></script>
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
                {{if .UserInfo}}
                <span><a class="navbar-brand" style="color:#fff;" href="/">{{.UserInfo.Username}}'s GRE3000</a></span>
                {{else}}
                <span><a class="navbar-brand" style="color:#fff;" href="/">Visitor's GRE3000</a></span>
                {{end}}
            </div>
            <div id="navbar" class="navbar-collapse collapse header-navbar">
                <ul class="nav navbar-nav navbar-right">
                    <li>
                        <a href="/about">About</a>
                    </li>
                    {{if .UserInfo}}
                    <li>
                        <a href="/words/rebuild">
                            VocabularyReset
                        </a>
                    </li>
                    <li><a href="/logout">Logout</a></li>
                    {{else}}
                    <li><a href="/login">Login</a></li>
                    <li><a href="/register">Register</a></li>
                    {{end}}
                </ul>
            </div>
        </div>
    </nav>
    <div class="container">
        {{embed}}
    </div>
</div>
<div class="container">
    <br>
    <div class="text-center">©2023 Powered by Mars<br>
    </div>
    <br>
</div>

</body>

</html>
