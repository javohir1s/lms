package handler

import (
	"context"
	"fmt"
	"lms_back/api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateStudent godoc
// @Router 		/student [POST]
// @Summary 	create a student
// @Description This api is creates a new student and returns all data
// @Tags 		student
// @Accept		json
// @Produce		json
// @Param		car  body      models.CreateStudent true "student"
// @Success		200  {object}  models.Student
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) CreateStudent(c *gin.Context) {
	student := models.Student{}

	if err := c.ShouldBindJSON(&student); err != nil {

		handleResponse(c, "error while decoding request body", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.Service.Student().Create(context.Background(), student)
	if err != nil {
		handleResponse(c, "error while creating student", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "created successfully", http.StatusOK, id)
}

// UpdateStudent godoc
// @Router                /student/{id} [PUT]
// @Summary 			  update a student
// @Description:          this api updates student information
// @Tags 			      student
// @Accept 			      json
// @Produce 		      json
// @Param 			      id path string true "Student ID"
// @Param       		  car body models.UpdateStudent true "student"
// @Success 		      200 {object} models.Student
// @Failure 		      400 {object} models.Response
// @Failure               404 {object} models.Response
// @Failure 		      500 {object} models.Response
func (h Handler) UpdateStudent(c *gin.Context) {

	student := models.Student{}
	if err := c.ShouldBindJSON(&student); err != nil {
		handleResponse(c, "error while decoding request body", http.StatusBadRequest, err.Error())
		return
	}

	student.ID = c.Query("id")
	err := uuid.Validate(student.ID)
	if err != nil {
		handleResponse(c, "error while validating", http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.Service.Student().Update(context.Background(), student)
	if err != nil {
		handleResponse(c, "error while updating student", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(c, "updated successfully", http.StatusOK, id)
}

// GetAllStudents godoc
// @Router 			/student [GET]
// @Summary 		get all students
// @Description 	This API returns student list
// @Tags 			student
// Accept			json
// @Produce 		json
// @Param 			page query int false "page number"
// @Param 			limit query int false "limit per page"
// @Param 			search query string false "search keyword"
// @Success 		200 {object} models.GetAllStudentsResponse
// @Failure 		400 {object} models.Response
// @Failure         404 {object} models.Response
// @Failure 		500 {object} models.Response
func (h Handler) GetAllStudent(c *gin.Context) {
	var (
		request = models.GetAllStudentsRequest{}
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

	student, err := h.Service.Student().GetAll(context.Background(), request)
	if err != nil {
		handleResponse(c, "error while getting student", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(c, "", http.StatusOK, student)
}

// GetByIDStudent godoc
// @Router       /student/{id} [GET]
// @Summary      return a student by ID
// @Description  Retrieves a student by its ID
// @Tags         student
// @Accept       json
// @Produce      json
// @Param        id path string true "Student ID"
// @Success      200 {object} models.GetStudent
// @Failure      400 {object} models.Response
// @Failure      404 {object} models.Response
// @Failure      500 {object} models.Response
func (h Handler) GetByIDStudent(c *gin.Context) {

	id := c.Param("id")
	fmt.Println("id: ", id)

	student, err := h.Service.Student().GetByID(context.Background(), id)
	if err != nil {
		handleResponse(c, "error while getting student by id", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(c, "", http.StatusOK, student)
}

// DeleteStudent godoc
// @Router          /student/{id} [DELETE]
// @Summary         delete a student by ID
// @Description     Deletes a student by its ID
// @Tags            student
// @Accept          json
// @Produce         json
// @Param           id path string true "Student ID"
// @Success         200 {string} models.Response
// @Failure         400 {object} models.Response
// @Failure         404 {object} models.Response
// @Failure         500 {object} models.Response
func (h Handler) DeleteStudent(c *gin.Context) {

	id := c.Param("id")
	fmt.Println("id: ", id)

	err := uuid.Validate(id)
	if err != nil {
		handleResponse(c, "error while validating id", http.StatusBadRequest, err.Error())
		return
	}
	err = h.Service.Student().Delete(context.Background(), id)
	if err != nil {
		handleResponse(c, "error while deleting student", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(c, "deleted student", http.StatusOK, id)
}
