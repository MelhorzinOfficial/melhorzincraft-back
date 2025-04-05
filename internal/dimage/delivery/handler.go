package delivery

import (
	"net/http"
	"strconv"

	"github.com/MelhorzinOfficial/melhorzincraft-back/internal/dimage/usecase"
	"github.com/MelhorzinOfficial/melhorzincraft-back/internal/dimage/usecase/dto"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(g *gin.RouterGroup, imageUC *usecase.ImageUC) {
	h := handler{
		imageUC: imageUC,
	}

	g.POST("/images", h.Create)
	g.PUT("/images/:id", h.Update)
	g.GET("/images/:id", h.Get)
	g.DELETE("/images/:id", h.Delete)
}

type handler struct {
	imageUC *usecase.ImageUC
}

func (h *handler) Create(c *gin.Context) {
	var req dto.CreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := h.imageUC.Create(c.Request.Context(), &req)
	c.JSON(resp.Code, resp)
}

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
