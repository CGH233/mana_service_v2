package banner

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

// 首先插入一条记录，顺便测试 Update 接口
func TestUpdate(t *testing.T) {
	g := router
	uri := "/api/ios/banner"
	u := model.BannerConfig{
		List: []model.BannerItem{{
			Img: "https://xueer.muxixyz.com",
			URL: "https://bar.foo",
			Num: 0,
		}, {
			Img: "https://xueer.muxixyz.com",
			URL: "https://bar.foo",
			Num: 1,
		},
			{
				Img: "https://xueer.muxixyz.com",
				URL: "https://bar.foo",
				Num: 2,
			},
		},
	}
	jsonByte, err := json.Marshal(u)
	if err != nil {
		t.Errorf("Test Error: %s", err.Error())
	}
	w := util.PerformRequestWithBody(http.MethodPut, g, uri, jsonByte, true)
	result := w.Result()
	if result.StatusCode != http.StatusOK {
		t.Errorf("Test Error: StatusCode Error:%d", result.StatusCode)
	}
}

// 测试获取 Config
func TestGet(t *testing.T) {
	g := router
	uri := "/api/ios/banner"
	w := util.PerformRequest(http.MethodGet, g, uri, false)
	result := w.Result()

	if result.StatusCode != http.StatusOK {
		t.Errorf("Test Error: StatusCode Error:%d", result.StatusCode)
	}

	// 读取响应 body 并解析
	var resp []model.BannerItem

	if err := json.Unmarshal([]byte(w.Body.String()), &resp); err != nil {
		t.Errorf("Test error: Prase Response Error:%s", err.Error())
	}
}

// 测试 Redis 缓存失效，fallback 到 DB
func TestGetFallBack(t *testing.T) {
	g := router
	err := model.GetRedis().Del("bannerConfig").Err()
	if err != nil {
		t.Errorf("Test Error: %s", err.Error())
	}
	uri := "/api/ios/banner"
	w := util.PerformRequest(http.MethodGet, g, uri, false)
	result := w.Result()

	if result.StatusCode != http.StatusOK {
		t.Errorf("Test Error: StatusCode Error:%d", result.StatusCode)
	}

	// 读取响应 body 并解析
	var resp []model.BannerItem

	if err := json.Unmarshal([]byte(w.Body.String()), &resp); err != nil {
		t.Errorf("Test error: Prase Response Error:%s", err.Error())
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
		u.GET("/ios/banner", Get)
		u.PUT("/ios/banner", Update)
	}

	return g
}
