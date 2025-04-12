package delivery

import (
	"net/http"
	"strconv"

	"github.com/MelhorzinOfficial/melhorzincraft-back/internal/gtemplate/usecase"
	"github.com/MelhorzinOfficial/melhorzincraft-back/internal/gtemplate/usecase/dto"
	"github.com/gin-gonic/gin"

	_ "github.com/MelhorzinOfficial/melhorzincraft-back/internal/core/response"
)

// RegisterRoutes registra as rotas de templates no router
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
// @Summary Criar um novo template de jogo
// @Description Cria um novo template de jogo com o nome e docker compose especificados
// @Tags templates
// @Accept json
// @Produce json
// @Param request body dto.CreateRequest true "Dados do template a ser criado"
// @Success 200 {object} response.Response[dto.CreateResponse] "Template criado com sucesso"
// @Failure 400 {object} response.Response[any] "Requisição inválida"
// @Failure 500 {object} response.Response[any] "Erro interno do servidor"
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
// @Summary Atualizar um template de jogo existente
// @Description Atualiza um template de jogo com o nome e/ou docker compose especificados
// @Tags templates
// @Accept json
// @Produce json
// @Param id path int true "ID do Template"
// @Param request body dto.UpdateRequest true "Dados do template a ser atualizado"
// @Success 200 {object} response.Response[dto.UpdateResponse] "Template atualizado com sucesso"
// @Failure 400 {object} response.Response[any] "ID inválido ou requisição inválida"
// @Failure 404 {object} response.Response[any] "Template não encontrado"
// @Failure 500 {object} response.Response[any] "Erro interno do servidor"
// @Router /game/templates/{id} [put]
func (h *handler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
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
// @Summary Obter um template de jogo
// @Description Obtém um template de jogo pelo ID
// @Tags templates
// @Produce json
// @Param id path int true "ID do Template"
// @Success 200 {object} response.Response[dto.ShowResponse] "Template encontrado com sucesso"
// @Failure 400 {object} response.Response[any] "ID inválido"
// @Failure 404 {object} response.Response[any] "Template não encontrado"
// @Failure 500 {object} response.Response[any] "Erro interno do servidor"
// @Router /game/templates/{id} [get]
func (h *handler) Get(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
		return
	}

	req := dto.ShowRequest{Id: id}
	resp := h.templateUC.Show(c.Request.Context(), &req)
	c.JSON(resp.Code, resp)
}

// Delete godoc
// @Summary Excluir um template de jogo
// @Description Exclui um template de jogo pelo ID
// @Tags templates
// @Produce json
// @Param id path int true "ID do Template"
// @Success 200 {object} response.Response[dto.DeleteResponse] "Template excluído com sucesso"
// @Failure 400 {object} response.Response[any] "ID inválido"
// @Failure 404 {object} response.Response[any] "Template não encontrado"
// @Failure 500 {object} response.Response[any] "Erro interno do servidor"
// @Router /game/templates/{id} [delete]
func (h *handler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
		return
	}

	req := dto.DeleteRequest{Id: id}
	resp := h.templateUC.Delete(c.Request.Context(), &req)
	c.JSON(resp.Code, resp)
}
