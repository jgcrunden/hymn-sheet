package service

import (
	"encoding/json"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/jgcrunden/hymn-sheet/model"
)

type ProperBuilder struct {
	reader           FileReader
	properDay        string
	suppressedSunday string
	config           model.Config
}

func NewProperBuilder(reader FileReader, config model.Config) ProperBuilder {
	return ProperBuilder{reader: reader, config: config}
}

func (p *ProperBuilder) GetOrdo(filename string) error {
	content, err := p.reader.ReadFile(filename)
	if err != nil {
		return err
	}

	var ordo model.Ordo
	err = json.Unmarshal(content, &ordo)
	if err != nil {
		return err
	}
	p.properDay, err = ordo.GetProperDay(p.config.Date, false)
	if err != nil {
		return err
	}
	p.suppressedSunday, err = ordo.GetSuppressedSunday(p.config.Date)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProperBuilder) GetPropers(filename string) (model.Propers, error) {
	content, err := p.reader.ReadFile(filename)
	if err != nil {
		return model.Propers{}, err
	}

	var c model.Calendar
	err = json.Unmarshal(content, &c)
	if err != nil {
		return model.Propers{}, err
	}

	propers, err := c.GetPropers(p.properDay)
	if err != nil {
		return model.Propers{}, err
	}
	return propers, nil
}

func (p *ProperBuilder) DeriveCycles() (model.Cycles, error) {
	cycles := model.Cycles{}
	switch p.config.Date.Year() % 3 {
	case 0:
		cycles.LectionaryYearSunday = model.C
	case 1:
		cycles.LectionaryYearSunday = model.A
	case 2:
		cycles.LectionaryYearSunday = model.B
	}

	switch p.config.Date.Year() % 2 {
	case 0:
		cycles.LectionaryYearWeekday = 2
	case 1:
		cycles.LectionaryYearWeekday = 1
	}
	properDayTmp := p.properDay

	startOfSeason := []string{"holy-family", "baptism-of-the-lord"}
	if slices.Contains(startOfSeason, p.properDay) {
		cycles.PsalterWeek = 1
		return cycles, nil
	}

	if p.suppressedSunday != "" {
		properDayTmp = p.suppressedSunday
	}

	split := strings.Split(properDayTmp, "-")
	finalSplit := split[len(split)-1]

	weekNum, err := strconv.Atoi(finalSplit)
	if err != nil {
		return model.Cycles{}, fmt.Errorf("cannot derive psalter week from %s: no start of the season, does not have a numbered week and does not have suppressed sunday info", properDayTmp)
	}

	psalterWeek := weekNum % 4
	if psalterWeek == 0 {
		psalterWeek = 4
	}
	cycles.PsalterWeek = uint(psalterWeek)
	return cycles, nil
}
