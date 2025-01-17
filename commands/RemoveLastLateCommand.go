package commands

import (
	"LateMateBot/dao"
	"fmt"
	tele "gopkg.in/telebot.v4"
	"strings"
)

func RemoveLastLateCommand(c tele.Context) error {
	var deletionReason string

	for _, e := range c.Entities() {
		if e.Type == tele.EntityCommand {
			deletionReason = c.Text()[e.Offset+e.Length:]
			break
		}
	}

	if strings.TrimSpace(deletionReason) == "" {
		return c.Send("Nope. Add some reason! (like so /oops I accidentally added him because I'm stewpead")
	}

	lastLatee := dao.GetLast()
	lastLatee.IsDeleted = true
	lastLatee.DeletionReason = deletionReason
	dao.Save(&lastLatee)

	return c.Send(fmt.Sprintf("Ok, I removed late for %s at %s", lastLatee.Name, lastLatee.AddedAt))
}
