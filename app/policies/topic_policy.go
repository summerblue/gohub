// Package policies 用户授权
package policies

import (
	"gohub/app/models/topic"
	"gohub/pkg/auth"

	"github.com/gin-gonic/gin"
)

func CanModifyTopic(c *gin.Context, _topic topic.Topic) bool {
	return auth.CurrentUID(c) == _topic.UserID
}
