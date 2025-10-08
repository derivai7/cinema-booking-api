package handler

import (
	"cinema-booking-api/internal/domain/dto"
	"cinema-booking-api/internal/pkg/response"
	"cinema-booking-api/internal/usecase"

	"github.com/gin-gonic/gin"
)

type ScheduleHandler struct {
	scheduleUsecase usecase.ScheduleUsecase
}

func NewScheduleHandler(scheduleUsecase usecase.ScheduleUsecase) *ScheduleHandler {
	return &ScheduleHandler{
		scheduleUsecase: scheduleUsecase,
	}
}

func (h *ScheduleHandler) CreateSchedule(c *gin.Context) {
	var req dto.CreateScheduleRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request body", err)
		return
	}

	result, err := h.scheduleUsecase.CreateSchedule(req)
	if err != nil {
		response.BadRequest(c, err.Error(), nil)
		return
	}

	response.Created(c, "Schedule created successfully", result)
}

func (h *ScheduleHandler) GetAllSchedules(c *gin.Context) {
	result, err := h.scheduleUsecase.GetAllSchedules()
	if err != nil {
		response.InternalServerError(c, "Failed to retrieve schedules", err)
		return
	}

	response.Success(c, "Schedules retrieved successfully", result)
}

func (h *ScheduleHandler) GetScheduleByID(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		response.BadRequest(c, "Schedule ID is required", nil)
		return
	}

	result, err := h.scheduleUsecase.GetScheduleByID(id)
	if err != nil {
		response.NotFound(c, err.Error())
		return
	}

	response.Success(c, "Schedule retrieved successfully", result)
}

func (h *ScheduleHandler) UpdateSchedule(c *gin.Context) {
	id := c.Param("id")
	var req dto.UpdateScheduleRequest

	if id == "" {
		response.BadRequest(c, "Schedule ID is required", nil)
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request body", err)
		return
	}

	result, err := h.scheduleUsecase.UpdateSchedule(id, req)
	if err != nil {
		response.BadRequest(c, err.Error(), nil)
		return
	}

	response.Success(c, "Schedule updated successfully", result)
}

func (h *ScheduleHandler) DeleteSchedule(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		response.BadRequest(c, "Schedule ID is required", nil)
		return
	}

	err := h.scheduleUsecase.DeleteSchedule(id)
	if err != nil {
		response.NotFound(c, err.Error())
		return
	}

	response.Success(c, "Schedule deleted successfully", nil)
}
