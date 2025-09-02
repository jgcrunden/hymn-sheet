package model

import (
	"bytes"
	"encoding/json"
	"time"
)

type Config struct {
	Date              Date   `json:"date"`
	YearAndPsalterRef bool   `json:"yearAndPsalterRef"`
	EuchPrayerOption  uint   `json:"euchPrayerOption"`
	EntranceAntiphon  bool   `json:"entranceAntiphon"`
	CommunionAntiphon bool   `json:"communionAntiphon"`
	Hymns             []Hymn `json:"hymns"`
}

type Date struct {
	time.Time
}

func (d Date) String() string {
	return d.Time.Format(DateLayout)
}

const DateLayout = "2006-01-02"

func (d *Date) UnmarshalJSON(b []byte) error {
	var str string
	err := json.Unmarshal(b, &str)
	if err != nil {
		return err
	}

	d.Time, err = time.Parse(DateLayout, str)
	return err
}

func (d Date) MarshalJSON() ([]byte, error) {
	if d.Time.IsZero() {
		return []byte("null"), nil
	}
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(d.String())
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

type Hymn struct {
	Position     Position `json:"position"`
	FriendlyName string   `json:"friendlyName"`
	HymnBook     HymnBook `json:"hymnBook"`
	HymnNum      uint     `json:"hymnNum"`
	Verses       uint     `json:"verses"`
	Columns      bool     `json:"columns"`
	Lyrics       string
}

type Position int

// Goal type
const (
	PROCESSIONAL Position = iota + 1
	OFFERTORY
	COMMUNION
	RECESSIONAL
)

func (p Position) String() string {
	return posToString[p]
}

var posToString = map[Position]string{
	PROCESSIONAL: "PROCESSIONAL",
	OFFERTORY:    "OFFERTORY",
	COMMUNION:    "COMMUNION",
	RECESSIONAL:  "RECESSIONAL",
}

var posToID = map[string]Position{
	"PROCESSIONAL": PROCESSIONAL,
	"OFFERTORY":    OFFERTORY,
	"COMMUNION":    COMMUNION,
	"RECESSIONAL":  RECESSIONAL,
}

// MarshalJSON marshals the enum as a quoted json string
func (p Position) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(posToString[p])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (p *Position) UnmarshalJSON(b []byte) error {
	var str string
	err := json.Unmarshal(b, &str)
	if err != nil {
		return err
	}

	*p = posToID[str]
	return nil
}

type HymnBook int

const (
	LAU HymnBook = iota + 1
	LHOAN
)

var hymnToString = map[HymnBook]string{
	LAU:   "LAU",
	LHOAN: "LHOAN",
}

var hymnToID = map[string]HymnBook{
	"LAU":   LAU,
	"LHOAN": LHOAN,
}

func (h HymnBook) String() string {
	return hymnToString[h]
}

// MarshalJSON marshals the enum as a quoted json string
func (h HymnBook) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(hymnToString[h])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (h *HymnBook) UnmarshalJSON(b []byte) error {
	var str string
	err := json.Unmarshal(b, &str)
	if err != nil {
		return err
	}

	*h = hymnToID[str]
	return nil
}
