<div class="row">
    <div class="col-md-9">
        {{if .IsLogin}}
        <div class="panel panel-default">
            <div class="panel-heading">{{.UserInfo.Username}}的单词表</div>
            <div class="panel-body">
                {{range .UserWords}}
                <div class="media">
                    <div class="media-body">
                        <div class="title">
                            <a>{{.Word.Word}}</a>
                            <a href="/words/mark/{{.Word.Id}}"><span class="label label-primary">&emsp;没记住+1&emsp;</span></a>

                        </div>
                        <p>
                            <span>•</span>
                            <span>{{.Word.Means}}</span>
                        </p>
                        <p>
                            {{if ne .CountMarks 0}}
                            <span>•</span>
                            <span>没记住{{.CountMarks}}次, 上次标记是在{{.LastMark | timeago}}</span>
                            {{end}}
                        </p>
                    </div>
                </div>
                <div class="divide mar-top-5"></div>
                {{end}}
            </div>
        </div>
        {{else}}
        <div class="panel panel-default">
            <div class="panel-heading">游客的单词表</div>
            <div class="panel-body">
                {{range .RawWords}}
                <div class="media">
                    <div class="media-body">
                        <div class="title">
                            <a>{{.Word}}</a>
                            <a href="#"><span class="label label-primary">&emsp;没记住+1&emsp;</span></a>

                        </div>
                        <p>
                            <!--<span>&emsp;•</span>-->
                            <span>•</span>
                            <span>{{.Means}}</span>
                        </p>
                        <p>
                            <span>•</span>
                            <span>没记住5次</span>
                        </p>
                    </div>
                </div>
                <div class="divide mar-top-5"></div>
                {{end}}
            </div>
        </div>
        {{end}}
    </div>
    <div class="col-md-3 hidden-sm hidden-xs">

    </div>
</div>