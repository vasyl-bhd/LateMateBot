package bot

import (
	"LateMateBot/commands"
	"fmt"
	tele "gopkg.in/telebot.v4"
	"log"
	"os"
	"time"
)

func InitAndStartBot() {
	pref := tele.Settings{
		Token:  os.Getenv("BOT_TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/late", commands.LateCommand)
	b.Handle("/all", commands.GetAllCommand)
	b.Handle("/recent", commands.GetRecentCommand)
	b.Handle("/oops", commands.RemoveLastLateCommand)
	b.Handle("/get_oopsies", commands.GetRemovedCommand)
	b.Handle("/help", commands.GetHelp)

	b.Handle(tele.OnText, func(c tele.Context) error {
		err := c.Send(c.Text())
		if err != nil {
			fmt.Println(err)
		}

		return nil
	})

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	b.Start()
}
