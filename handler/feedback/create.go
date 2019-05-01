package feedback

import (
	. "github.com/asynccnu/mana_service_v2/handler"
	"github.com/asynccnu/mana_service_v2/model"
	"github.com/asynccnu/mana_service_v2/pkg/errno"
	"github.com/asynccnu/mana_service_v2/util"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

func Create(c *gin.Context) {
	log.Info("feedback create function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	var r model.FeedbackItem
	if err := c.Bind(&r); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	err := r.Create()
	if err != nil {
		SendError(c, err, nil, err.Error())
		return
	}

	rsp := CreateResponse{}

	SendResponse(c, nil, rsp)
}
