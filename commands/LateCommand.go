package commands

import (
	"LateMateBot/dao"
	"LateMateBot/model"
	"LateMateBot/utils"
	"fmt"
	tele "gopkg.in/telebot.v4"
	"time"
)

func LateCommand(c tele.Context) error {
	mentions := utils.GetMentions(c.Entities(), c.Text())

	for _, mention := range mentions {
		latee := model.Latee{Name: mention, AddedAt: time.Now()}
		dao.Save(&latee)

	}
	err := c.Send(fmt.Sprintf("Oopsie! Someone is late: %s, wink wink ;)", mentions))

	return err
}
