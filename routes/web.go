package routes

import (
	"fmt"
	"goravel/app/helpers"
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
	facades.Route().Get("cukai/{name}", func(ctx http.Context) http.Response {
		// Get the "name" parameter from the URL
		name := ctx.Request().Input("name")

		// Return the view with the template file name
		return ctx.Response().View().Make(name+".tmpl", map[string]any{
			"assets": helpers.Asset(""),
		})
	})
}
