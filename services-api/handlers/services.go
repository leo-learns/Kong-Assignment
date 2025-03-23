package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"services-api/services-api/db"
	"services-api/services-api/models"
	"strconv"
)

// GetServices returns a list of services with filtering, sorting, and pagination
func GetServices(c *gin.Context) {
	var services []models.Service
	query := db.DB

	// Filtering by search term (name or description)
	if search := c.Query("search"); search != "" {
		query = query.Where("name LIKE ? OR description LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	// Sorting
	sort := c.Query("sort")
	order := c.Query("order")
	if sort != "" {
		if order == "desc" {
			query = query.Order(sort + " desc")
		} else {
			query = query.Order(sort) // Default to ascending
		}
	}

	// Pagination
	limit := 10 // Default limit
	if l := c.Query("limit"); l != "" {
		if n, err := strconv.Atoi(l); err == nil && n > 0 {
			limit = n
		}
	}
	offset := 0 // Default offset
	if o := c.Query("offset"); o != "" {
		if n, err := strconv.Atoi(o); err == nil && n >= 0 {
			offset = n
		}
	}
	query = query.Limit(limit).Offset(offset)

	// Execute query
	if err := query.Find(&services).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, services)
}

// GetService fetches a specific service by ID
func GetService(c *gin.Context) {
	id := c.Param("id")
	var service models.Service
	if err := db.DB.Where("id = ?", id).First(&service).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Service not found"})
		return
	}
	c.JSON(http.StatusOK, service)
}

// GetServiceVersions retrieves all versions of a specific service
func GetServiceVersions(c *gin.Context) {
	id := c.Param("id")
	var versions []models.Version
	if err := db.DB.Where("service_id = ?", id).Find(&versions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, versions)
}
