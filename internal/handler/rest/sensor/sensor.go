package sensor

import (
	"iot/internal/entity"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h Handler) GetListSensors(ctx *gin.Context) {
	sensors, err := h.uc.GetListSensors(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var data []entity.Sensor
	data = append(data, sensors...)

	ctx.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func (h Handler) GetSensor(ctx *gin.Context) {
	id := ctx.Param("id")

	sensor, err := h.uc.GetSensor(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": sensor,
	})
}

func (h Handler) CreateSensor(ctx *gin.Context) {
	var input entity.CreateSensor

	err := ctx.BindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	sensor, err := h.uc.CreateSensor(ctx.Request.Context(), input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": sensor,
	})
}

func (h Handler) UpdateSensor(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 32, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var input entity.UpdateSensor
	err = ctx.BindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	input.ID = uint(id)

	sensor, err := h.uc.UpdateSensor(ctx.Request.Context(), input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": sensor,
	})
}

func (h Handler) DeleteSensor(ctx *gin.Context) {
	id := ctx.Param("id")

	err := h.uc.DeleteSensor(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": nil,
	})
}
