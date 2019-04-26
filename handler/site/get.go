package site

import (
	. "github.com/asynccnu/mana_service_v2/handler"

	"github.com/gin-gonic/gin"
)

var siteList = []SiteItem{
	SiteItem{
		Site: "网上一站式服务门户",
		URL:  "http://one.ccnu.edu.cn",
	},
	SiteItem{
		Site: "教务管理系统（本科生）",
		URL:  "http://xk.ccnu.edu.cn/ssoserver/login?ywxt=jw&url=xtgl/index_initMenu.html",
	},
	SiteItem{
		Site: "素质课程管理平台",
		URL:  "http://122.204.187.9/jwglxt/xtgl/dl_loginForward.html",
	},
	SiteItem{
		Site: "图书馆首页",
		URL:  "http://lib.ccnu.edu.cn",
	},
	SiteItem{
		Site: "教务处",
		URL:  "http://jwc.ccnu.edu.cn/",
	},
	SiteItem{
		Site: "尔雅课",
		URL:  "http://ccnu.benke.chaoxing.com",
	},
	SiteItem{
		Site: "全国大学四六级",
		URL:  "http://www.cet.edu.cn",
	},
	SiteItem{
		Site: "微课程平台",
		URL:  "http://wk.ccnu.edu.cn/hsd-web/index.do",
	},
	SiteItem{
		Site: "校医院",
		URL:  "http://hosp.ccnu.edu.cn",
	},
	SiteItem{
		Site: "网上自主缴费平台",
		URL:  "http://218.199.196.90",
	},
	SiteItem{
		Site: "图书馆研习室预约",
		URL:  "http://202.114.34.12:8088/reserveSystem/readerIndex.html",
	},
	SiteItem{
		Site: "志愿者管理系统",
		URL:  "http://zyz.ccnu.edu.cn",
	},
}

func Get(c *gin.Context) {
	SendResponse(c, nil, &siteList)
}
