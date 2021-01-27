package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
)

type AuthController struct {
}

func NewAuthController() *AuthController {
	return &AuthController{}
}

func (this *AuthController) Name() string {
	return "AuthController"
}

func (this *AuthController) Auth(ctx *gin.Context) goft.Json {
	return gin.H{"message": "ok"}
}

func (this *AuthController) Build(goft *goft.Goft) {
	goft.Handle("ANY", "/auth", this.Auth)
	gin.New().Any()
}
