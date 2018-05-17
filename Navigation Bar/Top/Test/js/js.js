$(document).ready(function () {
    $(".fullscreen").css({ height: window.innerHeight - 40 });

    var html = $("#section-html").text();
    $("#section-preview").html(html);

    var li = $(".hp-li");
    var li0 = $(".hp-li:eq(0)");
    var li1 = $(".hp-li:eq(1)");
    var li2 = $(".hp-li:eq(2)");
    var li3 = $(".hp-li:eq(3)");
    li0.click(function(){
        setAllDisplayNone(li);
        li0.addClass("selected");
        $("#section-preview").css("display", "block");
    });
    li1.click(function(){
        setAllDisplayNone(li);
        li1.addClass("selected");
        $("#section-html").css("display", "block");
    });
    li2.click(function(){
        setAllDisplayNone(li);
        li2.addClass("selected");
        $("#section-css").css("display", "block");
    });
    li3.click(function(){
        setAllDisplayNone(li);
        li3.addClass("selected");
        $("#section-js").css("display", "block");
    });

    var box1 = $("#td-box-1");
    var box2 = $("#td-box-2");
    var box3 = $("#td-box-3");
    var box4 = $("#td-box-4");
    var b1 = b2 = b3 = b4 = false;

    $(box1).click(function(){
        if (b1) {
            $(this).removeClass('border-blue');
            $(this).addClass('border-black');
            b1 = false; 
        } else if (!b1) {
            if (b2) {
                $(box2).removeClass('border-blue');
                $(box2).addClass('border-black');
                b2 = false;
            }
            if (b3) {
                $(box3).removeClass('border-blue');
                $(box3).addClass('border-black');
                b3 = false;
            }
            if (b4) {
                $(box4).removeClass('border-blue');
                $(box4).addClass('border-black');
                b4 = false;
            }
            $(this).removeClass('border-black');
            $(this).addClass('border-blue');
            b1 = true;
        }
    });
    $(box2).click(function(){
        if (b2) {
            $(this).removeClass('border-blue');
            $(this).addClass('border-black');
            b2 = false; 
        } else if (!b2) {
            if (b1) {
                $(box1).removeClass('border-blue');
                $(box1).addClass('border-black');
                b1 = false;
            }
            if (b3) {
                $(box3).removeClass('border-blue');
                $(box3).addClass('border-black');
                b3 = false;
            }
            if (b4) {
                $(box4).removeClass('border-blue');
                $(box4).addClass('border-black');
                b4 = false;
            }
            $(this).removeClass('border-black');
            $(this).addClass('border-blue');
            b2 = true;
        }
    });
    $(box3).click(function(){
        if (b3) {
            $(this).removeClass('border-blue');
            $(this).addClass('border-black');
            b3 = false; 
        } else if (!b3) {
            if (b1) {
                $(box1).removeClass('border-blue');
                $(box1).addClass('border-black');
                b1 = false;
            }
            if (b2) {
                $(box2).removeClass('border-blue');
                $(box2).addClass('border-black');
                b2 = false;
            }
            if (b4) {
                $(box4).removeClass('border-blue');
                $(box4).addClass('border-black');
                b4 = false;
            }
            $(this).removeClass('border-black');
            $(this).addClass('border-blue');
            b3 = true;
        }
    });
    $(box4).click(function(){
        if (b4) {
            $(this).removeClass('border-blue');
            $(this).addClass('border-black');
            b4 = false; 
        } else if (!b4) {
            if (b1) {
                $(box1).removeClass('border-blue');
                $(box1).addClass('border-black');
                b1 = false;
            }
            if (b3) {
                $(box3).removeClass('border-blue');
                $(box3).addClass('border-black');
                b3 = false;
            }
            if (b3) {
                $(box3).removeClass('border-blue');
                $(box3).addClass('border-black');
                b4 = false;
            }
            $(this).removeClass('border-black');
            $(this).addClass('border-blue');
            b4 = true;
        }
    });
      
});

function setAllDisplayNone(li){
    if (li.hasClass("selected")) {
        li.removeClass("selected");
    }
    $("#section-preview").css("display", "none");
    $("#section-html").css("display", "none");
    $("#section-css").css("display", "none");
    $("#section-js").css("display", "none");
}

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