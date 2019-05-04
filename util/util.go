package util

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/teris-io/shortid"
)

func GenShortId() (string, error) {
	return shortid.Generate()
}

func GetReqID(c *gin.Context) string {
	v, ok := c.Get("X-Request-Id")
	if !ok {
		return ""
	}
	if requestID, ok := v.(string); ok {
		return requestID
	}
	return ""
}

// SendNotification ... 发送钉钉通知
func SendNotification(text string) {
	// debug 环境下不发送通知
	if viper.GetString("runmode") != "debug" {
		log.Println("发送警报", viper.GetString("runmode"))
		message := map[string]interface{}{
			"msgtype": "text",
			"text": map[string]string{
				"content": text,
			},
		}

		bytesRepresentation, err := json.Marshal(message)
		if err != nil {
			log.Println(err)
		}

		_, err = http.Post("https://oapi.dingtalk.com/robot/send?access_token=0fc384d57235fdb1cc6dfa83408d8154507d0699f3649589dd5a7898012ee690", "application/json", bytes.NewBuffer(bytesRepresentation))
		if err != nil {
			log.Println(err)
		}
	}
}
