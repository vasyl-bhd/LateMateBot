package main

import (
	"LateMateBot/bot"
	"LateMateBot/dao"
)

func main() {
	dao.Init()

	bot.InitAndStartBot()

}
