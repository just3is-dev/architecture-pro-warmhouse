package handlers

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"smarthome/services"

	"github.com/gin-gonic/gin"
)

// SensorHandler handles sensor-related requests
type SensorHandler struct {
	TemperatureService *services.TemperatureService
}

// NewSensorHandler creates a new SensorHandler
func NewSensorHandler(temperatureService *services.TemperatureService) *SensorHandler {
	return &SensorHandler{
		TemperatureService: temperatureService,
	}
}

// RegisterRoutes registers the sensor routes
func (h *SensorHandler) RegisterRoutes(router *gin.RouterGroup) {
	sensors := router.Group("/sensors")
	{
		sensors.GET("", h.GetSensors)
		sensors.GET("/:id", h.GetSensorByID)
		sensors.POST("", h.CreateSensor)
		sensors.PUT("/:id", h.UpdateSensor)
		sensors.DELETE("/:id", h.DeleteSensor)
		sensors.PATCH("/:id/value", h.UpdateSensorValue)
		sensors.GET("/temperature/:location", h.GetTemperatureByLocation)
	}
}

// GET /api/v1/sensors
func (h *SensorHandler) GetSensors(c *gin.Context) {
	resp, err := http.Get("http://device-service:8082/devices")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch devices"})
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	c.Data(resp.StatusCode, "application/json", body)
}

// GET /api/v1/sensors/:id
func (h *SensorHandler) GetSensorByID(c *gin.Context) {
	id := c.Param("id")

	resp, err := http.Get("http://device-service:8082/devices/" + id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch device"})
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	c.Data(resp.StatusCode, "application/json", body)
}

// POST /api/v1/sensors
func (h *SensorHandler) CreateSensor(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)

	resp, err := http.Post(
		"http://device-service:8082/devices",
		"application/json",
		bytes.NewBuffer(body),
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to call device-service"})
		return
	}

	defer resp.Body.Close()
	responseBody, _ := io.ReadAll(resp.Body)

	c.Data(resp.StatusCode, "application/json", responseBody)
}

// PUT /api/v1/sensors/:id
func (h *SensorHandler) UpdateSensor(c *gin.Context) {
	id := c.Param("id")

	body, _ := io.ReadAll(c.Request.Body)

	req, _ := http.NewRequest(
		http.MethodPut,
		"http://device-service:8082/devices/"+id,
		bytes.NewBuffer(body),
	)

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to call device-service"})
		return
	}

	defer resp.Body.Close()

	responseBody, _ := io.ReadAll(resp.Body)

	c.Data(resp.StatusCode, "application/json", responseBody)
}

// DELETE /api/v1/sensors/:id
func (h *SensorHandler) DeleteSensor(c *gin.Context) {
	id := c.Param("id")

	req, _ := http.NewRequest(
		http.MethodDelete,
		"http://device-service:8082/devices/"+id,
		nil,
	)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to call device-service"})
		return
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	c.Data(resp.StatusCode, "application/json", body)
}

// PATCH /api/v1/sensors/:id/value
func (h *SensorHandler) UpdateSensorValue(c *gin.Context) {
	id := c.Param("id")

	body, _ := io.ReadAll(c.Request.Body)

	req, _ := http.NewRequest(
		http.MethodPatch,
		"http://device-service:8082/devices/"+id+"/status",
		bytes.NewBuffer(body),
	)

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to call device-service"})
		return
	}

	defer resp.Body.Close()

	responseBody, _ := io.ReadAll(resp.Body)

	c.Data(resp.StatusCode, "application/json", responseBody)
}

// GET /api/v1/sensors/temperature/:location
func (h *SensorHandler) GetTemperatureByLocation(c *gin.Context) {
	location := c.Param("location")

	if location == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Location is required"})
		return
	}

	tempData, err := h.TemperatureService.GetTemperature(location)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Failed to fetch temperature data: %v", err),
		})
		return
	}

	c.JSON(http.StatusOK, tempData)
}
