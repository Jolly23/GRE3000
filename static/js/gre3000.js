$(document).ready(function () {
    if (window.location.pathname !== '/words') {
        return;
    }

    randomDisplayToggle();
    asyncLoadWords();
});

function meansDisplayToggle(whichObj) {
    if (whichObj.textContent === "展示翻译") {
        whichObj.textContent = "隐藏翻译";
        const displayBlock = document.getElementsByClassName('mean_block');
        for (let i = 0; i < displayBlock.length; i++) {
            displayBlock[i].style.display = 'block';
        }

        const displayBtn = document.getElementsByClassName('show1button');
        for (let i = 0; i < displayBtn.length; i++) {
            displayBtn[i].disabled = true;
        }

    } else if (whichObj.textContent === "隐藏翻译") {
        whichObj.textContent = "展示翻译";
        const displayBlock = document.getElementsByClassName('mean_block');
        for (let i = 0; i < displayBlock.length; i++) {
            displayBlock[i].style.display = 'none';
        }

        const displayBtn = document.getElementsByClassName('show1button');
        for (let i = 0; i < displayBtn.length; i++) {
            displayBtn[i].disabled = false;
        }
    }
}

function randomDisplayToggle() {
    let isRandomSort = window.location.search.indexOf('rs=true') !== -1;

    const btnUserSort = document.getElementById('sort-btn-user');
    if (btnUserSort !== null) {
        if (isRandomSort) {
            btnUserSort.href = '/words';
            btnUserSort.textContent = '标记排序';
        } else {
            btnUserSort.href = '/words?rs=true';
            btnUserSort.textContent = '随机排序';
        }
        return;
    }

    const btnVisitor = document.getElementById('sort-btn');
    if (btnVisitor !== null) {
        if (isRandomSort) {
            btnVisitor.href = '/words';
            btnVisitor.textContent = '默认排序';
        } else {
            btnVisitor.href = '/words?rs=true';
            btnVisitor.textContent = '随机排序';
        }
    }
}

function ShowMean(whichObj) {
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


function asyncLoadWords() {

    let isRandomSort = window.location.search.indexOf('rs=true') !== -1;
    let load_url = "/words/load";
    if (isRandomSort) {
        load_url += "?rs=true";
    }

    let str = "";
    let i = 0;

    $.ajaxSetup({cache: false});
    $.get(
        load_url, function (data) {
            if (data.length > 0 && data[0]['c'] !== undefined) {
                for (i = 0; i < data.length; i++) {
                    const notice_text = data[i]['c'] === 0 ? "" : "<span>• 没记住" + data[i]['c'] + "次, 上次标记是在" + data[i]['t'] + "</span>";
                    str += `<div class="media each_word_block">
                    <div class="col-md-7">
                        <div class="media-body">
                            <div class="title">
                                <a href="https://www.merriam-webster.com/dictionary/{{.Word}}" target="_blank"
                                   class="each_word">` + data[i]['w'] + `</a>
                            </div>
                            <p class="mean_block">
                                <span>• ` + data[i]['m'] + `</span>
                            </p>
                            <p>
                                ` + notice_text + `
                            </p>
                        </div>
                    </div>
                    <div class="col-md-offset-7">
                        <button value="` + data[i]['w'] + `" onclick="ShowMean(this)" type="button"
                                class="btn btn-info btn-lg show1button">
                            显示意思
                        </button>
                        <button value="` + data[i]['i'] + `" onclick="MarkWord(this)" type="button"
                                class="btn btn-warning btn-lg">没记住+1
                        </button>
                        <button value="` + data[i]['i'] + `" onclick="DeleteWord(this)" type="button"
                                class="btn btn-danger btn-sm">删除
                        </button>
                    </div>
                </div>`
                }
            } else {
                for (i = 0; i < data.length; i++) {
                    str += `<div class="media each_word_block">
                    <div class="col-md-7">
                        <div class="media-body">
                            <div class="title">
                                <a href="https://www.merriam-webster.com/dictionary/{{.Word}}" target="_blank"
                                   class="each_word">` + data[i]['w'] + `</a>
                            </div>
                            <p class="mean_block">
                                <span>• ` + data[i]['m'] + `</span>
                            </p>
                        </div>
                    </div>
                    <div class="col-md-offset-7">
                        <button onclick="ShowMean(this)" type="button" class="btn btn-info btn-lg show1button">
                            显示意思
                        </button>
                    </div>
                </div>`
                }
            }
            document.getElementById('words_head').innerHTML = "";
            document.getElementById("words_head").insertAdjacentHTML('beforeend', str);

            // if (data.length > 0 && data[0]['CountMarks'] !== undefined) {
            //     $.get(
            //         '/words/statistics',  function (data) {
            //             str = "<br><p style=\"text-align: center\">我的词表共有" + data['All'] + "个单词，标记单词" + data['Marked'] + "个</p>";
            //             document.getElementById("words_head").insertAdjacentHTML('beforeend', str);
            //         }
            //     );
            // }
        }
    );
}