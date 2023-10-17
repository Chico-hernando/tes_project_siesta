package routes

import (
	"tes_project_siesta/controllers/loan"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) *gin.Engine {

	r.GET("/loan", loan.CalculateLoanMonth)
	r.POST("/loan", loan.CreateLoan)

	return r
}
