package model

import (
	"encoding/json"
	"fmt"
)

type calendar map[string]*Propers
type Calendar struct {
	internal calendar
}

type Propers struct {
	ProperDay         string `json:"-"`
	EntranceAntiphon  string `json:"entranceAntiphon"`
	FirstReading      string `json:"firstReading"`
	ResponsorialPsalm string `json:"responsorialPsalm"`
	SecondReading     string `json:"secondReading"`
	GospelAcclamation string `json:"gospelAcclamation"`
	Gospel            string `json:"gospel"`
	CommunionAntiphon string `json:"communionAntiphon"`
}

type Cycles struct {
	LectionaryYearSunday  Year
	LectionaryYearWeekday uint
	PsalterWeek           uint
}

func (c *Calendar) UnmarshalJSON(b []byte) error {
	var tmp calendar
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	c.internal = tmp
	return nil
}

func (c Calendar) GetPropers(proper string) (Propers, error) {
	if c.internal[proper] == nil {
		return Propers{}, fmt.Errorf("could not find propers for %s", proper)
	} else {
		return *c.internal[proper], nil
	}
}
