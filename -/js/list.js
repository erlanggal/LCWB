$(document).ready(function () {
    $(".fullscreen").css({ height: window.innerHeight - 40 });

    var html = $("#section-html").text();
    $("#section-preview").html(html);

    var li = $(".hp-li");
    var li0 = $(".hp-li:eq(0)");
    var li1 = $(".hp-li:eq(1)");
    var li2 = $(".hp-li:eq(2)");
    var li3 = $(".hp-li:eq(3)");
    var li4 = $(".hp-li:eq(4)");
    li1.click(function(){
        if (li.hasClass("selected")) {
            li.removeClass("selected");
        }
        li1.addClass("selected");
    });
    li2.click(function(){
        if (li.hasClass("selected")) {
            li.removeClass("selected");
        }
        li2.addClass("selected");
    });
    li3.click(function(){
        if (li.hasClass("selected")) {
            li.removeClass("selected");
        }
        li3.addClass("selected");
    });
    li4.click(function(){
        if (li.hasClass("selected")) {
            li.removeClass("selected");
        }
        li4.addClass("selected");
    });
});

function openSidebar() {
    document.getElementById("hm-menu").style.width = "250px";
    document.getElementById("close").style.display = "block";
    document.getElementById("open").style.display = "none";
    disableScroll();
}

function closeSidebar() {
    document.getElementById("hm-menu").style.width = "0";
    document.getElementById("close").style.display = "none";
    document.getElementById("open").style.display = "block";
    enableScroll();
}

var keys = { 37: 1, 38: 1, 39: 1, 40: 1 };

function preventDefault(e) {
    e = e || window.event;
    if (e.preventDefault)
        e.preventDefault();
    e.returnValue = false;
}

function preventDefaultForScrollKeys(e) {
    if (keys[e.keyCode]) {
        preventDefault(e);
        return false;
    }
}

function disableScroll() {
    if (window.addEventListener)
        window.addEventListener('DOMMouseScroll', preventDefault, false);
    window.onwheel = preventDefault;
    window.onmousewheel = document.onmousewheel = preventDefault;
    window.ontouchmove = preventDefault;
    document.onkeydown = preventDefaultForScrollKeys;
}

function enableScroll() {
    if (window.removeEventListener)
        window.removeEventListener('DOMMouseScroll', preventDefault, false);
    window.onmousewheel = document.onmousewheel = null;
    window.onwheel = null;
    window.ontouchmove = null;
    document.onkeydown = null;
}