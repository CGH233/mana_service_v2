package banner

import (
	"encoding/json"

	. "github.com/asynccnu/mana_service_v2/handler"
	"github.com/asynccnu/mana_service_v2/model"
	"github.com/lexkong/log"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	var config *model.BannerConfig

	// 从 Redis 获取
	val, err := model.DB.Redis.Get("bannerConfig").Result()

	// fallback 到 db
	if err != nil {
		log.Info("Redis error, fallback banner config from mongo")
		config, err = model.GetBanner()
		if err != nil {
			SendError(c, err, nil, err.Error())
			return
		}
		SendResponse(c, err, &config.List)
		return
	}

	err = json.Unmarshal([]byte(val), &config)
	if err != nil {
		SendError(c, err, nil, err.Error())
		return
	}
	SendResponse(c, nil, &config.List)
}
