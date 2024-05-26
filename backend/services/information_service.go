package services

import "acci-backend/repositories"

func GetInformationService(id uint) (interface{}, error) {
	info, err := repositories.FindInformationById(id)
	if err != nil {
		return nil, err
	}
	return info, nil
}
