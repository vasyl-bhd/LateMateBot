package main

import (
	"LateMateBot/commands"
	"fmt"
	tele "gopkg.in/telebot.v4"
	"log"
	"time"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
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
