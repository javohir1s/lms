package handler

import (
	"context"
	"fmt"
	_ "lms_back/api/docs"
	"lms_back/api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// CreateAdmin godoc
// @Router     /admin [POST]
// @Summary    create a admin
// Description This api creates a new admin return its id
// @Tags       admin
// @Accept     json
// @Produce    json
// @Param      admin body   models.CreateAdmin true "admin"
// @Success    200 {object} models.Admin
// @Failure    400 {object} models.Response
// @Failure    404 {object} models.Response
// @Failure    500 {object} models.Response
func (h Handler) CreateAdmin(c *gin.Context) {
	admin := models.Admin{}

	if err := c.ShouldBindJSON(&admin); err != nil {
		handleResponse(c, "error while decoding request body", http.StatusBadRequest, err.Error())
		return
	}

	hashedpass, err := bcrypt.GenerateFromPassword([]byte(
		admin.Password,
	), bcrypt.DefaultCost)

	admin.Password = string(hashedpass)
	id, err := h.Service.Admin().Create(context.Background(), admin)
	if err != nil {
		handleResponse(c, "error while creating admin", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "created admin", http.StatusOK, id)
}

// UpdateAdmin godoc
// @Router                /admin/{id} [PUT]
// @Summary               Update Admin
// @Description           This API Updates admin Information
// @Tags   	  			  admin
// @Accept     	          json
// @Produce               json
// @Param				  id path string true "Admin Id"
// @Param                 admin body models.UpdateAdmin true "admin"
// @Success 		      200 {object} models.Admin
// @Failure 		      400 {object} models.Response
// @Failure               404 {object} models.Response
// @Failure 		      500 {object} models.Response
func (h Handler) UpdateAdmin(c *gin.Context) {
	admin := models.Admin{}
	if err := c.ShouldBindJSON(&admin); err != nil {
		handleResponse(c, "error while decoding request body", http.StatusBadRequest, err.Error())
		return
	}
	admin.Id = c.Param("id")
	err := uuid.Validate(admin.Id)
	if err != nil {
		handleResponse(c, "error while validating", http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.Service.Admin().Update(context.Background(), admin)
	if err != nil {
		handleResponse(c, "error while updating admin", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(c, "updated successfully", http.StatusOK, id)
}

// GetAllAdmin godoc
// @Router 			/admin [GET]
// @Summary 		get all admin
// @Description 	This API returns admin list
// @Tags 			admin
// Accept			json
// @Produce 		json
// @Param 			page query int false "page number"
// @Param 			limit query int false "limit per page"
// @Param 			search query string false "search keyword"
// @Success 		200 {object} models.GetAllAdminsResponse
// @Failure 		400 {object} models.Response
// @Failure         404 {object} models.Response
// @Failure 		500 {object} models.Response
func (h Handler) GetAllAdmins(c *gin.Context) {
	var (
		request = models.GetAllAdminsRequest{}
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

	admins, err := h.Service.Admin().GetAll(context.Background(), request)
	if err != nil {
		handleResponse(c, "error while getting admins", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(c, "", http.StatusOK, admins)
}

// GetByIDAdmin godoc
// @Router       /admin/{id} [GET]
// @Summary      return a admin by ID
// @Description  Retrieves a admin by its ID
// @Tags         admin
// @Accept       json
// @Produce      json
// @Param        id path string true "Admin ID"
// @Success      200 {object} models.GetAdmin
// @Failure      400 {object} models.Response
// @Failure      404 {object} models.Response
// @Failure      500 {object} models.Response
func (h Handler) GetByIDAdmin(c *gin.Context) {

	id := c.Param("id")
	fmt.Println("id: ", id)

	admin, err := h.Service.Admin().GetByID(context.Background(), id)
	if err != nil {
		handleResponse(c, "error while getting admin by id", http.StatusInternalServerError, err)
		return
	}
	handleResponse(c, "", http.StatusOK, admin)
}

// DeleteAdmin godoc
// @Router          /admin/{id} [DELETE]
// @Summary         delete a admin by ID
// @Description     Deletes a admin by its ID
// @Tags            admin
// @Accept          json
// @Produce         json
// @Param           id path string true "Admin ID"
// @Success         200 {string} models.Response
// @Failure         400 {object} models.Response
// @Failure         404 {object} models.Response
// @Failure         500 {object} models.Response
func (h Handler) DeleteAdmin(c *gin.Context) {

	id := c.Param("id")
	fmt.Println("id: ", id)

	err := uuid.Validate(id)
	if err != nil {
		handleResponse(c, "error while validating id", http.StatusBadRequest, err.Error())
		return
	}
	err = h.Service.Admin().Delete(context.Background(), id)
	if err != nil {
		handleResponse(c, "error while deleting admin", http.StatusInternalServerError, err)
		return
	}
	handleResponse(c, "deleted admin", http.StatusOK, id)
}
