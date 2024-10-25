package utilities

import (
	"encoding/csv"
	"fmt"
	"os"
	"tenders/models"
)

func FileReader() ([][]string, error) {
	file, err := os.Open("data/published_contracts.csv")
	if err != nil {
		fmt.Println("Error opening published contracts file")
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	var ContractsData [][]string
	ContractsData, err = reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading csv data")
		return nil, err
	}

	return ContractsData, nil
}

func ContractsMapper(s [][]string) models.Contracts {
	var ContractsData models.Contracts
	var contractStruct models.Contract
	for _, val := range s {
		contractStruct.ContractNumber = val[0]
		contractStruct.Amount = val[1]
		contractStruct.AwardDate = val[2]
		contractStruct.TenderTitle = val[3]
		contractStruct.EvalCompletionDate = val[4]
		contractStruct.NotificationOfAwardDate = val[5]
		contractStruct.SignDate = val[6]
		contractStruct.StartDate = val[7]
		contractStruct.EndDate = val[8]
		contractStruct.AgpoCertificateNumber = val[9]
		contractStruct.AwardedAgpoGroupId = val[10]
		contractStruct.CreatedBy = val[11]
		contractStruct.Terminated = val[12]
		contractStruct.FinancialYear = val[13]
		contractStruct.Quater = val[14]
		contractStruct.TenderRef = val[15]
		contractStruct.PEName = val[16]
		contractStruct.SupplierName= val[17]
		contractStruct.NoOfBOI = val[18]
		contractStruct.CreatedAt = val[19]

		ContractsData = append(ContractsData, contractStruct)
	}
	return ContractsData
}
