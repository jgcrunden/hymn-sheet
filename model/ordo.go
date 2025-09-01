package model

import (
	"encoding/json"
	"errors"
)

type ordo map[string]map[string]string

type Ordo struct {
	internal ordo
}

func (o *Ordo) UnmarshalJSON(b []byte) error {
	var tmp ordo
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	o.internal = tmp
	return nil
}

func (o Ordo) GetProperDay(date Date, optional bool) (string, error) {
	search := "_"
	if optional {
		search = "optional"

	}
	d := o.internal[date.String()]
	if d == nil {
		return "", errors.New("date does not exist in calendar")
	}

	if d[search] == "" {
		return "", errors.New("no optional memorial available for today")
	}

	return d[search], nil
}

func (o Ordo) GetSuppressedSunday(date Date) (string, error) {
	d := o.internal[date.String()]
	if d == nil {
		return "", errors.New("date does not exist in calendar")
	}

	return d["suppressed"], nil
}

type Year int

const (
	A Year = iota + 1
	B
	C
)

var yearToString = map[Year]string{
	A: "A",
	B: "B",
	C: "C",
}

func (y Year) String() string {
	return yearToString[y]
}
