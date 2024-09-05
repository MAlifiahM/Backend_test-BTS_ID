package main

import (
	"Intersolusi_Teknologi_Asia/internal/checklist"
	"Intersolusi_Teknologi_Asia/internal/infrastructure"
	jwt "Intersolusi_Teknologi_Asia/internal/middleware"
	"Intersolusi_Teknologi_Asia/internal/user"
)

func main() {
	app := infrastructure.SetupFiberApp()

	db := infrastructure.ConnectDB()

	userRepo := user.NewUserRepository(db)
	checklistRepo := checklist.NewChecklistRepository(db)

	userHandler := user.NewUserHandler(userRepo)
	checklistHandler := checklist.NewChecklistHandler(checklistRepo)

	// user
	app.Post("/register", userHandler.Register)
	app.Post("/login", userHandler.Login)

	app.Use(jwt.JWTMiddleware())

	// checklist
	app.Post("/checklist", checklistHandler.CreateChecklist)
	app.Get("/checklist", checklistHandler.GetChecklists)
	app.Delete("/checklist/:id", checklistHandler.DeleteChecklist)

	// checklist Item
	app.Post("/checklist/:id/item", checklistHandler.AddItem)
	app.Get("/checklist/:id/item", checklistHandler.GetItems)
	app.Put("/checklist/:id/item/:idItem", checklistHandler.UpdateItem)
	app.Delete("/checklist/:id/item/:idItem", checklistHandler.DeleteItem)
	app.Put("/checklist/:id/item/:idItem/status", checklistHandler.UpdateItemStatus)

	app.Listen(":8080")
}
