package service

import (
	"LateMateBot/model"
	"LateMateBot/utils"
)

func GetRecentLatees(latees []model.Latee) (recentLatees []model.Latee) {
	return utils.Filter(latees, func(latee model.Latee) bool {
		return latee.AddedAt.After(utils.StartOfPastWeek())
	})
}
