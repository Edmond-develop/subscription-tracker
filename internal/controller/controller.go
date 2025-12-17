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

func (h *SubscriptionHandler) GetAll(c *gin.Context) {
	subs, err := h.service.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, subs)
}

func (h *SubscriptionHandler) GetID(c *gin.Context) {
	id := c.Param("id")

	sub, err := h.service.GetID(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, sub)
}

func (h *SubscriptionHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	if err := h.service.Delete(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "subscription deleted"})
}

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
