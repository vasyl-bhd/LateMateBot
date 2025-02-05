package commands

import (
	"LateMateBot/dao"
	"LateMateBot/utils"
	"fmt"
	tele "gopkg.in/telebot.v4"
)

func GetAllCommand(c tele.Context) error {
	mentions := utils.GetMentions(c.Entities(), c.Text())
	latees := getLatees(mentions, dao.GetAll())

	response := getStarterResponse(mentions)

	for _, latee := range latees {
		response += fmt.Sprintf("Name: %s was late at: %s\n",
			latee.Name,
			latee.AddedAt.Local().Format("2006-01-02 15:04:05"))
	}

	err := c.Send(response)

	return err
}

func getStarterResponse(mentions []string) string {
	if len(mentions) == 0 {
		return "Here's a list of latees for you(like all of them):\n"
	}

	return fmt.Sprintf("Here's a list of all lates for %v:\n", mentions)
}
