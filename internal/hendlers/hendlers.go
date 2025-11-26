package hendlers

import (
	"database/sql"
	"fmt"
	internal "github.com/Edmond-develop/subscription-tracker/internal/database"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// @Summary Создать подписку
// @Router /subscriptions [post]
func CreateSubscriptions(c *gin.Context, db *sql.DB) {
	var s internal.Subscription
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
		return
	}

	if s.ServiceName == "" || s.Price <= 0 || s.UserName == "" || s.StartDate == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing or invalid fields"})
		return
	}

	start, err := time.Parse("01-2006", s.StartDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid start_date format"})
		return
	}
	var end *time.Time
	if s.EndDate != "" {
		e, err := time.Parse("01-2006", s.EndDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid end_date format"})
			return
		}
		end = &e
	}
	err = db.QueryRow(
		`INSERT INTO subscriptions (service_name, price, user_name, start_date, end_date) 
				VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		s.ServiceName, s.Price, s.UserName, start, end).Scan(&s.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database Create error: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, s)
}

// @Summary Получить список подписок
// @Router /subscriptions [get]
func ListSubscriptions(c *gin.Context, db *sql.DB) {
	rows, err := db.Query(`SELECT id, service_name, price, user_name, start_date, end_date 
								 FROM subscriptions`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database List error: " + err.Error()})
		return
	}
	defer rows.Close()
	list := []internal.Subscription{}

	for rows.Next() {
		var s internal.Subscription
		var start, end *time.Time

		err = rows.Scan(&s.ID, &s.ServiceName, &s.Price, &s.UserName, &start, &end)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database List error: " + err.Error()})
			return
		}

		if start != nil {
			s.StartDate = start.Format("01-2006")
		}

		if end != nil {
			s.EndDate = end.Format("01-2006")
		}

		list = append(list, s)
	}
	c.JSON(http.StatusOK, list)
}

// @Summary Получить подписку
// @Router /subscriptions/{id} [get]
func GetSubscription(c *gin.Context, db *sql.DB) {
	id := c.Param("id")
	var s internal.Subscription
	var start, end *time.Time

	err := db.QueryRow(`SELECT id, service_name, price, user_name, start_date, end_date 
							  FROM subscriptions WHERE id = $1`, id).Scan(&s.ID, &s.ServiceName, &s.Price, &s.UserName, &start, &end)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "subscription not found"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database Get error: " + err.Error()})
		return
	}

	if start != nil {
		s.StartDate = start.Format("01-2006")
	}

	if end != nil {
		s.EndDate = end.Format("01-2006")
	}

	c.JSON(http.StatusOK, s)
}

// @Summary Удалить подписку
// @Router /subscriptions/{id} [delete]
func DeleteSubscription(c *gin.Context, db *sql.DB) {
	id := c.Param("id")
	_, err := db.Exec(`DELETE FROM subscriptions WHERE id = $1`, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database Delete error: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "subscription deleted"})
}

// @Summary Получить сумму подписок за период
// @Router /subscriptions/summary [get]
func Summary(c *gin.Context, db *sql.DB) {
	periodStart := c.Query("period_start")
	periodEnd := c.Query("period_end")
	serviceName := c.Query("service_name")
	userName := c.Query("user_name")

	if periodStart == "" || periodEnd == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing or invalid fields"})
		return
	}
	start, err1 := time.Parse("01-2006", periodStart)
	end, err2 := time.Parse("01-2006", periodEnd)
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid date format (MM-YYYY)"})
		return
	}

	query := `SELECT SUM(price) FROM subscriptions WHERE start_date <= $2 AND (end_date IS NULL OR end_date >= $1)`

	args := []interface{}{start, end}
	argPos := 3

	if serviceName != "" {
		query += fmt.Sprintf(" AND service_name = $%d", argPos)
		args = append(args, serviceName)
		argPos++
	}

	if userName != "" {
		query += fmt.Sprintf(" AND user_name = $%d", argPos)
		args = append(args, userName)
		argPos++
	}

	var total sql.NullInt64
	err := db.QueryRow(query, args...).Scan(&total)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database Summary error: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"total_price": total.Int64})
}
