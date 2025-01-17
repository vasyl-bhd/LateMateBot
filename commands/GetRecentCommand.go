package commands

import (
	"LateMateBot/dao"
	"LateMateBot/model"
	"LateMateBot/service"
	"LateMateBot/utils"
	"fmt"
	tele "gopkg.in/telebot.v4"
	"slices"
)

func GetRecentCommand(c tele.Context) error {
	mentions := utils.GetMentions(c.Entities(), c.Text())
	recentLatees := getLatees(mentions, service.GetRecentLatees(dao.GetAll()))

	response := getRecentStarterResponse(mentions)

	for _, latee := range recentLatees {
		response += fmt.Sprintf("Name: %s was late at: %s\n",
			latee.Name,
			latee.AddedAt.Local().Format("2006-01-02 15:04:05"))
	}

	err := c.Send(response)

	return err
}

func getLatees(mentions []string, latees []model.Latee) []model.Latee {
	if len(mentions) != 0 {
		return utils.Filter(latees, func(latee model.Latee) bool {
			return slices.Contains(mentions, latee.Name)
		})
	}

	return latees
}

func getRecentStarterResponse(mentions []string) string {
	if len(mentions) == 0 {
		return "Here's a list of the **recent** latees for you:\n"
	}

	return fmt.Sprintf("Here's a list of **recent** lates for %v:\n", mentions)
}
