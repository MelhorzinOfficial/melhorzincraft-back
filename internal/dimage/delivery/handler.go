package delivery

import (
	"net/http"
	"strconv"

	"github.com/MelhorzinOfficial/melhorzincraft-back/internal/dimage/usecase"
	"github.com/MelhorzinOfficial/melhorzincraft-back/internal/dimage/usecase/dto"
	"github.com/gin-gonic/gin"

	_ "github.com/MelhorzinOfficial/melhorzincraft-back/internal/core/response"
)

// RegisterRoutes registers image routes in the router
func RegisterRoutes(g *gin.Engine, imageUC *usecase.ImageUC) {
	h := handler{
		imageUC: imageUC,
	}

	g.POST("/api/v1/images", h.Create)
	g.PUT("/api/v1/images/:id", h.Update)
	g.GET("/api/v1/images/:id", h.Get)
	g.DELETE("/api/v1/images/:id", h.Delete)
}

type handler struct {
	imageUC *usecase.ImageUC
}

// Create godoc
// @Summary Create a new Docker image
// @Description Creates a new Docker image with the specified tag and repository
// @Tags images
// @Accept json
// @Produce json
// @Param request body dto.CreateRequest true "Image data to create"
// @Success 200 {object} response.Response[dto.CreateResponse] "Image created successfully"
// @Failure 400 {object} response.Response[any] "Bad request"
// @Failure 409 {object} response.Response[any] "Image already exists"
// @Failure 500 {object} response.Response[any] "Internal server error"
// @Router /images [post]
func (h *handler) Create(c *gin.Context) {
	var req dto.CreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := h.imageUC.Create(c.Request.Context(), &req)
	c.JSON(resp.Code, resp)
}

// Update godoc
// @Summary Update an existing Docker image
// @Description Updates a Docker image with the specified tag and/or repository
// @Tags images
// @Accept json
// @Produce json
// @Param id path int true "Image ID"
// @Param request body dto.UpdateRequest true "Image data to update"
// @Success 200 {object} response.Response[dto.UpdateResponse] "Image updated successfully"
// @Failure 400 {object} response.Response[any] "Invalid ID or bad request"
// @Failure 404 {object} response.Response[any] "Image not found"
// @Failure 500 {object} response.Response[any] "Internal server error"
// @Router /images/{id} [put]
func (h *handler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id invalid"})
		return
	}

	var req dto.UpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.Id = id
	resp := h.imageUC.Update(c.Request.Context(), &req)
	c.JSON(resp.Code, resp)
}

// Get godoc
// @Summary Get a Docker image
// @Description Gets a Docker image by ID
// @Tags images
// @Produce json
// @Param id path int true "Image ID"
// @Success 200 {object} response.Response[dto.ShowResponse] "Image found successfully"
// @Failure 400 {object} response.Response[any] "Invalid ID"
// @Failure 404 {object} response.Response[any] "Image not found"
// @Failure 500 {object} response.Response[any] "Internal server error"
// @Router /images/{id} [get]
func (h *handler) Get(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id invalid"})
		return
	}

	req := dto.ShowRequest{Id: id}
	resp := h.imageUC.Show(c.Request.Context(), &req)
	c.JSON(resp.Code, resp)
}

// Delete godoc
// @Summary Delete a Docker image
// @Description Deletes a Docker image by ID
// @Tags images
// @Produce json
// @Param id path int true "Image ID"
// @Success 200 {object} response.Response[dto.DeleteResponse] "Image deleted successfully"
// @Failure 400 {object} response.Response[any] "Invalid ID"
// @Failure 404 {object} response.Response[any] "Image not found"
// @Failure 500 {object} response.Response[any] "Internal server error"
// @Router /images/{id} [delete]
func (h *handler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id invalid"})
		return
	}

	req := dto.DeleteRequest{Id: id}
	resp := h.imageUC.Delete(c.Request.Context(), &req)
	c.JSON(resp.Code, resp)
}
