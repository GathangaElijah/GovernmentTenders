package models

type Contract struct{
	ContractNumber string
	Amount string
	AwardDate string
	TenderTitle string
	EvalCompletionDate string
	NotificationOfAwardDate string
	SignDate string
	StartDate string
	EndDate string
	AgpoCertificateNumber string
	AwardedAgpoGroupId string
	CreatedBy string
	Terminated string
	FinancialYear string
	Quater string
	TenderRef string
	PEName string
	SupplierName string
	NoOfBOI string
	CreatedAt string
}

type Contracts []Contract