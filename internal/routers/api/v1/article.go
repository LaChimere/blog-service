package v1

import (
	"github.com/lachimere/blog-service/global"
	"github.com/lachimere/blog-service/internal/service"
	"github.com/lachimere/blog-service/pkg/app"
	"github.com/lachimere/blog-service/pkg/convert"
	"github.com/lachimere/blog-service/pkg/errcode"

	"github.com/gin-gonic/gin"
)

type Article struct{}

func NewArticle() Article {
	return Article{}
}

// Get godoc
// @Summary Get an article
// @Produce json
// @Param id path int true "Article ID"
// @Success 200 {object} model.Article "Success"
// @Failure 400 {object} errcode.Error "request error"
// @Failure 500 {object} errcode.Error "internal error"
// @Router /api/v1/articles/{id} [get]
func (a Article) Get(c *gin.Context) {
	param := service.ArticleRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	article, err := svc.GetArticle(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetArticle err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetArticleFail)
		return
	}

	response.ToResponse(article)
	return
}

// List godoc
// @Summary List articles
// @Produce json
// @Param name query string false "Article Name"
// @Param tag_id query int false "Tag ID"
// @Param state query int false "State"
// @Param page query int false "Page"
// @Param page_size query int false "Page Size"
// @Success 200 {object} model.ArticleSwagger "Success"
// @Failure 400 {object} errcode.Error "request error"
// @Failure 500 {object} errcode.Error "internal error"
// @Router /api/v1/articles [get]
func (a Article) List(c *gin.Context) {
	param := service.ArticleListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	articles, totalRows, err := svc.GetArticleList(&param, &pager)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetArticleList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetArticlesFail)
		return
	}

	response.ToResponseList(articles, totalRows)
	return
}

// Create godoc
// @Summary Create article
// @Produce json
// @Param tag_id body string true "Tag ID"
// @Param title body string true "Article Title"
// @Param desc body string false "Article Description"
// @Param cover_image_url body string true "Cover Image URL"
// @Param content body string true "Article Content"
// @Param created_by body int true "Creator"
// @Param state body int false "State"
// @Success 200 {object} model.Article "Success"
// @Failure 400 {object} errcode.Error "request error"
// @Failure 500 {object} errcode.Error "internal error"
// @Router /api/v1/articles [post]
func (a Article) Create(c *gin.Context) {
	param := service.CreateArticleRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CreatedArticle(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.CreateArticle err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateArticleFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}

// Update godoc
// @Summary Update article
// @Produce json
// @Param tag_id body string false "Tag ID"
// @Param title body string false "Article Title"
// @Param desc body string false "Article Description"
// @Param cover_image_url body string false "Cover Image URL"
// @Param content body string false "Article Content"
// @Param modified_by body string true "Modifier"
// @Success 200 {object} model.Article "Success"
// @Failure 400 {object} errcode.Error "request error"
// @Failure 500 {object} errcode.Error "internal error"
// @Router /api/v1/articles/{id} [put]
func (a Article) Update(c *gin.Context) {
	param := service.UpdateArticleRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.UpdateArticle(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.UpdateArticle err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateArticleFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}

// Delete godoc
// @Summary Delete article
// @Produce  json
// @Param id path int true "Article ID"
// @Success 200 {string} string "Success"
// @Failure 400 {object} errcode.Error "request error"
// @Failure 500 {object} errcode.Error "internal error"
// @Router /api/v1/articles/{id} [delete]
func (a Article) Delete(c *gin.Context) {
	param := service.DeleteArticleRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.DeleteArticle(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.DeleteArticle err: %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteArticleFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}
