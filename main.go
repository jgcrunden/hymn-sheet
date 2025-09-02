package main

import (
	"flag"
	"os"
	"os/exec"

	"github.com/jgcrunden/hymn-sheet/service"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type OSFileReader struct{}

func (o OSFileReader) ReadFile(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}
func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).With().Caller().Logger()

	filename := flag.String("file", "config.json", "The JSON config file containing hymn date, configuration and hymn references")
	flag.Parse()

	osFileReader := OSFileReader{}
	configParser := service.NewConfigParser(osFileReader)
	err := configParser.ReadConfigFile(*filename)
	if err != nil {
		log.Error().Msg(err.Error())
		return
	}

	properBuilder := service.NewProperBuilder(osFileReader, configParser.Config)

	err = properBuilder.GetOrdo("./resources/ordo.json")
	if err != nil {
		log.Error().Msg(err.Error())
		return
	}

	cycles, err := properBuilder.DeriveCycles()
	if err != nil {
		log.Error().Msg(err.Error())
		return
	}

	propers, err := properBuilder.GetPropers("./resources/calendar.json")
	if err != nil {
		log.Error().Msg(err.Error())
		return
	}

	hymnBuilder := service.NewHymnBuilder(osFileReader, configParser.Config)
	hymns, err := hymnBuilder.GetHymns()
	if err != nil {
		log.Error().Msg(err.Error())
		return
	}

	configParser.Config.Hymns, err = hymnBuilder.TagHymnVerses(hymns)
	if err != nil {
		log.Error().Msg(err.Error())
		return
	}

	latexFile, err := service.GenerateLatex(configParser.Config, propers, cycles)
	if err != nil {
		log.Error().Msg(err.Error())
		return
	}

	cmd := exec.Command("lualatex", latexFile)
	if err := cmd.Run(); err != nil {
		log.Fatal().Msg(err.Error())
	}
}
