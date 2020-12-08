var a = 0;
$('#fullscreen').on('click', function() {
    a++;
    a % 2 == 1 ? enterfullscreen() : exitfullscreen();
})

//进入全屏
function enterfullscreen() { 
    var docElm = document.documentElement;

    //W3C
    if(docElm.requestFullscreen) {
        docElm.requestFullscreen();
    }
    //FireFox
    else if(docElm.mozRequestFullScreen) {
        docElm.mozRequestFullScreen();
    }
    //Chrome等
    else if(docElm.webkitRequestFullScreen) {
        docElm.webkitRequestFullScreen();
    }
    //IE11
    else if(elem.msRequestFullscreen) {
        elem.msRequestFullscreen();
    }
}

//退出全屏
function exitfullscreen() { 
     //W3C
    if(document.exitFullscreen) {
        document.exitFullscreen();
    }
    //FireFox
    else if (document.mozCancelFullScreen) {
        document.mozCancelFullScreen();
    }
    //Chrome等
    else if (document.webkitCancelFullScreen) {
        document.webkitCancelFullScreen();
    }
    //IE11
    else if (document.msExitFullscreen) {
        document.msExitFullscreen();
    }
}