<div class="row">
    <div class="col-md-10 center-block">
        <div class="panel panel-default">
            {{if .UserInfo}}
            <div class="panel-heading">
                {{.UserInfo.Username}}的单词表
                <button onclick="meansDisplayToggle(this)" class="btn btn-sm btn-success">展示翻译</button>
                <a href="" class="btn btn-sm btn-success" id="sort-btn-user"></a>
            </div>

            <div class="panel-body" id="words_head">
                <p style="text-align: center">词表加载中，请稍等</p>

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
            <div class="panel-heading">游客的单词表
                <button onclick="meansDisplayToggle(this)" class="btn btn-sm btn-success">展示翻译</button>
                <a href="" class="btn btn-sm btn-success" id="sort-btn"></a>
            </div>

            <div class="panel-body" id="words_head">
                <p style="text-align: center">词表加载中，请稍等</p>
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
