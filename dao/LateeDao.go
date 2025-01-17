package dao

import (
	"LateMateBot/model"
	"fmt"
)

func Save(latee *model.Latee) {
	DB.Save(&latee)
}

func GetAll() []model.Latee {
	var latees []model.Latee
	DB.Find(&latees).Order("AddedAt desc")
	fmt.Printf("%+v\n", latees)

	return latees
}
func GetLast() (result model.Latee) {
	DB.Last(result)
	return
}
