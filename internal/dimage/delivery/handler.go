package delivery

import (
	"github.com/MelhorzinOfficial/melhorzincraft-back/internal/dimage/usecase"
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

}

func (h *handler) Update(c *gin.Context) {

}

func (h *handler) Get(c *gin.Context) {

}

func (h *handler) Delete(c *gin.Context) {

}
