package model

type TopBorrower struct {
	AuthorID     uint   `json:"author_id"`
	AuthorName   string `json:"author_name"`
	BorrowerID   uint   `json:"borrower_id"`
	BorrowerName string `json:"borrower_name"`
	LoanCount    int    `json:"loan_count"`
}
