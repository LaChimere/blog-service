package v1

import (
	"github.com/lachimere/blog-service/global"
	"github.com/lachimere/blog-service/internal/service"
	"github.com/lachimere/blog-service/pkg/app"
	"github.com/lachimere/blog-service/pkg/convert"
	"github.com/lachimere/blog-service/pkg/errcode"

	"github.com/gin-gonic/gin"
)

type Tag struct{}

func NewTag() Tag {
	return Tag{}
}

// List godoc
// @Summary List tags
// @Produce  json
// @Param name query string false "Tag Name" maxlength(100)
// @Param state query int false "State" Enums(0, 1) default(1)
// @Param page query int false "Page"
// @Param page_size query int false "Page Size"
// @Success 200 {object} model.TagSwagger "Success"
// @Failure 400 {object} errcode.Error "request error"
// @Failure 500 {object} errcode.Error "internal error"
// @Router /api/v1/tags [get]
func (t Tag) List(c *gin.Context) {
	param := service.ListTagRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	totalRows, err := svc.CountTag(&service.CountTagRequest{Name: param.Name, State: param.State})
	if err != nil {
		global.Logger.Errorf(c, "svc.CountTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorCountTagFail)
		return
	}
	tags, err := svc.GetTagList(&param, &pager)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetTagList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetTagListFail)
		return
	}

	response.ToResponseList(tags, totalRows)
	return
}

// Create godoc
// @Summary Create tag
// @Produce  json
// @Param name body string true "Tag Name" minlength(3) maxlength(100)
// @Param state body int false "State" Enums(0, 1) default(1)
// @Param created_by body string false "Creator" minlength(3) maxlength(100)
// @Success 200 {object} model.Tag "Success"
// @Failure 400 {object} errcode.Error "request error"
// @Failure 500 {object} errcode.Error "internal error"
// @Router /api/v1/tags [post]
func (t Tag) Create(c *gin.Context) {
	param := service.CreateTagRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CreateTag(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.CreateTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateTagFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}

// Update godoc
// @Summary Update tag
// @Produce  json
// @Param id path int true "Tag ID"
// @Param name body string false "Tag Name" minlength(3) maxlength(100)
// @Param state body int false "State" Enums(0, 1) default(1)
// @Param modified_by body string true "Modifier" minlength(3) maxlength(100)
// @Success 200 {array} model.Tag "Success"
// @Failure 400 {object} errcode.Error "request error"
// @Failure 500 {object} errcode.Error "internal error"
// @Router /api/v1/tags/{id} [put]
func (t Tag) Update(c *gin.Context) {
	param := service.UpdateTagRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.UpdateTag(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.UpdateTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateTagFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}

// Delete godoc
// @Summary Delete tag
// @Produce  json
// @Param id path int true "Tag ID"
// @Success 200 {string} string "Success"
// @Failure 400 {object} errcode.Error "request error"
// @Failure 500 {object} errcode.Error "internal error"
// @Router /api/v1/tags/{id} [delete]
func (t Tag) Delete(c *gin.Context) {
	param := service.DeleteTagRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.DeleteTag(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.DeleteTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteTagFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}
