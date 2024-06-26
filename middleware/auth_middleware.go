package middleware

import (
	"final-project-olib/service"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware interface {
	CheckToken() gin.HandlerFunc
}
type authMiddleware struct {
	jwtService service.JwtService
}
type AuthHeader struct {
	Autheader string `header:"Authorization" required:"true"`
}

// CheckToken implements AuthMiddleware.
func (a *authMiddleware) CheckToken() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var header AuthHeader

		ctx.ShouldBindHeader(&header)
		token := strings.Replace(header.Autheader, "Bearer ", "", -1)
		claims, err := a.jwtService.ValidateToken(token)
		if err != nil {
			log.Println(err)
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		ctx.Set("author", claims["authorId"])
	}

}

func NewAuthMiddleware(jwtService service.JwtService) AuthMiddleware {
	return &authMiddleware{jwtService: jwtService}
}
