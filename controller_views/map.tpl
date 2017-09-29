<!DOCTYPE html>
<html>
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
    <meta name="viewport" content="initial-scale=1.0, user-scalable=no"/>
    <style type="text/css">
        body, html, #big_map {
            width: 100%;
            height: 100%;
            overflow: hidden;
            margin: 0;
        }
    </style>
    <script type="text/javascript"
            src="https://api.map.baidu.com/api?v=2.0&ak=lIxKv2vLqpxN7a81kR3McI4iAlXj16Gx"></script>
    <title>行踪</title>
</head>

<body>

<div id="big_map"></div>

</body>

<script type="text/javascript">
    map = new BMap.Map("big_map");
    var navigationControl = new BMap.NavigationControl({
        anchor: BMAP_ANCHOR_TOP_LEFT,
        type: BMAP_NAVIGATION_CONTROL_LARGE,
        enableGeolocation: true
    });
    map.addControl(navigationControl);
    map.centerAndZoom(new BMap.Point({{ .center_lng }}, {{ .center_lat }}), 12);

    var opts = {
        width: 100,
        height: 30,
        title: "记录时间",
        enableMessage: false
    };

    var pins = JSON.parse('{{ .pins }}');

    for (var i = 0; i < pins.length; i++) {
        var gps_point = new BMap.Point(pins[i].lng, pins[i].lat);
        var marker = new BMap.Marker(gps_point);
        var content = pins[i].time;
        map.addOverlay(marker);
        addClickHandler(content, marker);
        if (i === 0) {
            marker.setAnimation(BMAP_ANIMATION_BOUNCE);

            var point_dlnu = new BMap.Point(121.772636, 39.042618);
            alert('最近行踪 ' + pins[i].time + ' 距离大连民族大学：' + (map.getDistance(point_dlnu, gps_point)).toFixed(2) + ' 米。');
        }
    }

    function addClickHandler(content, marker) {
        marker.addEventListener("click", function (e) {
                openInfo(content, e)
            }
        );
    }

    function openInfo(content, e) {
        var p = e.target;
        var point = new BMap.Point(p.getPosition().lng, p.getPosition().lat);
        var infoWindow = new BMap.InfoWindow(content, opts);
        map.openInfoWindow(infoWindow, point);
    }

</script>

</html>
