package websocket

import (
	"battery-analysis-platform/app/web/constant"
	"github.com/gorilla/websocket"
	"net/http"
)

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  constant.WsReadBufferSize,
	WriteBufferSize: constant.WsWriteBufferSize,
	// 取消 ws 跨域校验，http 头部的 origin 不对，会返回 403，这里禁止这个功能
	CheckOrigin: func(r *http.Request) bool { return true },
}

func upgradeHttpConn(writer http.ResponseWriter, request *http.Request) (*websocket.Conn, error) {
	conn, err := wsUpgrader.Upgrade(writer, request, nil)
	if err != nil {
		return nil, err
	}
	return conn, err
}

// 注意，WebSocket 的关闭是应用层的而不是传输层，
// 关闭的一方会发一个包代表关闭，但 TCP 连接还保持着，
// 收到包的一方也必须回一个关闭包，整个 TCP 连接才会关闭，
// 所以不能把监控 TCP 关闭的方法应用在 WebSocket。
// 监测对端是否关闭的方法，只有读对端
func monitorWsClosed(conn *websocket.Conn) chan struct{} {
	closed := make(chan struct{})
	go func() {
		_, _, _ = conn.ReadMessage()
		closed <- struct{}{}
	}()
	return closed
}
