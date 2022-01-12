package v1

import (
	"gohub/app/models/user"
	"gohub/app/requests"
	"gohub/pkg/auth"
	"gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

type UsersController struct {
	BaseAPIController
}

// CurrentUser 当前登录用户信息
func (ctrl *UsersController) CurrentUser(c *gin.Context) {
	userModel := auth.CurrentUser(c)
	response.Data(c, userModel)
}

// Index 所有用户
func (ctrl *UsersController) Index(c *gin.Context) {
	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}

	data, pager := user.Paginate(c, 10)
	response.JSON(c, gin.H{
		"data":  data,
		"pager": pager,
	})
}

func (ctrl *UsersController) UpdateProfile(c *gin.Context) {

	request := requests.UserUpdateProfileRequest{}
	if ok := requests.Validate(c, &request, requests.UserUpdateProfile); !ok {
		return
	}

	currentUser := auth.CurrentUser(c)
	currentUser.Name = request.Name
	currentUser.City = request.City
	currentUser.Introduction = request.Introduction
	rowsAffected := currentUser.Save()
	if rowsAffected > 0 {
		response.Data(c, currentUser)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}
