package msg

import (
	"encoding/json"

	. "github.com/asynccnu/mana_service_v2/handler"
	"github.com/asynccnu/mana_service_v2/model"
	"github.com/asynccnu/mana_service_v2/pkg/errno"
	"github.com/lexkong/log"

	"github.com/gin-gonic/gin"
)

// Get ... 获取消息
func Get(c *gin.Context) {
	var r GetRequest
	if err := c.ShouldBindQuery(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	var msg *model.MessageItem

	// 从 Redis 获取
	val, err := model.DB.Redis.Get("msg-" + r.OS + "-" + r.Page).Result()

	// fallback 到 db
	if err != nil {
		log.Info("Redis error, fallback msg from mongo")
		msg, err := model.GetMsg(r.OS, r.Page)
		if err != nil {
			SendError(c, err, nil, err.Error())
			return
		}
		SendResponse(c, err, &msg)
		return
	}

	err = json.Unmarshal([]byte(val), &msg)
	if err != nil {
		SendError(c, err, nil, err.Error())
		return
	}
	SendResponse(c, nil, &msg)
}
