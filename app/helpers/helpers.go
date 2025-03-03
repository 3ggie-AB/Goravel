package helpers

import (
	"fmt"

	"github.com/goravel/framework/facades"
)

// Asset: Fungsi mirip asset() di Laravel
func Asset(path string) string {
	config := facades.Config()
	baseURL := config.Env("APP_URL", "https://localhost:3000").(string)
	return baseURL + "public/" + path
}

func Rupiah(amount int) string {
	return fmt.Sprintf("Rp %d", amount)
}
