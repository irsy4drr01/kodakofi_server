package handlers

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kodakofidev/kodakofi_server/internal/models"
	"github.com/kodakofidev/kodakofi_server/internal/repositories"
	"github.com/kodakofidev/kodakofi_server/pkg"
)

type OrderHandlers struct {
	repo repositories.OrderRepoInterface
}

func NewOrder(repo repositories.OrderRepoInterface) *OrderHandlers {
	return &OrderHandlers{repo: repo}
}

func (h *OrderHandlers) PostOrderHandler(ctx *gin.Context) {
	responder := models.NewResponse(ctx)

	order := models.CreateOrderRequest{}

	if err := ctx.ShouldBindJSON(&order); err != nil {
		log.Println("Validation Error:", err.Error())
		responder.BadRequest("Validation Error", "Invalid request payload format")
		return
	}

	createOrder, err := h.repo.CreateOrder(ctx, &order)
	if err != nil {
		log.Println("Create order error:", err.Error())
		responder.InternalServerError("Internal Server Error", "Failed to create order")
		return
	}

	responder.Created("Order created successfully", createOrder)
}

// handlers get history order
func (h *OrderHandlers) GetHistoryOrders(ctx *gin.Context) {
	claims, _ := ctx.Get("payloads")
	userClaims := claims.(*pkg.Claims)

	response := models.NewResponse(ctx)

	// tangkap query
	pageQ := ctx.Query("page")
	statusQ := ctx.Query("status")
	var offset int
	var pageQInt int
	if pageQ != "" {
		pageQNum, err := strconv.Atoi(pageQ)
		if err != nil {
			log.Println("Page query conversion error:", err.Error())
			response.InternalServerError("Internal Server Error", "A server error occurred")
			return
		}
		pageQInt += pageQNum
	}

	if pageQInt == 1 {
		offset = 0
	} else if pageQInt == 0 {
		offset = -1
	} else {
		offset = pageQInt*4 - 4
	}

	log.Println("offset", offset)
	log.Println("statusQ", statusQ)

	result, err := h.repo.GetHistoryOrders(ctx, offset, statusQ, userClaims.Uuid)
	if err != nil {
		log.Println("GetHistoryOrders error:", err.Error())
		response.InternalServerError("Internal Server Error", "A server error occurred while fetching order history")
		return
	}
	println(len(result))
	if len(result) == 0 {
		log.Println("No history orders found for user:", userClaims.Uuid)
		response.NotFound("Not Found", "No order history available")
		return
	}

	response.Success("success", result)
}

func (h *OrderHandlers) FetchDetailOrderByUser(ctx *gin.Context) {
	responder := models.NewResponse(ctx)

	// Ambil klaim dari middleware JWT
	claims, exists := ctx.Get("payloads")
	if !exists {
		responder.Unauthorized("Unauthorized", "Missing token claims")
		return
	}

	userClaims, ok := claims.(*pkg.Claims)
	if !ok {
		responder.Unauthorized("Unauthorized", "Invalid token claims")
		return
	}

	userID := userClaims.Uuid

	// Ambil order_id dari path parameter
	orderIDParam := ctx.Param("order_id")
	orderID, err := strconv.Atoi(orderIDParam)
	if err != nil {
		responder.BadRequest("Bad Request", "Invalid Fetch Order")
		return
	}

	// Panggil repository
	result, err := h.repo.GetDetailOrderByUser(ctx.Request.Context(), userID, orderID)
	if err != nil {
		responder.InternalServerError("Internal Server Error", "Failed to get order detail")
		return
	}

	responder.Success("Order detail fetched successfully", result)
}
