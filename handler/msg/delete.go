package msg

import (
	. "github.com/asynccnu/mana_service_v2/handler"
	"github.com/asynccnu/mana_service_v2/model"
	"github.com/asynccnu/mana_service_v2/pkg/errno"
	"github.com/asynccnu/mana_service_v2/util"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"

	"github.com/gin-gonic/gin"
)

// Delete ... 删除消息
func Delete(c *gin.Context) {
	log.Info("Msg delete function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	var r DeleteRequest
	if err := c.Bind(&r); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	err := model.DeleteMsg(r.OS, r.Page)
	if err != nil {
		SendError(c, err, nil, err.Error())
		return
	}

	rsp := DeleteResponse{}

	SendResponse(c, nil, rsp)
}
