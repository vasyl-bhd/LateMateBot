package commands

import (
	"LateMateBot/dao"
	"LateMateBot/model"
	"fmt"
	tele "gopkg.in/telebot.v4"
	"strings"
	"time"
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
	if (lastLatee == model.Latee{}) {
		return c.Send("Nothing to revert, you stewpead!")
	}
	lastLatee.IsDeleted = true
	lastLatee.DeletionReason = deletionReason
	lastLatee.RemovedAt = time.Now()
	dao.Save(&lastLatee)

	return c.Send(fmt.Sprintf("Ok, I removed late for %s at %s", lastLatee.Name,
		lastLatee.AddedAt.Local().Format("2006-01-02 15:04:05")))
}
