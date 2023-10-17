package loan

import (
	"net/http"
	"tes_project_siesta/init/loan"
	"tes_project_siesta/init/response"
	"tes_project_siesta/models/model_loan"

	"github.com/gin-gonic/gin"
)

func CreateLoan(c *gin.Context) {

	var response response.Response
	var request loan.Loan
	if err := c.ShouldBindJSON(&request); err != nil {

		response.Status = http.StatusBadRequest
		response.Message = err.Error()
		response.Data = nil
		c.JSON(http.StatusBadRequest, response)
		return
	}

	result, err := model_loan.CreateLoan(request)
	if err != nil {
		response.Status = http.StatusBadRequest
		response.Message = err.Error()
		response.Data = nil
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response.Status = http.StatusCreated
	response.Message = "Success Response"
	response.Data = result

	c.JSON(http.StatusCreated, response)
}

func CalculateLoanMonth(c *gin.Context) {

	var response response.Response

	idUser := c.Query("id_user")
	result, err := model_loan.CalculateLoanMonth(idUser)
	if err != nil {
		response.Status = http.StatusBadRequest
		response.Message = err.Error()
		response.Data = nil
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response.Status = http.StatusOK
	response.Message = "Success Response"
	response.Data = result

	c.JSON(http.StatusOK, response)
}