package middleware

import (
	"github.com/gin-gonic/gin"
	"kBase-permissions/src/common"
	"strconv"
)

const (
	GET_USERID_KEY = "userId"
)

type AuthMiddleware struct {
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (this *AuthMiddleware) OnRequest(ctx *gin.Context) error {
	user := GetUserId(ctx)
	var userStr string
	if user == 1 {
		userStr = "shengyi"
	} else if user == 2 {
		userStr = "lisi"
	}
	access, err := common.E.Enforce(userStr, ctx.Request.RequestURI, ctx.Request.Method)
	if err != nil || !access {
		s := ""
		if err != nil {
			s = err.Error()
		}
		ctx.AbortWithStatusJSON(403, gin.H{"message": "forbidden" + s})
	} else {
		ctx.Next()
	}
	return nil
}

func (this *AuthMiddleware) OnResponse(result interface{}) (interface{}, error) {
	return nil, nil
}

func GetUserId(ctx *gin.Context) int {
	userIdStr := ctx.Request.Header.Get(GET_USERID_KEY)
	userId, _ := strconv.Atoi(userIdStr)
	return userId
}
