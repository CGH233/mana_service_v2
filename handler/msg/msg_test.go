package msg

import (
	"encoding/json"
	"net/http"
	"os"
	"testing"

	"github.com/asynccnu/mana_service_v2/config"
	"github.com/asynccnu/mana_service_v2/model"
	"github.com/asynccnu/mana_service_v2/router/middleware"
	"github.com/asynccnu/mana_service_v2/util"

	"github.com/gin-gonic/gin"
)

var (
	g *gin.Engine
)

var router = getRouter(true)

func TestMain(m *testing.M) {

	// init config
	if err := config.Init(""); err != nil {
		panic(err)
	}
	// init db
	model.DB.Init()
	defer model.DB.Close()

	os.Exit(m.Run())
}

func TestCreate(t *testing.T) {
	g := router
	uri := "/api/msg"
	u := model.MessageItem{
		OS:      "iOS",
		Page:    "com.muxistudio.ele",
		Msg:     "大促销",
		Detail:  "foobar",
		Time:    "foobar",
		Version: "foobar",
	}
	jsonByte, err := json.Marshal(u)
	if err != nil {
		t.Errorf("Test Error: %s", err.Error())
	}
	w := util.PerformRequestWithBody(http.MethodPost, g, uri, jsonByte, true)
	result := w.Result()
	if result.StatusCode != http.StatusOK {
		t.Errorf("Test Error: StatusCode Error:%d", result.StatusCode)
	}
}

// 测试获取 Msg
func TestGet(t *testing.T) {
	g := router
	uri := "/api/msg?os=iOS&page=com.muxistudio.ele"
	w := util.PerformRequest(http.MethodGet, g, uri, false)
	result := w.Result()

	if result.StatusCode != http.StatusOK {
		t.Errorf("Test Error: StatusCode Error:%d", result.StatusCode)
	}
}

// 测试 Redis 缓存失效，fallback 到 DB
func TestGetFallback(t *testing.T) {
	g := router

	err := model.GetRedis().Del("msg-iOS-com.muxistudio.ele").Err()
	if err != nil {
		t.Errorf("Test Error: %s", err.Error())
	}

	uri := "/api/msg?os=iOS&page=com.muxistudio.ele"
	w := util.PerformRequest(http.MethodGet, g, uri, false)
	result := w.Result()

	if result.StatusCode != http.StatusOK {
		t.Errorf("Test Error: StatusCode Error:%d", result.StatusCode)
	}
}

func TestUpdate(t *testing.T) {
	g := router
	uri := "/api/msg"
	u := model.MessageItem{
		OS:      "iOS",
		Page:    "com.muxistudio.ele",
		Msg:     "跳楼价",
		Detail:  "foobar",
		Time:    "foobar",
		Version: "foobar",
	}
	jsonByte, err := json.Marshal(u)
	if err != nil {
		t.Errorf("Test Error: %s", err.Error())
	}
	w := util.PerformRequestWithBody(http.MethodPost, g, uri, jsonByte, true)
	result := w.Result()
	if result.StatusCode != http.StatusOK {
		t.Errorf("Test Error: StatusCode Error:%d", result.StatusCode)
	}
}

// 测试删除 Msg
func TestDelete(t *testing.T) {
	g := router
	r := DeleteRequest{
		OS:   "iOS",
		Page: "com.muxistudio.ele",
	}
	jsonByte, err := json.Marshal(r)
	if err != nil {
		t.Errorf("Test Error: %s", err.Error())
	}

	uri := "/api/msg"
	w := util.PerformRequestWithBody(http.MethodDelete, g, uri, jsonByte, true)
	result := w.Result()

	if result.StatusCode != http.StatusOK {
		t.Errorf("Test Error: StatusCode Error:%d", result.StatusCode)
	}
}

// Helper function to create a router during testing
func getRouter(withRouter bool) *gin.Engine {
	g = gin.New()
	if withRouter {
		loadRouters(
			// Cores.
			g,
			middleware.RequestId(),
		)
	}
	return g
}

// Load loads the middlewares, routes, handlers about Test
func loadRouters(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Middlewares.
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	u := g.Group("/api")

	{
		u.GET("/msg", Get)
		u.POST("/msg", Create)
		u.PUT("/msg", Update)
		u.DELETE("/msg", Delete)
	}

	return g
}
