package file

import (
	"battery-analysis-platform/app/web/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"path/filepath"
)

func fileResponse(c *gin.Context, s service.FileServicer) {
	if path, err := s.Do(); err != nil {
		// 返回 500 错误码，并且会在控制台打印错误
		c.AbortWithError(500, err)
	} else {
		_, filename := filepath.Split(path)
		c.Writer.Header().Set("content-disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))
		// 代表可能是任何类型
		c.Writer.Header().Add("Content-Type", "application/octet-stream")
		// 这里必须传入完整路径，因为内部调用的函数还会有一次分割
		c.File(path)
	}
}
