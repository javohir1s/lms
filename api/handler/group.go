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

// CreateGroup godoc
// @Router 		   /group [POST]
// @Summary 	   create a group
// @Description    This api is creates a new group and returns its id
// @Tags 		   group
// @Accept		   json
// @Produce		   json
// @Param		   group body    models.CreateGroup true "group"
// @Success		   200  {object}  models.Group
// @Failure		   400  {object}  models.Response
// @Failure		   404  {object}  models.Response
// @Failure		   500  {object}  models.Response
func (h Handler) CreateGroup(c *gin.Context) {
	group := models.Group{}

	if err := c.ShouldBindJSON(&group); err != nil {
		handleResponse(c, "error while decoding request body", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.Service.Group().Create(context.Background(), group)
	if err != nil {
		handleResponse(c, "error while creating group", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, id)
}

// UpdateGroup godoc
// @Router                /group/{id} [PUT]
// @Summary 			  update a group
// @Description:          this api updates group information
// @Tags 			      group
// @Accept 			      json
// @Produce 		      json
// @Param 			      id path string true "group ID"
// @Param       		  group body models.UpdateGroup true "group"
// @Success 		      200 {object} models.Group
// @Failure 		      400 {object} models.Response
// @Failure               404 {object} models.Response
// @Failure 		      500 {object} models.Response
func (h Handler) UpdateGroup(c *gin.Context) {
	group := models.Group{}
	if err := c.ShouldBindJSON(&group); err != nil {
		handleResponse(c, "error while decoding request body", http.StatusBadRequest, err.Error())
		return
	}
	group.Id = c.Param("id")
	err := uuid.Validate(group.Id)
	if err != nil {
		handleResponse(c, "error while validating", http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.Service.Group().Update(context.Background(), group)
	if err != nil {
		handleResponse(c, "error while updating group", http.StatusInternalServerError, err)
		return
	}
	handleResponse(c, "updated successfully", http.StatusOK, id)
}

// GetAllGroup godoc
// @Router 			/group [GET]
// @Summary 		get all Group
// @Description 	This API returns group list
// @Tags 			group
// Accept			json
// @Produce 		json
// @Param 			page query int false "page number"
// @Param 			limit query int false "limit per page"
// @Param 			search query string false "search keyword"
// @Success 		200 {object} models.GetAllGroupsResponse
// @Failure 		400 {object} models.Response
// @Failure         404 {object} models.Response
// @Failure 		500 {object} models.Response
func (h Handler) GetAllGroups(c *gin.Context) {
	var (
		request = models.GetAllGroupsRequest{}
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

	groups, err := h.Service.Group().GetAll(context.Background(), request)
	if err != nil {
		handleResponse(c, "error while getting groups", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(c, "", http.StatusOK, groups)
}

// GetByIDGroup godoc
// @Router       /group/{id} [GET]
// @Summary      return a group by ID
// @Description  Retrieves a group by its ID
// @Tags         group
// @Accept       json
// @Produce      json
// @Param        id path string true "Group ID"
// @Success      200 {object} models.GetGroup
// @Failure      400 {object} models.Response
// @Failure      404 {object} models.Response
// @Failure      500 {object} models.Response
func (h Handler) GetByIDGroup(c *gin.Context) {
	
	id := c.Param("id")
	fmt.Println("id: ", id)

	group, err := h.Service.Group().GetByID(context.Background(), id)
	if err != nil {
		fmt.Println("error while getting group by id")
		handleResponse(c, "", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(c, "", http.StatusOK, group)
}

// DeleteGroup godoc
// @Router          /group/{id} [DELETE]
// @Summary         delete a group by ID
// @Description     Deletes a group by its ID
// @Tags            group
// @Accept          json
// @Produce         json
// @Param           id path string true "Group ID"
// @Success         200 {string} models.Response
// @Failure         400 {object} models.Response
// @Failure         404 {object} models.Response
// @Failure         500 {object} models.Response
func (h Handler) DeleteGroup(c *gin.Context) {

	id := c.Param("id")
	fmt.Println("id: ", id)

	err := uuid.Validate(id)
	if err != nil {
		handleResponse(c, "error while validating id", http.StatusBadRequest, err.Error())
		return
	}
	err = h.Service.Group().Delete(context.Background(), id)
	if err != nil {
		handleResponse(c, "error while deleting group", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(c, "deleted successfully", http.StatusOK, id)
}
