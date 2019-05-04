package config

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

func Update(c *gin.Context) {
	log.Info("iOS Config Update function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	var r UpdateRequest
	if err := c.Bind(&r); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	err := r.Config.Update()
	if err != nil {
		SendError(c, err, nil, err.Error())
		return
	}

	// 同步 Redis 缓存
	str, err := json.Marshal(r.Config)
	if err != nil {
		SendError(c, err, nil, err.Error())
		return
	}

	err = model.DB.Redis.Set("iosConfig", string(str), 0).Err()
	if err != nil {
		SendError(c, err, nil, err.Error())
		return
	}

	util.SendNotification("iOS元信息配置已更新")

	rsp := UpdateResponse{}

	SendResponse(c, nil, rsp)
}
