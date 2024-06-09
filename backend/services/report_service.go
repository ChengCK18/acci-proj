package services

import (
	"acci-backend/repositories"
)

func GetReportService(id uint) (interface{}, error) {
	reports, err := repositories.FindReportById(id)
	if err != nil {
		return nil, err
	}

	return reports, nil
}
