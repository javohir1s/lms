package handler

import (
	"context"
	"fmt"
	_ "lms_back/api/docs"
	"lms_back/api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateSchedule godoc
// @Router 		/schedule [POST]
// @Summary 	create a schedule
// @Description This api is creates a new schedule and returns its id
// @Tags 		schedule
// @Accept		json
// @Produce		json
// @Param		schedule body  models.CreateSchedule true "schedule"
// @Success		200  {object}  models.Schedule
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) CreateSchedule(c *gin.Context) {
	schedule := models.Schedule{}

	if err := c.ShouldBindJSON(&schedule); err != nil {
		handleResponse(c, "error while decoding request body", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.Service.Schedule().Create(context.Background(), schedule)
	if err != nil {
		handleResponse(c, "error while creating schedule", http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, "created successfully", http.StatusOK, id)
}

// UpdateSchedule godoc
// @Router                /schedule/{id} [PUT]
// @Summary 			  update a schedule
// @Description:          this api updates schedule information
// @Tags 			      schedule
// @Accept 			      json
// @Produce 		      json
// @Param 			      id path string true "Schedule ID"
// @Param       		  schedule body models.UpdateSchedule true "schedule"
// @Success 		      200 {object} models.Schedule
// @Failure 		      400 {object} models.Response
// @Failure               404 {object} models.Response
// @Failure 		      500 {object} models.Response
func (h Handler) UpdateSchedule(c *gin.Context) {

	schedule := models.Schedule{}
	if err := c.ShouldBindJSON(&schedule); err != nil {
		handleResponse(c, "error while decoding request body", http.StatusBadRequest, err.Error())
		return
	}

	schedule.Id = c.Param("id")
	err := uuid.Validate(schedule.Id)
	if err != nil {
		handleResponse(c, "error while validating", http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.Service.Schedule().Update(context.Background(), schedule)
	if err != nil {
		handleResponse(c, "error while updating schedule", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(c, "", http.StatusOK, id)
}

// GetAllSchedules godoc
// @Router 			/schedule [GET]
// @Summary 		get all schedules
// @Description 	This API returns schedule list
// @Tags 			schedule
// Accept			json
// @Produce 		json
// @Param 			page query int false "page number"
// @Param 			limit query int false "limit per page"
// @Param 			search query string false "search keyword"
// @Success 		200 {object} models.GetAllSchedulesResponse
// @Failure 		400 {object} models.Response
// @Failure         404 {object} models.Response
// @Failure 		500 {object} models.Response
func (h Handler) GetAllSchedule(c *gin.Context) {
	var (
		request = models.GetAllSchedulesRequest{}
	)
	request.Search = c.Query("search")

	page, err := ParsePageQueryParam(c)
	if err != nil {
		handleResponse(c, "error while parsing page", http.StatusInternalServerError, err.Error())
		return
	}
	limit, err := ParseLimitQueryParam(c)
	if err != nil {
		handleResponse(c, "error while parsing limit", http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println("page: ", page)
	fmt.Println("limit: ", limit)

	request.Page = page
	request.Limit = limit

	schedule, err := h.Service.Schedule().GetAll(context.Background(), request)
	if err != nil {
		handleResponse(c, "error while getting schedule", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(c, "", http.StatusOK, schedule)
}

// GetByIDSchedule godoc
// @Router       /schedule/{id} [GET]
// @Summary      Return a schedule by ID
// @Description  Retrieves a schedule by its ID
// @Tags         schedule
// @Accept       json
// @Produce      json
// @Param        id path string true "Schedule ID"
// @Success      200 {object} models.GetSchedule
// @Failure      400 {object} models.Response
// @Failure      404 {object} models.Response
// @Failure      500 {object} models.Response
func (h Handler) GetByIDSchedule(c *gin.Context) {

	id := c.Param("id")
	fmt.Println("id: ", id)

	schedule, err := h.Service.Schedule().GetByID(context.Background(), id)
	if err != nil {
		handleResponse(c, "error while getting schedule by id", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(c, "", http.StatusOK, schedule)
}

// DeleteSchedule godoc
// @Router          /schedule/{id} [DELETE]
// @Summary         delete a schedule by ID
// @Description     Deletes a schedule by its ID
// @Tags            schedule
// @Accept          json
// @Produce         json
// @Param           id path string true "Schedule ID"
// @Success         200 {string} models.Response
// @Failure         400 {object} models.Response
// @Failure         404 {object} models.Response
// @Failure         500 {object} models.Response
func (h Handler) DeleteSchedule(c *gin.Context) {

	id := c.Param("id")
	fmt.Println("id: ", id)

	err := uuid.Validate(id)
	if err != nil {
		handleResponse(c, "error while validating id", http.StatusBadRequest, err.Error())
		return
	}
	err = h.Service.Schedule().Delete(context.Background(), id)
	if err != nil {
		handleResponse(c, "error while deleting schedule", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(c, "schedule deleted", http.StatusOK, id)
}
