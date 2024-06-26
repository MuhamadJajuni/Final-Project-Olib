package controller

import (
	"final-project-olib/middleware"
	"final-project-olib/model/dto"
	commonresponse "final-project-olib/model/dto/common_response"
	"final-project-olib/usecase"

	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authUc         usecase.AuthUseCase
	routerGroup    *gin.RouterGroup
	authMiddleware middleware.AuthMiddleware
}

func (a *AuthController) loginHandler(ctx *gin.Context) {
	var payload dto.AuthReqDto
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		commonresponse.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}
	token, err := a.authUc.Login(payload)
	if err != nil {
		commonresponse.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}
	commonresponse.SendSingleResponse(ctx, token, "Success Login")

}
func (a *AuthController) loginHandlerAdmin(ctx *gin.Context) {
	var payload dto.AuthReqDto
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		commonresponse.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}
	token, err := a.authUc.LoginAdmin(payload)
	if err != nil {
		commonresponse.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}
	commonresponse.SendSingleResponse(ctx, token, "Success Login")

}

func (a *AuthController) Route() {
	a.routerGroup.POST("/auth/login", a.loginHandler)
	a.routerGroup.POST("/auth/login/admin", a.loginHandlerAdmin)
}

func NewAuthController(authUc usecase.AuthUseCase, rg *gin.RouterGroup, authMiddleware middleware.AuthMiddleware) *AuthController {
	return &AuthController{authUc: authUc,
		authMiddleware: authMiddleware, routerGroup: rg}
}
