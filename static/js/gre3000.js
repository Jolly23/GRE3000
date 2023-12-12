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