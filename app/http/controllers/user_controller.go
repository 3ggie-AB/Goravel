package controllers

import (
	"fmt"
	"goravel/app/models"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

// ✅ CREATE: Tambah User Baru
func (r *UserController) Store(ctx http.Context) http.Response {
	// ✅ 1. Definisikan aturan validasi
	rules := map[string]string{
		"name":     "required|min_len:3",
		"email":    "required|email|unique:users,email",
		"password": "required|min_len:6",
		"nik":      "required|numeric|digits:16",
	}

	// ✅ 2. Jalankan validasi dengan facades.Validation()
	validator, err := facades.Validation().Make(ctx.Request().All(), rules)
	if err != nil {
		fmt.Println("Error validating user:", err)
		return ctx.Response().Json(500, http.Json{"error": "Failed to validate user", "message": err.Error()})
	}
	// ✅ 3. Cek apakah validasi gagal
	if validator.Fails() {
		return ctx.Response().Json(400, http.Json{
			"error":   "Validation failed",
			"message": validator.Errors().All(), // ✅ Ambil semua error
		})
	}

	// ✅ 4. Hash password sebelum menyimpan
	hashedPassword, _ := facades.Hash().Make(ctx.Request().Input("password"))

	// ✅ 5. Simpan data user
	user := models.User{
		Name:     ctx.Request().Input("name"),
		Email:    ctx.Request().Input("email"),
		Password: hashedPassword,
		Role:     "user",
		Nik:      ctx.Request().Input("nik"),
	}

	err = facades.Orm().Query().Create(&user)
	if err != nil {
		fmt.Println("Error creating user:", err)
		return ctx.Response().Json(500, http.Json{"error": "Failed to create user", "message": err.Error()})
	}

	return ctx.Response().Json(201, http.Json{
		"success": true,
		"message": "User created successfully",
		"data":    user,
	})
}

// ✅ READ: Ambil Semua User
func (r *UserController) Index(ctx http.Context) http.Response {
	var users []models.User

	err := facades.Orm().Query().Model(&users).Get(&users)
	if err != nil {
		fmt.Println("Error fetching users:", err)
		return ctx.Response().Json(500, http.Json{"error": "Failed to fetch users"})
	}

	return ctx.Response().Json(200, http.Json{
		"success": true,
		"users":   users,
	})
}

// ✅ READ: Ambil User Berdasarkan ID
func (r *UserController) Show(ctx http.Context) http.Response {
	var user models.User
	userID := ctx.Request().InputInt("id") // ✅ Pastikan ID bertipe int

	err := facades.Orm().Query().Model(&user).Where("id", userID).First(&user)
	if err != nil {
		fmt.Println("Error fetching user:", err)
		return ctx.Response().Json(404, http.Json{"error": "User not found"})
	}

	return ctx.Response().Json(200, http.Json{
		"success": true,
		"data":    user,
	})
}

// ✅ UPDATE: Edit User Berdasarkan ID
func (r *UserController) Update(ctx http.Context) http.Response {
	var user models.User
	userID := ctx.Request().InputInt("id") // ✅ Gunakan InputInt agar pasti integer

	// Cek apakah user ada
	err := facades.Orm().Query().Model(&user).Where("id", userID).First(&user)
	if err != nil {
		fmt.Println("User not found:", err)
		return ctx.Response().Json(404, http.Json{"error": "User not found"})
	}

	// Data yang akan diperbarui
	updateData := map[string]any{
		"name":  ctx.Request().Input("name"),
		"email": ctx.Request().Input("email"),
	}

	// ✅ Simpan jumlah baris yang diperbarui dan error
	affectedRows, err := facades.Orm().Query().Model(&user).Where("id", userID).Update(updateData)
	if err != nil {
		fmt.Println("Error updating user:", err)
		return ctx.Response().Json(500, http.Json{"error": "Failed to update user"})
	}

	// ✅ Cek apakah ada data yang diperbarui
	if affectedRows == nil { // ✅ Ubah dari `nil` ke `0`
		return ctx.Response().Json(400, http.Json{"error": "No changes made"})
	}

	return ctx.Response().Json(200, http.Json{
		"success": true,
		"message": "User updated successfully",
		"data":    updateData,
	})
}

// ✅ DELETE: Hapus User Berdasarkan ID
func (r *UserController) Destroy(ctx http.Context) http.Response {
	userID := ctx.Request().InputInt("id") // ✅ Gunakan InputInt agar ID bertipe int

	// Hapus user
	affectedRows, err := facades.Orm().Query().Model(&models.User{}).Where("id", userID).Delete()
	if err != nil {
		fmt.Println("Error deleting user:", err)
		return ctx.Response().Json(500, http.Json{"error": "Failed to delete user"})
	}

	if affectedRows == nil { // ✅ Ubah dari `nil` ke `0`
		return ctx.Response().Json(404, http.Json{"error": "User not found"})
	}

	return ctx.Response().Json(200, http.Json{
		"success": true,
		"message": "User deleted successfully",
	})
}
