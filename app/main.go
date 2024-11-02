package main

import (
	"log"
	"github.com/Resolution-hash/s1-report/config"
	"github.com/Resolution-hash/s1-report/internal/bot"
)

func main() {
	// totalTicket, err := api.GetStatistics()
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(totalTicket)

	botCfg, err := config.LoadBOTConfig()
	if err != nil {
		log.Fatalln(err)
	}
	bot.InitBOT(botCfg)

}
