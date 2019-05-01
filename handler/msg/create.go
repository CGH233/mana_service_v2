package msg

import (
	"encoding/json"

	. "github.com/asynccnu/mana_service_v2/handler"
	"github.com/asynccnu/mana_service_v2/model"
	"github.com/asynccnu/mana_service_v2/pkg/errno"
	"github.com/asynccnu/mana_service_v2/util"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

// Create ... 创建消息
func Create(c *gin.Context) {
	log.Info("Msg Create function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	var r model.MessageItem
	if err := c.Bind(&r); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	err := r.Update()
	if err != nil {
		SendError(c, err, nil, err.Error())
		return
	}

	// 同步 Redis 缓存
	str, err := json.Marshal(r)
	if err != nil {
		SendError(c, err, nil, err.Error())
		return
	}

	err = model.DB.Redis.Set("msg-"+r.OS+"-"+r.Page, string(str), 0).Err()
	if err != nil {
		SendError(c, err, nil, err.Error())
		return
	}

	rsp := CreateResponse{}

	SendResponse(c, nil, rsp)
}
