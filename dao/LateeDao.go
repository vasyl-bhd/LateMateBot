package dao

import (
	"LateMateBot/model"
	"fmt"
	"log"
)

func Save(latee *model.Latee) {
	log.Printf("Saving latee: %+v\n", latee)
	DB.Save(&latee)
}

func GetAll() []model.Latee {
	var latees []model.Latee
	DB.Model(&model.Latee{}).Order("added_at desc").Where("is_deleted = false").Find(&latees)
	fmt.Printf("%+v\n", latees)

	return latees
}
func GetLast() (result model.Latee) {
	DB.Model(&model.Latee{}).Limit(1).Order("added_at desc").Where("is_deleted = false").Find(&result)
	return
}
