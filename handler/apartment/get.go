package apartment

import (
	. "github.com/asynccnu/mana_service_v2/handler"

	"github.com/gin-gonic/gin"
)

var apartmentInfoList = []ApartmentInfoItem{
	ApartmentInfoItem{
		Apartment: "学生事务大厅",
		Phone:     []string{"67865591"},
		Place:     "文华公书林（老图书馆）3楼",
	},
	ApartmentInfoItem{
		Apartment: "校医院",
		Phone:     []string{"67867176"},
		Place:     "九号楼侧边下阶梯右转处",
	},
	ApartmentInfoItem{
		Apartment: "水电费缴纳处",
		Phone:     []string{"67861701"},
		Place:     "东南门附近的水电修建服务中心内",
	},
	ApartmentInfoItem{
		Apartment: "校园卡管理中心",
		Phone:     []string{"67868524"},
		Place:     "田家炳楼侧边阶梯旁",
	},
	ApartmentInfoItem{
		Apartment: "教务处",
		Phone:     []string{"67868057"},
		Place:     "行政楼2楼",
	},
	ApartmentInfoItem{
		Apartment: "学生资助中心",
		Phone:     []string{"67865161"},
		Place:     "学生事务大厅10号窗口",
	},
	ApartmentInfoItem{
		Apartment: "党校办公室",
		Phone:     []string{"67868011"},
		Place:     "行政楼副楼",
	},
	ApartmentInfoItem{
		Apartment: "素质教育办公室",
		Phone:     []string{"67868057"},
		Place:     "暂无信息",
	},
	ApartmentInfoItem{
		Apartment: "档案馆",
		Phone:     []string{"67867198"},
		Place:     "科学会堂一楼",
	},
	ApartmentInfoItem{
		Apartment: "国际交流事务办",
		Phone:     []string{"67861299"},
		Place:     "法学院前面的路口左拐（有路标）",
	},
	ApartmentInfoItem{
		Apartment: "心理咨询中心",
		Phone:     []string{"67868274"},
		Place:     "文华公书林（老图书馆）3楼",
	},
	ApartmentInfoItem{
		Apartment: "保卫处",
		Phone:     []string{"67868110"},
		Place:     "化学逸夫楼斜对面",
	},
}

func Get(c *gin.Context) {
	SendResponse(c, nil, &apartmentInfoList)
}
