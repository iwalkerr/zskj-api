initWebSocket()

// 收到消息处理
function onmessage(e) {
    count = 0
    // 处理消息
    let data = JSON.parse(e.data)
    switch (data.grp_id) {
        case "backend-110":
            // console.log(data)
            break
    }
}

let lock = false; count = 0
// websocket
function initWebSocket() {
    // 浏览器提供了WebSocket类型，在Firefox中为MozWebSocket
    if (window.WebSocket || window.MozWebSocket) {
        conn()
    } else {
        console.log("浏览器不支持websocket")
    }
}
function conn() {
    //可以看到客户端JS，很容易的就通过WebSocket函数建立了一个与服务器的连接sock，当握手成功后，会触发WebScoket对象的onopen事件，告诉客户端连接已经成功建立。客户端一共绑定了四个事件。
    let ws = new WebSocket($("#websocketConn").val())
    let userId = $("#userId").val()
    // 收到消息后触发
    ws.onmessage = onmessage
    //建立连接后触发
    ws.onopen = function () {
        let obj = { from_user: userId, to_user: userId, grp_id: "", content: "ping", content_type: 1 }
        obj = JSON.stringify(obj)
        ws.send(obj)
    };
    // 关闭连接时候触发
    ws.onclose = function (e) {
        reconnect()
        console.log('关闭websocket连接！')
    };
    //发生错误的时候触发
    ws.onerror = function (e) { reconnect() }
}

// 每隔1秒重连
function reconnect() {
    if (lock) return; lock = true
    setTimeout(function () {       // 没连接上会一直重连，设置延迟避免请求过多
        lock = false
        // 重连50次后还没有连上，则关闭websocket
        if (count < 5) { initWebSocket() } else count = 0
        count++
    }, 1000)
}