package device

import (
	"iot/internal/entity"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h Handler) GetListDevices(ctx *gin.Context) {
	devices, err := h.uc.GetListDevices(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var data []gin.H
	for _, device := range devices {
		data = append(data, gin.H{
			"id":         device.ID,
			"name":       device.Name,
			"type":       device.Types,
			"created_at": device.CreatedAt,
			"updated_at": device.UpdatedAt,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func (h Handler) GetDevice(ctx *gin.Context) {
	id := ctx.Param("id")

	device, err := h.uc.GetDevice(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": device,
	})
}

func (h Handler) CreateDevice(ctx *gin.Context) {
	var input entity.CreateDevice

	err := ctx.BindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	device, err := h.uc.CreateDevice(ctx.Request.Context(), input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": device,
	})
}

func (h Handler) UpdateDevice(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 32, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var input entity.UpdateDevice
	err = ctx.BindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	input.ID = uint(id)

	device, err := h.uc.UpdateDevice(ctx.Request.Context(), input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": device,
	})
}

func (h Handler) DeleteDevice(ctx *gin.Context) {
	id := ctx.Param("id")

	err := h.uc.DeleteDevice(ctx.Request.Context(), id)
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
