//本地缓存处理
let storage = {
    set: function (key, value) {
        window.localStorage.setItem(key, value);
    },
    get: function (key) {
        return window.localStorage.getItem(key);
    },
    remove: function (key) {
        window.localStorage.removeItem(key);
    },
    clear: function () {
        window.localStorage.clear();
    }
};

/**
 * iframe tab处理
 */

$(function () {
    //通过遍历给菜单项加上data-index属性
    $(".menuItem").each(function (index) {
        if (!$(this).attr('data-index')) {
            $(this).attr('data-index', index);
        }
    });

    $('.menuItem').on('click', menuItem);
    // 左移按扭
    $('.tabLeft').on('click', scrollTabLeft);
    // 右移按扭
    $('.tabRight').on('click', scrollTabRight);
    $('.menuTabs').on('click', '.menuTab i', closeTab);
    // 点击选项卡菜单
    $('.menuTabs').on('click', '.menuTab', activeTab);
    $('.tabShowActive').on('click', showActiveTab);
    // 页签刷新按钮
    $('.tabReload').on('click', refreshTab);
    // 关闭当前
    $('.tabCloseCurrent').on('click', tabCloseCurrent);
    // 关闭其他
    $('.tabCloseOther').on('click', tabCloseOther);
    // 关闭全部
    $('.tabCloseAll').on('click', tabCloseAll);

    //滚动到已激活的选项卡
    function showActiveTab() {
        scrollToTab($('.menuTab.active'));
    }

    function menuItem() {
        // 获取标识数据
        let dataUrl = $(this).attr('href'),
            dataIndex = $(this).data('index'),
            menuName = $.trim($(this).text()),
            flag = true;
        $(".nav ul li, .nav li").removeClass("active");
        $(this).parent("li").addClass("active");
        setIframeUrl($(this).attr("href"));
        if (dataUrl == undefined || $.trim(dataUrl).length == 0) return false;

        // 选项卡菜单已存在
        $('.menuTab').each(function () {
            if ($(this).data('id') == dataUrl) {
                if (!$(this).hasClass('active')) {
                    $(this).addClass('active').siblings('.menuTab').removeClass('active');
                    scrollToTab(this);
                    // 显示tab对应的内容区
                    $('.mainContent .xframe').each(function () {
                        if ($(this).data('id') == dataUrl) {
                            $(this).show().siblings('.xframe').hide();
                            return false;
                        }
                    });
                }
                flag = false;
                return false;
            }
        });
        // 选项卡菜单不存在
        if (flag) {
            let str = '<a href="javascript:;" class="active menuTab" data-id="' + dataUrl + '">' + menuName + ' <i class="fa fa-times-circle"></i></a>';
            $('.menuTab').removeClass('active');

            // 添加选项卡对应的iframe
            let str1 = '<iframe class="xframe" name="iframe' + dataIndex + '" width="100%" height="100%" src="' + dataUrl + '" frameborder="0" data-id="' + dataUrl + '" seamless></iframe>';
            $('.mainContent').find('iframe.xframe').hide().parents('.mainContent').append(str1);
            $.modal.loading("数据加载中，请稍后...");

            $('.mainContent iframe:visible').load(function () {
                $.modal.closeLoading();
            });

            // 添加选项卡
            $('.menuTabs .page-tabs-content').append(str);
            scrollToTab($('.menuTab.active'));
        }
        return false;
    }

    //滚动到指定选项卡
    function scrollToTab(element) {
        let marginLeftVal = calSumWidth($(element).prevAll()),
            marginRightVal = calSumWidth($(element).nextAll());
        // 可视区域非tab宽度
        let tabOuterWidth = calSumWidth($(".content-tabs").children().not(".menuTabs"));
        //可视区域tab宽度
        let visibleWidth = $(".content-tabs").outerWidth(true) - tabOuterWidth;
        //实际滚动宽度
        let scrollVal = 0;
        if ($(".page-tabs-content").outerWidth() < visibleWidth) {
            scrollVal = 0;
        } else if (marginRightVal <= (visibleWidth - $(element).outerWidth(true) - $(element).next().outerWidth(true))) {
            if ((visibleWidth - $(element).next().outerWidth(true)) > marginRightVal) {
                scrollVal = marginLeftVal;
                let tabElement = element;
                while ((scrollVal - $(tabElement).outerWidth()) > ($(".page-tabs-content").outerWidth() - visibleWidth)) {
                    scrollVal -= $(tabElement).prev().outerWidth();
                    tabElement = $(tabElement).prev();
                }
            }
        } else if (marginLeftVal > (visibleWidth - $(element).outerWidth(true) - $(element).prev().outerWidth(true))) {
            scrollVal = marginLeftVal - $(element).prev().outerWidth(true);
        }
        $('.page-tabs-content').animate({
            marginLeft: 0 - scrollVal + 'px'
        }, "fast");
    }

    //查看左侧隐藏的选项卡
    function scrollTabLeft() {
        let marginLeftVal = Math.abs(parseInt($('.page-tabs-content').css('margin-left')));
        // 可视区域非tab宽度
        let tabOuterWidth = calSumWidth($(".content-tabs").children().not(".menuTabs"));
        //可视区域tab宽度
        let visibleWidth = $(".content-tabs").outerWidth(true) - tabOuterWidth;
        //实际滚动宽度
        let scrollVal = 0;
        if (($(".page-tabs-content").width()) < visibleWidth) {
            return false;
        } else {
            let tabElement = $(".menuTab:first");
            let offsetVal = 0;
            while ((offsetVal + $(tabElement).outerWidth(true)) <= marginLeftVal) { //找到离当前tab最近的元素
                offsetVal += $(tabElement).outerWidth(true);
                tabElement = $(tabElement).next();
            }
            offsetVal = 0;
            if (calSumWidth($(tabElement).prevAll()) > visibleWidth) {
                while ((offsetVal + $(tabElement).outerWidth(true)) < (visibleWidth) && tabElement.length > 0) {
                    offsetVal += $(tabElement).outerWidth(true);
                    tabElement = $(tabElement).prev();
                }
                scrollVal = calSumWidth($(tabElement).prevAll());
            }
        }
        $('.page-tabs-content').animate({
            marginLeft: 0 - scrollVal + 'px'
        }, "fast");
    }

    //查看右侧隐藏的选项卡
    function scrollTabRight() {
        let marginLeftVal = Math.abs(parseInt($('.page-tabs-content').css('margin-left')));
        // 可视区域非tab宽度
        let tabOuterWidth = calSumWidth($(".content-tabs").children().not(".menuTabs"));
        //可视区域tab宽度
        let visibleWidth = $(".content-tabs").outerWidth(true) - tabOuterWidth;
        //实际滚动宽度
        let scrollVal = 0;
        if ($(".page-tabs-content").width() < visibleWidth) {
            return false;
        } else {
            let tabElement = $(".menuTab:first");
            let offsetVal = 0;
            while ((offsetVal + $(tabElement).outerWidth(true)) <= marginLeftVal) { //找到离当前tab最近的元素
                offsetVal += $(tabElement).outerWidth(true);
                tabElement = $(tabElement).next();
            }
            offsetVal = 0;
            while ((offsetVal + $(tabElement).outerWidth(true)) < (visibleWidth) && tabElement.length > 0) {
                offsetVal += $(tabElement).outerWidth(true);
                tabElement = $(tabElement).next();
            }
            scrollVal = calSumWidth($(tabElement).prevAll());
            if (scrollVal > 0) {
                $('.page-tabs-content').animate({
                    marginLeft: 0 - scrollVal + 'px'
                }, "fast");
            }
        }
    }

    // 关闭选项卡菜单
    function closeTab() {
        let closeTabId = $(this).parents('.menuTab').data('id');
        let currentWidth = $(this).parents('.menuTab').width();
        let panelUrl = $(this).parents('.menuTab').data('panel');
        // 当前元素处于活动状态
        if ($(this).parents('.menuTab').hasClass('active')) {

            // 当前元素后面有同辈元素，使后面的一个元素处于活动状态
            if ($(this).parents('.menuTab').next('.menuTab').size()) {

                let activeId = $(this).parents('.menuTab').next('.menuTab:eq(0)').data('id');
                $(this).parents('.menuTab').next('.menuTab:eq(0)').addClass('active');

                $('.mainContent .xframe').each(function () {
                    if ($(this).data('id') == activeId) {
                        $(this).show().siblings('.xframe').hide();
                        return false;
                    }
                });

                let marginLeftVal = parseInt($('.page-tabs-content').css('margin-left'));
                if (marginLeftVal < 0) {
                    $('.page-tabs-content').animate({
                        marginLeft: (marginLeftVal + currentWidth) + 'px'
                    }, "fast");
                }

                //  移除当前选项卡
                $(this).parents('.menuTab').remove();

                // 移除tab对应的内容区
                $('.mainContent .xframe').each(function () {
                    if ($(this).data('id') == closeTabId) {
                        $(this).remove();
                        return false;
                    }
                });
            }

            // 当前元素后面没有同辈元素，使当前元素的上一个元素处于活动状态
            if ($(this).parents('.menuTab').prev('.menuTab').size()) {
                let activeId = $(this).parents('.menuTab').prev('.menuTab:last').data('id');
                $(this).parents('.menuTab').prev('.menuTab:last').addClass('active');
                $('.mainContent .xframe').each(function () {
                    if ($(this).data('id') == activeId) {
                        $(this).show().siblings('.xframe').hide();
                        return false;
                    }
                });

                //  移除当前选项卡
                $(this).parents('.menuTab').remove();

                // 移除tab对应的内容区
                $('.mainContent .xframe').each(function () {
                    if ($(this).data('id') == closeTabId) {
                        $(this).remove();
                        return false;
                    }
                });

                if ($.common.isNotEmpty(panelUrl)) {
                    $('.menuTab[data-id="' + panelUrl + '"]').addClass('active').siblings('.menuTab').removeClass('active');
                    $('.mainContent .xframe').each(function () {
                        if ($(this).data('id') == panelUrl) {
                            $(this).show().siblings('.xframe').hide();
                            return false;
                        }
                    });
                }
            }
        } else {  // 当前元素不处于活动状态
            //  移除当前选项卡
            $(this).parents('.menuTab').remove();

            // 移除相应tab对应的内容区
            $('.mainContent .xframe').each(function () {
                if ($(this).data('id') == closeTabId) {
                    $(this).remove();
                    return false;
                }
            });
        }
        scrollToTab($('.menuTab.active'));
        return false;
    }

    // 点击选项卡菜单
    function activeTab() {
        if (!$(this).hasClass('active')) {
            let currentId = $(this).data('id');
            // 显示tab对应的内容区
            $('.mainContent .xframe').each(function () {
                if ($(this).data('id') == currentId) {
                    $(this).show().siblings('.xframe').hide();
                    return false;
                }
            });
            $(this).addClass('active').siblings('.menuTab').removeClass('active');
            scrollToTab(this);
        }
    }

    // 刷新iframe
    function refreshTab() {
        let currentId = $('.page-tabs-content').find('.active').attr('data-id');
        let target = $('.xframe[data-id="' + currentId + '"]');
        let url = target.attr('src');
        target.attr('src', url).ready();
    }

    //计算元素集合的总宽度
    function calSumWidth(elements) {
        let width = 0;
        $(elements).each(function () {
            width += $(this).outerWidth(true);
        });
        return width;
    }

    // 设置锚点
    function setIframeUrl(href) {
        if ($.common.equals("history", mode)) {
            storage.set('publicPath', href);
        } else {
            let nowUrl = window.location.href;
            let newUrl = nowUrl.substring(0, nowUrl.indexOf("#"));
            window.location.href = newUrl + "#" + href;
        }
    }

    // 激活指定选项卡
    function setActiveTab(element) {
        if (!$(element).hasClass('active')) {
            let currentId = $(element).data('id');
            // 显示tab对应的内容区
            $('.xframe').each(function () {
                if ($(this).data('id') == currentId) {
                    $(this).show().siblings('.xframe').hide();
                }
            });
            $(element).addClass('active').siblings('.menuTab').removeClass('active');
            scrollToTab(element);
        }
    }

    // 关闭当前选项卡
    function tabCloseCurrent() {
        $('.page-tabs-content').find('.active i').trigger("click");
    }

    //关闭其他选项卡
    function tabCloseOther() {
        $('.page-tabs-content').children("[data-id]").not(":first").not(".active").each(function () {
            $('.xframe[data-id="' + $(this).data('id') + '"]').remove();
            $(this).remove();
        });
        $('.page-tabs-content').css("margin-left", "0");
    }

    // 关闭全部选项卡
    function tabCloseAll() {
        $('.page-tabs-content').children("[data-id]").not(":first").each(function () {
            $('.xframe[data-id="' + $(this).data('id') + '"]').remove();
            $(this).remove();
        });
        $('.page-tabs-content').children("[data-id]:first").each(function () {
            $('.xframe[data-id="' + $(this).data('id') + '"]').show();
            $(this).addClass("active");
        });
        $('.page-tabs-content').css("margin-left", "0");
    }

    // 右键菜单实现
    $.contextMenu({
        selector: ".menuTab",
        trigger: 'right',
        autoHide: true,
        items: {
            "close_current": {
                name: "关闭当前",
                icon: "fa-close",
                callback: function (key, opt) {
                    opt.$trigger.find('i').trigger("click");
                }
            },
            "close_other": {
                name: "关闭其他",
                icon: "fa-times-circle-o",
                callback: function (key, opt) {
                    setActiveTab(this);
                    tabCloseOther();
                }
            },
            "close_left": {
                name: "关闭左侧",
                icon: "fa-reply",
                callback: function (key, opt) {
                    setActiveTab(this);
                    this.prevAll('.menuTab').not(":last").each(function () {
                        if ($(this).hasClass('active')) {
                            setActiveTab(this);
                        }
                        $('.xframe[data-id="' + $(this).data('id') + '"]').remove();
                        $(this).remove();
                    });
                    $('.page-tabs-content').css("margin-left", "0");
                }
            },
            "close_right": {
                name: "关闭右侧",
                icon: "fa-share",
                callback: function (key, opt) {
                    setActiveTab(this);
                    this.nextAll('.menuTab').each(function () {
                        $('.menuTab[data-id="' + $(this).data('id') + '"]').remove();
                        $(this).remove();
                    });
                }
            },
            "close_all": {
                name: "全部关闭",
                icon: "fa-times-circle",
                callback: function (key, opt) {
                    tabCloseAll();
                }
            },
            "refresh": {
                name: "刷新页面",
                icon: "fa-refresh",
                callback: function (key, opt) {
                    setActiveTab(this);
                    let target = $('.xframe[data-id="' + this.data('id') + '"]');
                    let url = target.attr('src');
                    $.modal.loading("数据加载中，请稍后...");
                    target.attr('src', url).load(function () {
                        $.modal.closeLoading();
                    });
                }
            }
        }
    })

});