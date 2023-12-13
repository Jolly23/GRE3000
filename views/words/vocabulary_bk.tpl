<div class="row">
    <div class="col-md-9">
        {{if .UserInfo}}
        <div class="panel panel-default">
            {{ if .ShowMean }}
                {{ if .RandomSort }}
                <div class="panel-heading">{{.UserInfo.Username}}的单词表 <a href="/words?RandomSort=true" class="btn btn-sm btn-success">关闭翻译</a> <a href="/words?ShowMean=true" class="btn btn-sm btn-success">标记排序</a></div>
                {{ else }}
                <div class="panel-heading">{{.UserInfo.Username}}的单词表 <a href="/words" class="btn btn-sm btn-success">关闭翻译</a> <a href="/words?RandomSort=true&ShowMean=true" class="btn btn-sm btn-success">随机排序</a></div>
                {{ end }}
            {{ else }}
                {{ if .RandomSort }}
                <div class="panel-heading">{{.UserInfo.Username}}的单词表 <a href="/words?ShowMean=true&RandomSort=true" class="btn btn-sm btn-success">显示翻译</a> <a href="/words" class="btn btn-sm btn-success">标记排序</a></div>
                {{ else }}
                <div class="panel-heading">{{.UserInfo.Username}}的单词表 <a href="/words?ShowMean=true" class="btn btn-sm btn-success">显示翻译</a> <a href="/words?RandomSort=true" class="btn btn-sm btn-success">随机排序</a></div>
                {{ end }}
            {{ end }}
            <div class="panel-body" id="words_head">
                {{ if .RandomSort }}
                <p style="text-align: center">词表加载中，请稍等</p>
                {{ end }}
                {{range .UserWords}}
                <div class="media">
                    <div class="col-lg-7">
                        <div class="media-body">
                            <div class="title">
                                <a href="https://www.merriam-webster.com/dictionary/{{.Word}}" target="_blank" id="each_word">{{.Word}}</a>
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
                    <div class="col-lg-offset-7">
                        <button value="{{.Word}}" onclick="ShowMean(this)" type="button" class="btn btn-info btn-lg show1button">显示意思</button>
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
            {{ if .ShowMean }}
                {{ if .RandomSort }}
                <div class="panel-heading">游客的单词表 <a href="/words?RandomSort=true" class="btn btn-sm btn-success">关闭翻译</a> <a href="/words?ShowMean=true" class="btn btn-sm btn-success">默认排序</a></div>
                {{ else }}
                <div class="panel-heading">游客的单词表 <a href="/words" class="btn btn-sm btn-success">关闭翻译</a> <a href="/words?ShowMean=true&RandomSort=true" class="btn btn-sm btn-success">随机排序</a></div>
                {{ end }}
            {{ else }}
                {{ if .RandomSort }}
                <div class="panel-heading">游客的单词表 <a href="/words?ShowMean=true&RandomSort=true" class="btn btn-sm btn-success">显示翻译</a> <a href="/words" class="btn btn-sm btn-success">默认排序</a></div>
                {{ else }}
                <div class="panel-heading">游客的单词表 <a href="/words?ShowMean=true" class="btn btn-sm btn-success">显示翻译</a> <a href="/words?RandomSort=true" class="btn btn-sm btn-success">随机排序</a></div>
                {{ end }}
            {{ end }}
            <div class="panel-body" id="words_head">
                {{ if .RandomSort }}
                <p style="text-align: center">词表加载中，请稍等</p>
                {{ end }}
                {{range .RawWords}}
                <div class="media">
                    <div class="col-lg-7">
                        <div class="media-body">
                            <div class="title">
                                <a href="https://www.merriam-webster.com/dictionary/{{.Word}}" target="_blank" id="each_word">{{.Word}}</a>
                            </div>
                            <p class="mean_block">
                                <span>• {{.Mean}}</span>
                            </p>
                        </div>
                    </div>
                    <div class="col-lg-offset-7">
                        <button onclick="ShowMean(this)" type="button" class="btn btn-info btn-lg">显示意思</button>
                    </div>
                </div>
                <div class="divide mar-top-5"></div>
                {{end}}
            </div>
        </div>
        {{end}}
    </div>
</div>


<script type="text/javascript">
    $(document).ready(function () {

        LoadWords();
    });
</script>
