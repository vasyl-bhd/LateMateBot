package commands

import (
	tele "gopkg.in/telebot.v4"
)

func GetHelp(context tele.Context) error {
	response := "Available commands:\n" +
		"- /late <at-mention> - log person as late\n" +
		"- /all <at-mentions>(optional) - get all stats for a person(use mention to get stats only for those people)\n" +
		"- /recent <at-mentions>(optional) - get recent stats for a person(use mention to get stats only for those people)\n" +
		"- /oops <reason> - Use it when you fucked up and logged late accidentally"

	return context.Send(response)
}
