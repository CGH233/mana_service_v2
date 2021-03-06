package handler

import (
	"net/http"

	"github.com/asynccnu/mana_service_v2/util"

	"github.com/asynccnu/mana_service_v2/pkg/errno"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	TraceID string      `json:"traceId"`
}

func SendResponse(c *gin.Context, err error, data interface{}) {
	// 历史原因，200 情况直接返回数据
	c.JSON(http.StatusOK, data)
}

func SendBadRequest(c *gin.Context, err error, data interface{}, cause string) {
	code, message := errno.DecodeErr(err)
	log.Info(message, lager.Data{"X-Request-Id": util.GetReqID(c), "cause": cause})
	c.JSON(http.StatusBadRequest, Response{
		Code:    code,
		Message: message + ": " + cause,
		Data:    data,
		TraceID: util.GetReqID(c),
	})
}

func SendUnauthorized(c *gin.Context, err error, data interface{}, cause string) {
	code, message := errno.DecodeErr(err)
	log.Info(message, lager.Data{"X-Request-Id": util.GetReqID(c), "cause": cause})
	c.JSON(http.StatusUnauthorized, Response{
		Code:    code,
		Message: message + ": " + cause,
		Data:    data,
		TraceID: util.GetReqID(c),
	})
}

func SendError(c *gin.Context, err error, data interface{}, cause string) {
	code, message := errno.DecodeErr(err)
	log.Info(message, lager.Data{"X-Request-Id": util.GetReqID(c), "cause": cause})
	c.JSON(http.StatusInternalServerError, Response{
		Code:    code,
		Message: message + ": " + cause,
		Data:    data,
		TraceID: util.GetReqID(c),
	})
}
