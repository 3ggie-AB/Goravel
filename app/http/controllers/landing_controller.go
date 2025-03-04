package controllers

import (
	"goravel/app/helpers"

	"github.com/goravel/framework/contracts/http"
)

type LandingController struct{}

func NewLandingController() *LandingController {
	return &LandingController{}
}

func (r *LandingController) BlogIndex(ctx http.Context) http.Response {
	return ctx.Response().View().Make("blog/index.tmpl", map[string]any{
		"assets": helpers.Asset(""),
	})
}

func (r *LandingController) BlogBlog(ctx http.Context) http.Response {
	return ctx.Response().View().Make("blog/blog.tmpl", map[string]any{
		"assets": helpers.Asset(""),
	})
}

func (r *LandingController) BlogAbout(ctx http.Context) http.Response {
	return ctx.Response().View().Make("blog/about.tmpl", map[string]any{
		"assets": helpers.Asset(""),
	})
}

func (r *LandingController) BlogContact(ctx http.Context) http.Response {
	return ctx.Response().View().Make("blog/contact.tmpl", map[string]any{
		"assets": helpers.Asset(""),
	})
}

func (r *LandingController) BlogGeneric(ctx http.Context) http.Response {
	return ctx.Response().View().Make("blog/generic.tmpl", map[string]any{
		"assets": helpers.Asset(""),
	})
}

func (r *LandingController) BlogServices(ctx http.Context) http.Response {
	return ctx.Response().View().Make("blog/services.tmpl", map[string]any{
		"assets": helpers.Asset(""),
	})
}

func (r *LandingController) BlogSingle(ctx http.Context) http.Response {
	return ctx.Response().View().Make("blog/services.tmpl", map[string]any{
		"assets": helpers.Asset(""),
	})
}

func (r *LandingController) BlogStyles(ctx http.Context) http.Response {
	return ctx.Response().View().Make("blog/styles.tmpl", map[string]any{
		"assets": helpers.Asset(""),
	})
}
