package routes

import (
	"fmt"
	"goravel/app/http/controllers"
	"strings"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

func Web() {
	facades.Route().Get("/public/{file}", func(ctx http.Context) http.Response {
		originalFile := ctx.Request().Input("file")
		fmt.Println("Original file:", originalFile)
		newFile := strings.ReplaceAll(originalFile, "-", "/")
		fmt.Println("Modified file path:", newFile)
		filePath := "public/monica/" + newFile
		fmt.Println("Final file path:", filePath)
		return ctx.Response().File(filePath)
	})

	landingController := controllers.NewLandingController()
	facades.Route().Get("/", landingController.BlogIndex)
	facades.Route().Get("/about", landingController.BlogAbout)
	facades.Route().Get("/blog", landingController.BlogBlog)
	facades.Route().Get("/contact", landingController.BlogContact)
	facades.Route().Get("/generic", landingController.BlogGeneric)
	facades.Route().Get("/services", landingController.BlogServices)
	facades.Route().Get("/single", landingController.BlogSingle)
	facades.Route().Get("/styles", landingController.BlogStyles)
}
