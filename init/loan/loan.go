package loan

type CalculateLoan struct {
	Id                int    `json:"id"`
	LoanDate          string `json:"loan_date"`
	Amount            int    `json:"amount"`
	Tenor             int    `json:"tenor"`
	Principal         int    `json:"principal"`
	PrincipalPerMonth int    `json:"principal_per_month"`
	Fee               int    `json:"fee"`
	InterestPerMonth  int    `json:"interest_per_month"`
	TotalInterest     int    `json:"total_interest"`
	FeeStamp          int    `json:"fee_stamp"`
	TotalPayment      int    `json:"total_payment"`
	DemandPerMonth    int    `json:"demand_per_month"`
}

type LoanMonth struct {
	IdUser             int `json:"id_user"`
	Month              int `json:"month"`
	Fee                int `json:"fee"`
	FeeStamp           int `json:"fee_stamp"`
	Interest           int `json:"interest"`
	PrincipalPaid      int `json:"principal_paid"`
	Bill               int `json:"bill"`
	RemainingPrincipal int `json:"remainingPrincipal"`
}

type LoanTotal struct {
	IdUser             int         `json:"id_user"`
	Loan               []LoanMonth `json:"loan"`
	Fee                int         `json:"fee"`
	FeeStamp           int         `json:"fee_stamp"`
	Interest           int         `json:"interest"`
	PrincipalPaid      int         `json:"principal_paid"`
	Bill               int         `json:"bill"`
	RemainingPrincipal int         `json:"remainingPrincipal"`
}

type Loan struct {
	IdUser  int  `json:"id_user"`
	Amount  int  `json:"amount"`
	Tenor   int  `json:"tenor"`
	Preview bool `json:"preview"`
}

type GetLoan struct {
	Id         int `json:"id"`
	IdUser     int `json:"id_user"`
	Amount     int `json:"amount"`
	Tenor      int `json:"tenor"`
	StartMonth int `json:"start_month"`
}
