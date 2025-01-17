package commands

import (
	"LateMateBot/dao"
	"fmt"
	tele "gopkg.in/telebot.v4"
)

func GetRemovedCommand(c tele.Context) error {
	removed := dao.GetRemoved()

	if len(removed) == 0 {
		return c.Send("'Noone' is removed, calm down")
	}

	var response = "Here's a list of removed lads:\n"

	for _, r := range removed {
		response += fmt.Sprintf("Name:%s; Removed at: %s; Reason: %s\n",
			r.Name, r.RemovedAt.Local().Format("2006-01-02 15:04:05"), r.DeletionReason)
	}

	return c.Send(response)
}
