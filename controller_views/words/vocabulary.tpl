<div class="row">
    <div class="col-md-9">
        {{if .IsLogin}}
        <div class="panel panel-default">
            <div class="panel-heading">{{.UserInfo.Username}}的单词表</div>
            <div class="panel-body">
                {{range .UserWords}}
                <div id="{{.Word.Word}}" class="media">
                    <div class="col-lg-7">
                        <div class="media-body">
                            <div class="title">
                                <a style="font-size: 30px;">{{.Word.Word}}</a>
                            </div>
                            <p id="btn-means-{{.Word.Word}}-obj" style="display: none;font-size: 20px;">
                                <span>• {{.Word.Means}}</span>
                            </p>
                            <p>
                                {{if ne .CountMarks 0}}
                                <span>• 没记住{{.CountMarks}}次, 上次标记是在{{.LastMark | FuncFormatTimeAgo}}</span>
                                {{end}}
                            </p>
                        </div>
                    </div>
                    <div class="col-lg-offset-7">
                        <button onclick="ShowMeans(this)" type="button" class="btn btn-info btn-lg">显示意思</button>
                        <button value="{{.Id}}" onclick="MarkWord(this)" type="button" class="btn btn-warning btn-lg">没记住+1</button>
                        <button value="{{.Id}}" onclick="DeleteWord(this)" type="button" class="btn btn-danger btn-sm">删除</button>
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
                    <div class="col-lg-7">
                        <div class="media-body">
                            <div class="title">
                                <a style="font-size: 30px;">{{.Word}}</a>
                            </div>
                            <p id="btn-means-{{.Word}}-obj" style="display: none;font-size: 20px;">
                                <span>• {{.Means}}</span>
                            </p>
                        </div>
                    </div>
                    <div class="col-lg-offset-7">
                        <button id="btn-means-{{.Word}}" onclick="ShowMeans(this)" type="button" class="btn btn-info btn-lg">显示意思</button>
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
