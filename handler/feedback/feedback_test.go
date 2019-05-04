package feedback

import (
	"encoding/json"
	"net/http"
	"os"
	"testing"
	"time"

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
	uri := "/api/feedback"
	u := model.FeedbackItem{
		Contact:   "127794324",
		Content:   "hhhhhhhhhhhhhhhhhh",
		Timestamp: time.Now().Unix(),
	}
	jsonByte, err := json.Marshal(u)
	if err != nil {
		t.Errorf("Test Error: %s", err.Error())
	}
	// 插入三条
	for i := 0; i < 3; i++ {
		w := util.PerformRequestWithBody(http.MethodPost, g, uri, jsonByte, true)
		result := w.Result()
		if result.StatusCode != http.StatusOK {
			t.Errorf("Test Error: StatusCode Error:%d", result.StatusCode)
		}
	}
}

// 测试获取 Feedback
func TestGetPageOne(t *testing.T) {
	g := router
	uri := "/api/feedback?page=0&limit=1"
	w := util.PerformRequest(http.MethodGet, g, uri, false)
	result := w.Result()

	if result.StatusCode != http.StatusOK {
		t.Errorf("Test Error: StatusCode Error:%d", result.StatusCode)
	}

	// 读取响应 body 并解析
	var resp ListResponse

	if err := json.Unmarshal([]byte(w.Body.String()), &resp); err != nil {
		t.Errorf("Test error: Prase Response Error:%s", err.Error())
	}
}

// 测试获取 Feedback
func TestGetPageTwo(t *testing.T) {
	g := router
	uri := "/api/feedback?page=1&limit=1"
	w := util.PerformRequest(http.MethodGet, g, uri, false)
	result := w.Result()

	if result.StatusCode != http.StatusOK {
		t.Errorf("Test Error: StatusCode Error:%d", result.StatusCode)
	}

	// 读取响应 body 并解析
	var resp ListResponse

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
		u.POST("/feedback", Create)
		u.GET("/feedback", List)
	}

	return g
}
