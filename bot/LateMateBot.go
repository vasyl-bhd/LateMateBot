package bot

import (
	"LateMateBot/commands"
	"fmt"
	tele "gopkg.in/telebot.v4"
	"log"
	"time"
)

func InitAndStartBot() {
	pref := tele.Settings{
		Token:  "8094075135:AAH84a_BWdReSLCZ2kEGcW45LWfFivoTbr0",
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/late", commands.LateCommand)

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
