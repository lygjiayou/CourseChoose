package api

import (
	"CourseChoose/model"
	"CourseChoose/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateMember(c *gin.Context) {
	// 创建成员业务部分
	var req model.CreateMemberRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, model.CreateMemberRespone{
			Code: model.ParamInvalid,
		})
	} else {
		//if req.UserType == model.Admin
		// 参数校验
		if service.CheckCreateMemberParamService(&req) {
			// 校验通过，进行创建成员操作
			c.JSON(http.StatusOK, service.CreateMemberService(&req))
		}else {
			// 校验不通过
			c.JSON(http.StatusOK, model.CreateMemberRespone{
				Code: model.ParamInvalid,
			})
		}
	}
}
