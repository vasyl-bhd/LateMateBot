package dao

import (
	"LateMateBot/model"
)

func Save(latee *model.Latee) {
	DB.Create(&latee)
}

func GetAll() []model.Latee {
	var latees []model.Latee
	DB.Find(&latees)

	return latees

}
