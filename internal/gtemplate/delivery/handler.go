package delivery

import (
	"net/http"
	"strconv"

	"github.com/MelhorzinOfficial/melhorzincraft-back/internal/gtemplate/usecase"
	"github.com/MelhorzinOfficial/melhorzincraft-back/internal/gtemplate/usecase/dto"
	"github.com/gin-gonic/gin"

	_ "github.com/MelhorzinOfficial/melhorzincraft-back/internal/core/response"
)

// RegisterRoutes registers template routes in the router
func RegisterRoutes(g *gin.RouterGroup, templateUC *usecase.TemplateUC) {
	h := handler{
		templateUC: templateUC,
	}

	g.POST("/templates", h.Create)
	g.PUT("/templates/:id", h.Update)
	g.GET("/templates/:id", h.Get)
	g.DELETE("/templates/:id", h.Delete)
}

type handler struct {
	templateUC *usecase.TemplateUC
}

// Create godoc
// @Summary Create a new game template
// @Description Creates a new game template with the specified name and docker compose
// @Tags templates
// @Accept json
// @Produce json
// @Param request body dto.CreateRequest true "Template data to create"
// @Success 200 {object} response.Response[dto.CreateResponse] "Template created successfully"
// @Failure 400 {object} response.Response[any] "Bad request"
// @Failure 500 {object} response.Response[any] "Internal server error"
// @Router /game/templates [post]
func (h *handler) Create(c *gin.Context) {
	var req dto.CreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := h.templateUC.Create(c.Request.Context(), &req)
	c.JSON(resp.Code, resp)
}

// Update godoc
// @Summary Update an existing game template
// @Description Updates a game template with the specified name and/or docker compose
// @Tags templates
// @Accept json
// @Produce json
// @Param id path int true "Template ID"
// @Param request body dto.UpdateRequest true "Template data to update"
// @Success 200 {object} response.Response[dto.UpdateResponse] "Template updated successfully"
// @Failure 400 {object} response.Response[any] "Invalid ID or bad request"
// @Failure 404 {object} response.Response[any] "Template not found"
// @Failure 500 {object} response.Response[any] "Internal server error"
// @Router /game/templates/{id} [put]
func (h *handler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var req dto.UpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.Id = id
	resp := h.templateUC.Update(c.Request.Context(), &req)
	c.JSON(resp.Code, resp)
}

// Get godoc
// @Summary Get a game template
// @Description Gets a game template by ID
// @Tags templates
// @Produce json
// @Param id path int true "Template ID"
// @Success 200 {object} response.Response[dto.ShowResponse] "Template found successfully"
// @Failure 400 {object} response.Response[any] "Invalid ID"
// @Failure 404 {object} response.Response[any] "Template not found"
// @Failure 500 {object} response.Response[any] "Internal server error"
// @Router /game/templates/{id} [get]
func (h *handler) Get(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	req := dto.ShowRequest{Id: id}
	resp := h.templateUC.Show(c.Request.Context(), &req)
	c.JSON(resp.Code, resp)
}

// Delete godoc
// @Summary Delete a game template
// @Description Deletes a game template by ID
// @Tags templates
// @Produce json
// @Param id path int true "Template ID"
// @Success 200 {object} response.Response[dto.DeleteResponse] "Template deleted successfully"
// @Failure 400 {object} response.Response[any] "Invalid ID"
// @Failure 404 {object} response.Response[any] "Template not found"
// @Failure 500 {object} response.Response[any] "Internal server error"
// @Router /game/templates/{id} [delete]
func (h *handler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	req := dto.DeleteRequest{Id: id}
	resp := h.templateUC.Delete(c.Request.Context(), &req)
	c.JSON(resp.Code, resp)
}
