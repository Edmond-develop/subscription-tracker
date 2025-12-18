package controller

import (
	"github.com/Edmond-develop/subscription-tracker/internal/database"
	"github.com/Edmond-develop/subscription-tracker/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SubscriptionHandler struct {
	service *service.SubscriptionService
}

func NewSubscriptionHandler(service *service.SubscriptionService) *SubscriptionHandler {
	return &SubscriptionHandler{
		service: service,
	}
}

// Create subscription
// @Summary Create subscription
// @Description Create a new subscription
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param subscription body database.Subscription true "Subscription"
// @Success 201 {object} database.Subscription
// @Failure 400 {object} map[string]string
// @Router /subscriptions [post]
func (h *SubscriptionHandler) Create(c *gin.Context) {
	var sub database.Subscription

	if err := c.ShouldBindJSON(&sub); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Create(&sub); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, sub)
}

// Get all subscriptions
// @Summary Get all subscriptions
// @Tags subscriptions
// @Produce json
// @Success 200 {array} database.Subscription
// @Router /subscriptions [get]
func (h *SubscriptionHandler) GetAll(c *gin.Context) {
	subs, err := h.service.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, subs)
}

// Get subscription by ID
// @Summary Get subscription
// @Tags subscriptions
// @Produce json
// @Param id path string true "Subscription ID"
// @Success 200 {object} database.Subscription
// @Failure 404 {object} map[string]string
// @Router /subscriptions/{id} [get]
func (h *SubscriptionHandler) GetID(c *gin.Context) {
	id := c.Param("id")

	sub, err := h.service.GetID(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, sub)
}

// Delete subscription
// @Summary Delete subscription
// @Tags subscriptions
// @Param id path string true "Subscription ID"
// @Success 204
// @Router /subscriptions/{id} [delete]
func (h *SubscriptionHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	if err := h.service.Delete(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "subscription deleted"})
}

// Summary subscriptions
// @Summary Get total subscription cost
// @Description Calculate total cost of subscriptions for a period
// @Tags subscriptions
// @Produce json
// @Param period_start query string true "Start period (MM-YYYY)"
// @Param period_end query string true "End period (MM-YYYY)"
// @Param service_name query string false "Service name"
// @Param user_name query string false "User name"
// @Success 200 {object} map[string]int64
// @Failure 400 {object} map[string]string
// @Router /subscriptions/summary [get]
func (h *SubscriptionHandler) Summary(c *gin.Context) {
	startDate := c.Query("period_start")
	endDate := c.Query("period_end")
	serviceName := c.Query("service_name")
	userName := c.Query("user_name")

	total, err := h.service.Summary(startDate, endDate, serviceName, userName)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": total})
}
