package file

import (
	"battery-analysis-platform/app/web/service"
	"github.com/gin-gonic/gin"
)

func DownloadDlModel(c *gin.Context) {
	s := service.DownloadDlModelService{
		Id: c.Param("taskId"),
	}
	fileResponse(c, &s)
}
