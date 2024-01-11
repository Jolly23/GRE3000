<div class="row">
    <div class="col-md-10 center-block">
        <div class="panel panel-default">
            {{if .UserInfo}}
            <div class="panel-heading">
                {{.UserInfo.Username}}的单词表
                <button onclick="meansDisplayToggle(this)" class="btn btn-sm btn-success">展示翻译</button>
                <a href="" class="btn btn-sm btn-success" id="sort-btn-user"></a>
            </div>
            {{else}}
            <div class="panel-heading">游客的单词表
                <button onclick="meansDisplayToggle(this)" class="btn btn-sm btn-success">展示翻译</button>
                <a href="" class="btn btn-sm btn-success" id="sort-btn"></a>
            </div>
            {{end}}

            <div class="panel-body" id="words_head">
                <p style="text-align: center">词表加载中，请稍等</p>
            </div>
        </div>
    </div>
</div>
