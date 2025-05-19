package middlewares

import (
	"log"
	"slices"

	"github.com/gin-gonic/gin"
	"github.com/kodakofidev/kodakofi_server/internal/models"
	"github.com/kodakofidev/kodakofi_server/pkg"
)

func (m *Middleware) AccsessGate(allowedRole ...string) func(*gin.Context) {
	return func(ctx *gin.Context) {
		responder := models.NewResponse(ctx)

		payloads, exits := ctx.Get("payloads")
		if !exits {
			log.Println("Unauthorized: input payloads in context")
			responder.Unauthorized("Unauthorized", "Please login first!")
			return
		}
		userPayload, ok := payloads.(*pkg.Claims)
		if !ok {
			log.Println("Unauthorized: payloads is not of type *pkg.Claims")
			responder.Unauthorized("Unauthorized", "Your login identity is malformed, please login again!")
			return
		}
		if !slices.Contains(allowedRole, userPayload.Role) {
			log.Printf("Forbidden access: role '%s' is not allowed\n", userPayload.Role)
			responder.Forbidden("Forbidden", "You do not have permission to access this resource")
			return
		}
		ctx.Next()
	}
}
