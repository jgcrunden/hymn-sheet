package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jgcrunden/hymn-sheet/model"
	"github.com/jgcrunden/hymn-sheet/service"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).With().Caller().Logger()

	fileReader := service.NewOSFileReader()
	parser := service.NewConfigParser(fileReader)
	parser.ReadConfigFile("./examples/mass-config.json")
	content, err := os.ReadFile("./resources/ordo.json")
	if err != nil {
		log.Error().Msg(fmt.Sprintf("%v", err))
	}

	var ordo model.Ordo
	err = json.Unmarshal(content, &ordo)
	if err != nil {
		fmt.Println("Error marshalling data", err)
	}

	/*
		proper, err := ordo.GetProper(conf.Date.String(), false)
		if err != nil {
			fmt.Println(err)
		}
		content, err = os.ReadFile("./resources/calendar.json")
		if err != nil {
			fmt.Println("error reading file", err)
		}

		var calendar model.Calendar
		err = json.Unmarshal(content, &calendar)
		if err != nil {
			fmt.Println(err)
		}

		propers, err := calendar.GetPropers(proper)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("propers are:", propers)

		service.DeriveSundayCycle(conf.Date)
	*/
}
