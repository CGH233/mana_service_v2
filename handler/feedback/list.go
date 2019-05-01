package feedback

import (
	. "github.com/asynccnu/mana_service_v2/handler"
	"github.com/asynccnu/mana_service_v2/model"
	"github.com/asynccnu/mana_service_v2/pkg/errno"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {

	var r ListRequest
	if err := c.ShouldBindQuery(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	list, count, err := model.ListFeedback(r.Offset, r.Limit)
	if err != nil {
		SendError(c, err, nil, err.Error())
		return
	}

	SendResponse(c, nil, ListResponse{
		TotalCount: count,
		List:       list,
	})
}
