<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8"/>
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
    <title>Jolly</title>
    <link rel="stylesheet" type="text/css" href="../static/css/index_page.css">
</head>

<body>
<div class="vi">
    <div class="sidebar">

        <div class="header">
            <h1>Jolly</h1>
            <div class="quote">
                <p class="quote-text animate-init">我的征途，是星辰大海。</p>
                <p class="quote-author animate-init"> —— <strong>磊</strong></p>
            </div>
        </div>

        <div class="menu">
            <a href="https://github.com/Jolly23" class="animate-init">GitHub</a>
            <a href="https://jolly23.com" class="animate-init">个人博客</a>
        </div>

        <div class="location">
            <i class="location-icon"></i>
            <span class="location-text animate-init">Hohhot - China</span>
        </div>

    </div>

    <div class="content">
        <span class="close">close</span>
    </div>
</div>


<script type="text/javascript" src="../static/js/index_page.js"></script>
<script>
    $(document).ready(function () {
        var delay = 1;
        var DELAY_STEP = 200;
        var animationOptions = {opacity: 1, top: 0};

        $('h1').animate(animationOptions).promise().pipe(animateMain).pipe(animateLocationIcon);

        function animateMain() {
            var dfd = $.Deferred();
            var els = $('.animate-init');
            var size = els.size();

            els.each(function (index, el) {
                delay++;
                $(el).delay(index * DELAY_STEP).animate(animationOptions);
                (size - 1 === index) && dfd.resolve();
            });
            return dfd.promise();
        }

        function animateLocationIcon() {
            $('.location-icon').delay(delay * DELAY_STEP).animate({
                opacity: 1,
                top: 0
            }).promise().done(animationQuote);
        }

        function animationQuote() {
        }
    });
</script>
</body>
</html>
