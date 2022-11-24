package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events: ", event.Timestamp, event.Command, event.Parameters, event.Event.Channel)
	}
}
func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-3840796681478-4443250939056-BjaJ8kf5d1bUSuogtZKoLvmc")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A04BWUYPLEB-4416583248389-21bb29ffae77694813c1426076cb14e857e57cc64b99ddcfe3a6310726be306c")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("my year of birth is  <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		Examples:    []string{"hi", "YOB"},
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {

			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				println("error")
			}
			age := 2022 - yob
			r := fmt.Sprintf("age is %d", age)
			response.Reply(r)

		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
