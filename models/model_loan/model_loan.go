package model_loan

import (
	"log"
	"math"
	"tes_project_siesta/configs"
	"tes_project_siesta/init/loan"
	"tes_project_siesta/utils"
)

func CreateLoan(request loan.Loan) (response loan.CalculateLoan, err error) {

	DB := configs.DB
	var totalLoan int

	tx, err := DB.Begin()
	if err != nil {
		return response, err
	}

	query := `select sum(amount) from loan where id_user = ?`

	DB.QueryRow(query, request.IdUser).Scan(&totalLoan)
	response = CalculateLoan(request, totalLoan+request.Amount)

	if request.Preview {
		return
	}

	queryInsert := `insert into loans (id_user, amount, tenor, created_at, updated_at)
					values(?, ?, ?, ?, ?)`

	sql, err := tx.Exec(queryInsert, request.IdUser, request.Amount, request.Tenor, utils.DateTime(), utils.DateTime())

	if err != nil {
		tx.Rollback()
		return
	}

	idtarget, err := sql.LastInsertId()

	if err != nil {
		tx.Rollback()
		return
	}

	response.Id = int(idtarget)

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return
	}

	return
}

func CalculateLoanMonth(idUser string) (result loan.LoanTotal, err error) {

	DB := configs.DB
	var getLoan loan.GetLoan
	// var arGetLoan []loan.GetLoan

	query := `select id_user, amount, tenor, month(created_at) from loans where id_user = ?`

	rows, err := DB.Query(query, idUser)

	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {

		err = rows.Scan(&getLoan.IdUser, &getLoan.Amount, &getLoan.Tenor, &getLoan.StartMonth)
		if err != nil {
			return
		}

		loan := CalculateLoan(loan.Loan{
			Amount: getLoan.Amount,
			Tenor:  getLoan.Tenor,
		}, 10)

		log.Printf("%+v\n",loan)

	}

	return
}

func CalculateLoan(loan loan.Loan, totalLoan int) (result loan.CalculateLoan) {

	result.Amount = loan.Amount
	result.Tenor = loan.Tenor
	result.LoanDate = utils.Date()
	result.Principal = loan.Amount
	result.PrincipalPerMonth = int(math.Ceil(float64(result.Principal) / float64(result.Tenor)))
	result.Fee = int(float64(result.Principal) * 0.05)
	result.InterestPerMonth = int(float64(result.Principal) * 0.0199)
	result.TotalInterest = result.InterestPerMonth * 3
	if totalLoan > 5000000 {
		result.FeeStamp = 10000
	}
	result.TotalPayment = result.Principal + result.TotalInterest
	result.DemandPerMonth = int(math.Ceil(float64(result.TotalPayment) / float64(result.Tenor)))

	return
}

func CalculateLoanTotal(getLoan loan.GetLoan, loanTotal *loan.LoanTotal) {

	loanTotal.IdUser = getLoan.IdUser
	loan := CalculateLoan(loan.Loan{
		Amount: getLoan.Amount,
		Tenor:  getLoan.Tenor,
	}, 10)

	loanTotal.Fee = loanTotal.Fee + loan.Fee
	loanTotal.FeeStamp = loanTotal.FeeStamp + loan.FeeStamp
	loanTotal.Interest = loanTotal.Interest + loan.TotalInterest

}
