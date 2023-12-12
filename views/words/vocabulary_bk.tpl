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
        {{ if .ShowMean }}
        var disable_tag = " disabled";
        $("button.btn-info:not(:disabled)").attr('disabled', true);
        {{ else }}
        var disable_tag = "";
        {{end}}

        {{ if .RandomSort }}
        var sync_load_url = "/words/load_words?RandomSort=true";
        {{ else }}
        var sync_load_url = "/words/load_words";
        {{end}}

        var str = "";
        var i = 0;

        $.ajaxSetup({ cache: false });
        $.get(
            sync_load_url, function(data) {
                if (data.length > 0 && data[0]['CountMarks'] !== undefined) {
                    for (i = 0; i < data.length; i++) {
                        var notice_text = data[i]['CountMarks'] === 0 ? "" : "<span>• 没记住" + data[i]['CountMarks'] + "次, 上次标记是在" + data[i]['LastMark'] + "</span>" ;
                        str += "<div class=\"media\"><div class=\"col-lg-7\"><div class=\"media-body\"><div class=\"title\"><a href=\"https://www.merriam-webster.com/dictionary/" + data[i]['Word'] + "\" target=\"_blank\" id=\"each_word\">" + data[i]['Word'] + "</a></div><p id=\"mean\"><span>• " + data[i]['Mean'] + "</span></p><p>" + notice_text + "</p></div></div><div class=\"col-lg-offset-7\"><button onclick=\"ShowMean(this)\" type=\"button\" class=\"btn btn-info btn-lg" + disable_tag + "\">显示意思</button> <button value=\"" + data[i]['Id'] + "\" onclick=\"MarkWord(this)\" type=\"button\" class=\"btn btn-warning btn-lg\">没记住+1</button> <button value=\"" + data[i]['Id'] + "\" onclick=\"DeleteWord(this)\" type=\"button\" class=\"btn btn-danger btn-sm\">删除</button></div></div><div class=\"divide mar-top-5\"></div>";
                    }
                } else {
                    for (i = 0; i < data.length; i++) {
                        str += "<div class=\"media\"><div class=\"col-lg-7\"><div class=\"media-body\"><div class=\"title\"><a href=\"https://www.merriam-webster.com/dictionary/" + data[i]['word'] + "\" target=\"_blank\" id=\"each_word\">" + data[i]['word'] + "</a></div><p id=\"mean\"><span>• " + data[i]['mean'] + "</span></p></div></div><div class=\"col-lg-offset-7\"><button onclick=\"ShowMean(this)\" type=\"button\" class=\"btn btn-info btn-lg" + disable_tag + "\">显示意思</button></div></div><div class=\"divide mar-top-5\"></div>"
                    }
                }
                {{ if .RandomSort }}
                document.getElementById('words_head').innerHTML="";
                {{ end }}
                document.getElementById("words_head").insertAdjacentHTML('beforeend', str);

                if (data.length > 0 && data[0]['CountMarks'] !== undefined) {
                    $.get(
                        '/words/statistics',  function (data) {
                            str = "<br><p style=\"text-align: center\">我的词表共有" + data['All'] + "个单词，标记单词" + data['Marked'] + "个</p>";
                            document.getElementById("words_head").insertAdjacentHTML('beforeend', str);
                        }
                    );
                }
            }
        );
    });
</script>
