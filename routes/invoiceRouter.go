package routes

import "github.com/gin-gonic/gin"
import controller "go-resto-management/controllers"

func InvoiceRoutes(router *gin.Engine) {
	router.GET("/api/invoices", controller.GetInvoices())
	router.GET("/api/invoices/:invoices_id", controller.GetInvoice())
	router.POST("/api/invoices", controller.CreateInvoices())
	router.PATCH("/api/invoices/:invoices_id", controller.UpdateInvoices())
}
