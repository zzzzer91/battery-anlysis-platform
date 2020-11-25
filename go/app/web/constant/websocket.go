package constant

import "time"

const (
	WsReadBufferSize  = 4096
	WsWriteBufferSize = 4096
	WsSendInterval    = time.Second * 3
)
