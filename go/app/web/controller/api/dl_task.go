package api

import (
	"battery-analysis-platform/app/web/service"
	"github.com/gin-gonic/gin"
)

func CreateDlTask(c *gin.Context) {
	s := service.CreateDlTaskService{}
	if err := c.ShouldBindJSON(&s); err != nil {
		c.AbortWithError(500, err)
		return
	}
	jsonResponse(c, &s)
}

func GetDlTaskList(c *gin.Context) {
	s := service.GetDlTaskListService{}
	jsonResponse(c, &s)
}

func GetDlTaskTraningHistory(c *gin.Context) {
	s := service.GetDlTaskTraningHistoryService{
		Id:            c.Param("taskId"),
		ReadFromRedis: false,
	}
	jsonResponse(c, &s)
}

func GetDlEvalResultHistory(c *gin.Context) {
	s := service.GetDlTaskEvalResultService{
		Id: c.Param("taskId"),
	}
	jsonResponse(c, &s)
}

func DeleteDlTask(c *gin.Context) {
	s := service.DeleteDlTaskService{
		Id: c.Param("taskId"),
	}
	jsonResponse(c, &s)
}
