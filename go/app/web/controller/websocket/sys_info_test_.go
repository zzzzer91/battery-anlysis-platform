package websocket

import (
	"battery-analysis-platform/app/web/model"
	"github.com/gorilla/websocket"
	"testing"
)

// TODO
func TestSysInfo(t *testing.T) {
	url := "ws://localhost:5000/websocket/v1/sys-info"
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		t.Fatal(err)
	}
	var jd model.SysInfo
	if err = conn.ReadJSON(&jd); err != nil {
		t.Fatal(err)
	} else {
		t.Log(jd)
	}
}
