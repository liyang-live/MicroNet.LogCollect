package controller

import (
	"MicroNet.LogCollect/common"
	"MicroNet.LogCollect/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func GetLogRecord(c *gin.Context)  {
	c.JSON(http.StatusOK,
		 models.LogModel{Id:"1",System:"测试系统",Module:"测试模块",FuncName:"测试方法",
		                 FuncParameter:"测试参数",LogTime:common.TimFormat(time.Now().String(),"yyyy-MM-dd HH:mm:ss"),
			             LogContent:"测试内容"})
}

func RecordLog(c *gin.Context) {

	var logger models.LogModel

	go writeLog(logger)
}

func writeLog(model models.LogModel) {
	log.Print(model)
}
