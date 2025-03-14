// Package helper "helper package"
package helper

import (
	"github.com/gin-gonic/gin"
)

func RespondMessage(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, map[string]interface{}{
		"message": message,
		"error":   false,
	})
}

type ResponseData struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Error   bool        `json:"error"`
}

func RespondWithData(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, map[string]interface{}{
		"message": message,
		"error":   false,
		"data":    data,
	})
}

func RespondWithPagination(c *gin.Context, code int, message string, total int, page int, perPage int, dataName string, data interface{}) {
	c.JSON(code, map[string]interface{}{
		"message": message,
		"error":   false,
		"data": map[string]interface{}{
			"total":    total,
			"page":     page,
			"per_page": perPage,
			"dataName": dataName,
			"data":     data,
		},
	})
}

func RespondError(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, map[string]interface{}{
		"message": message,
		"error":   true,
	})
}

type PaginatedResponse struct {
	Message string                `json:"message"`
	Error   bool                  `json:"error"`
	Data    PaginatedResponseData `json:"data"`
}

type PaginatedResponseData struct {
	Total    int         `json:"total"`
	Page     int         `json:"page"`
	PerPage  int         `json:"per_page"`
	DataName string      `json:"dataName"`
	Data     interface{} `json:"data"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Error   bool   `json:"error"`
}
