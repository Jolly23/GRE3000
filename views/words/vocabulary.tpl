<div class="row">
    <div class="col-md-10 center-block">
        <div class="panel panel-default">
            {{if .UserInfo}}
            {{ if .RandomSort }}
            <div class="panel-heading">
                {{.UserInfo.Username}}的单词表
                <button onclick="meansDisplayToggle(this)" class="btn btn-sm btn-success">展示翻译</button>
                <a href="/words" class="btn btn-sm btn-success">标记排序</a>
            </div>
            {{ else }}
            <div class="panel-heading">
                {{.UserInfo.Username}}的单词表
                <button onclick="meansDisplayToggle(this)" class="btn btn-sm btn-success">展示翻译</button>
                <a href="/words?RandomSort=true" class="btn btn-sm btn-success">随机排序</a>
            </div>
            {{ end }}
            <div class="panel-body" id="words_head">
                {{range .UserWords}}
                <div class="media each_word_block">
                    <div class="col-md-7">
                        <div class="media-body">
                            <div class="title">
                                <a href="https://www.merriam-webster.com/dictionary/{{.Word}}" target="_blank"
                                   class="each_word">{{.Word}}</a>
                            </div>
                            <p class="mean_block">
                                <span>• {{.Mean}}</span>
                            </p>
                            <p>
                                {{if ne .CountMarks 0}}
                                <span>• 没记住{{.CountMarks}}次, 上次标记是在{{.LastMark | FuncFormatTimeAgo}}</span>
                                {{end}}
                            </p>
                        </div>
                    </div>
                    <div class="col-md-offset-7">
                        <button value="{{.Word}}" onclick="ShowMean(this)" type="button"
                                class="btn btn-info btn-lg show1button">
                            显示意思
                        </button>
                        <button value="{{.WordID}}" onclick="MarkWord(this)" type="button"
                                class="btn btn-warning btn-lg">没记住+1
                        </button>
                        <button value="{{.WordID}}" onclick="DeleteWord(this)" type="button"
                                class="btn btn-danger btn-sm">删除
                        </button>
                    </div>
                </div>
                <div class="divide mar-top-5"></div>
                {{end}}
            </div>
            {{else}}
            {{ if .RandomSort }}
            <div class="panel-heading">游客的单词表
                <button onclick="meansDisplayToggle(this)" class="btn btn-sm btn-success">展示翻译</button>
                <a href="/words" class="btn btn-sm btn-success">默认排序</a>
            </div>
            {{ else }}
            <div class="panel-heading">
                游客的单词表
                <button onclick="meansDisplayToggle(this)" class="btn btn-sm btn-success">展示翻译</button>
                <a href="/words?RandomSort=true" class="btn btn-sm btn-success">随机排序</a>
            </div>
            {{ end }}

            <div class="panel-body" id="words_head">
                {{range .RawWords}}
                <div class="media each_word_block">
                    <div class="col-md-7">
                        <div class="media-body">
                            <div class="title">
                                <a href="https://www.merriam-webster.com/dictionary/{{.Word}}" target="_blank"
                                   class="each_word">{{.Word}}</a>
                            </div>
                            <p class="mean_block">
                                <span>• {{.Mean}}</span>
                            </p>
                        </div>
                    </div>
                    <div class="col-md-offset-7">
                        <button onclick="ShowMean(this)" type="button" class="btn btn-info btn-lg show1button">
                            显示意思
                        </button>
                    </div>
                </div>
                <div class="divide mar-top-5"></div>
                {{end}}
            </div>
            {{end}}
        </div>
    </div>
</div>
