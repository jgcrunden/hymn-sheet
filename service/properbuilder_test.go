package service

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/jgcrunden/hymn-sheet/model"
)

const (
	validOrdo = `
{
	"2025-08-17": { "_": "c-ordinary-20" },
	"2024-12-29": { "_": "holy-family" },
	"2025-06-08": { "_": "pentecost", "suppressed": "c-ordinary-10" }
}
`
	exampleOrdoInvalidJSON = `
[
	{
		"2025-08-31": { "_": "c-ordinary-22" }
	}
]
`
	ordoMissingSuppressedSunday = `
{
	"2025-06-08": { "_": "pentecost" }
}
`
	validCalendar = `
{
    "c-ordinary-22": {
        "entranceAntiphon": "Have mercy on me, O Lord, for I cry to you all the day long. O Lord, you are good and forgiving, full of mercy to all who call to you.",
        "firstReading": "Sirach 3:17-18, 20, 28-29",
        "responsorialPsalm": "In your goodness, O God, you provided a home for the poor.",
        "secondReading": "Hebrews 12:18-19, 22-24a",
        "gospelAcclamation": "Alleluia, alleluia. Take my yoke upon you, says the Lord, and learn from me, for I am gentle and lowly in heart. Alleluia.",
        "gospel": "Luke 14:1, 7-14",
        "communionAntiphon": "How great is the goodness, Lord, that you keep for those who fear you."
    }
}
	`
	invalidCalendar = `
[
{
    "c-ordinary-22": {
        "entranceAntiphon": "Have mercy on me, O Lord, for I cry to you all the day long. O Lord, you are good and forgiving, full of mercy to all who call to you.",
        "firstReading": "Sirach 3:17-18, 20, 28-29",
        "responsorialPsalm": "In your goodness, O God, you provided a home for the poor.",
        "secondReading": "Hebrews 12:18-19, 22-24a",
        "gospelAcclamation": "Alleluia, alleluia. Take my yoke upon you, says the Lord, and learn from me, for I am gentle and lowly in heart.  Alleluia.",
        "gospel": "Luke 14:1, 7-14",
        "communionAntiphon": "How great is the goodness, Lord, that you keep for those who fear you."
    }
}
]
	`
	fileDoesNotExistError = "open NONSENSE_FILE.json: no such file or directory"
)

type mockFileReaderGetOrdoValid struct{}

// ReadFile implements FileReader.
func (m mockFileReaderGetOrdoValid) ReadFile(filename string) ([]byte, error) {
	return []byte(validOrdo), nil
}

type mockFileReaderGetOrdoInvalidJSON struct{}

// ReadFile implements FileReader.
func (m mockFileReaderGetOrdoInvalidJSON) ReadFile(filename string) ([]byte, error) {
	return []byte(exampleOrdoInvalidJSON), nil
}

func TestGetOrdoFailure(t *testing.T) {
	config := model.Config{}
	m := mockFileReaderFailure{}
	p := NewProperBuilder(m, config)
	err := p.GetOrdo("foo.json")

	if err.Error() != fileDoesNotExistError {
		t.Errorf("Expected %s err got %s", fileDoesNotExistError, err.Error())
	}
}

func TestGetOrdoInvalidJSON(t *testing.T) {
	config := model.Config{}
	m := mockFileReaderGetOrdoInvalidJSON{}
	p := NewProperBuilder(m, config)
	err := p.GetOrdo("foo.json")

	expectedError := "json: cannot unmarshal array into Go value of type model.ordo"
	if err.Error() != expectedError {
		t.Errorf("Expected %s err got %s", expectedError, err.Error())
	}
}

func TestGetOrdoDayDoesNotExist(t *testing.T) {
	d := model.Date{}
	d.Time, _ = time.Parse(model.DateLayout, "2025-08-31")
	config := model.Config{Date: d}
	m := mockFileReaderGetOrdoValid{}
	p := NewProperBuilder(m, config)
	err := p.GetOrdo("foo.json")

	expectedError := "date does not exist in calendar"
	if err.Error() != expectedError {
		t.Errorf("Expected no err got %s", err.Error())
	}
}

func TestGetOrdoSuccess(t *testing.T) {
	d := model.Date{}
	d.Time, _ = time.Parse(model.DateLayout, "2025-08-17")
	config := model.Config{Date: d}
	m := mockFileReaderGetOrdoValid{}
	p := NewProperBuilder(m, config)
	err := p.GetOrdo("foo.json")

	if err != nil {
		t.Errorf("Expected no err got %s", err.Error())
	}
}

func TestDeriveCyclesSuccess(t *testing.T) {
	d := model.Date{}
	d.Time, _ = time.Parse(model.DateLayout, "2025-08-17")
	config := model.Config{
		Date: d,
	}
	m := mockFileReaderGetOrdoValid{}
	p := NewProperBuilder(m, config)
	p.properDay = "c-ordinary-time-20"
	cycles, err := p.DeriveCycles()
	if err != nil {
		t.Errorf("Expected no error, got %s", err.Error())
	}

	expectedSundayYear := model.C
	if cycles.LectionaryYearSunday != expectedSundayYear {
		t.Errorf("Expected %v, got %v", expectedSundayYear, cycles.LectionaryYearSunday)
	}

	expectedWeekdayYear := 1
	if cycles.LectionaryYearWeekday != uint(expectedWeekdayYear) {
		t.Errorf("Expected %d, got %d", expectedWeekdayYear, cycles.LectionaryYearWeekday)
	}

	expectedPsalterWeek := 4
	if cycles.PsalterWeek != uint(expectedPsalterWeek) {
		t.Errorf("Expected %d, got %d", expectedPsalterWeek, cycles.PsalterWeek)
	}
}

func TestDeriveLectionaryYearA(t *testing.T) {
	d := model.Date{}
	d.Time, _ = time.Parse(model.DateLayout, "2026-08-31")
	config := model.Config{
		Date: d,
	}
	m := mockFileReaderGetOrdoValid{}
	p := NewProperBuilder(m, config)
	p.properDay = "c-ordinary-time-20"
	cycles, err := p.DeriveCycles()
	if err != nil {
		t.Errorf("Expected no error, got %s", err.Error())
	}

	expectedYear := model.A
	if cycles.LectionaryYearSunday != expectedYear {
		t.Errorf("Expected %v, got %v", expectedYear, cycles.LectionaryYearSunday)
	}
}

func TestDeriveLectionaryYearB(t *testing.T) {
	d := model.Date{}
	d.Time, _ = time.Parse(model.DateLayout, "2027-08-31")
	config := model.Config{
		Date: d,
	}
	m := mockFileReaderGetOrdoValid{}
	p := NewProperBuilder(m, config)
	p.properDay = "c-ordinary-time-20"
	cycles, err := p.DeriveCycles()
	if err != nil {
		t.Errorf("Expected no error, got %s", err.Error())
	}

	expectedYear := model.B
	if cycles.LectionaryYearSunday != expectedYear {
		t.Errorf("Expected %v, got %v", expectedYear, cycles.LectionaryYearSunday)
	}
}

func TestDerivePsalterWeekFromSeasonStart(t *testing.T) {
	d := model.Date{}
	d.Time, _ = time.Parse(model.DateLayout, "2024-12-29")
	config := model.Config{
		Date: d,
	}
	m := mockFileReaderGetOrdoValid{}
	p := NewProperBuilder(m, config)
	err := p.GetOrdo("foo.json")

	if err != nil {
		t.Errorf("Expected no err got %s", err.Error())
	}
	cycles, err := p.DeriveCycles()
	if err != nil {
		t.Errorf("Expected no err got %s", err.Error())
	}

	expectedPsalterWeek := 1
	if cycles.PsalterWeek != uint(expectedPsalterWeek) {
		t.Errorf("Expected psalter week %d, got %d", expectedPsalterWeek, cycles.PsalterWeek)

	}
}

type mockFileReaderGetOrdoMissingSuppressedSunday struct{}

// ReadFile implements FileReader.
func (m mockFileReaderGetOrdoMissingSuppressedSunday) ReadFile(filename string) ([]byte, error) {
	return []byte(ordoMissingSuppressedSunday), nil
}

func TestDerivePsalterWeekMissingSuppressedSunday(t *testing.T) {
	d := model.Date{}
	d.Time, _ = time.Parse(model.DateLayout, "2025-06-08")
	config := model.Config{
		Date: d,
	}
	m := mockFileReaderGetOrdoMissingSuppressedSunday{}
	p := NewProperBuilder(m, config)
	err := p.GetOrdo("foo.json")

	if err != nil {
		t.Errorf("Expected no err got %s", err.Error())
	}
	_, err = p.DeriveCycles()
	expectedError := "cannot derive psalter week from pentecost: no start of the season, does not have a numbered week and does not have suppressed sunday info"
	if err.Error() != expectedError {
		t.Errorf("Expected err %s, got %s", expectedError, err.Error())
	}
}

func TestDerivePsalterWeekWithSuppressedSunday(t *testing.T) {
	d := model.Date{}
	d.Time, _ = time.Parse(model.DateLayout, "2025-06-08")
	config := model.Config{
		Date: d,
	}
	m := mockFileReaderGetOrdoValid{}
	p := NewProperBuilder(m, config)
	err := p.GetOrdo("foo.json")
	if err != nil {
		t.Errorf("Expected no err got %s", err.Error())
	}
	cycles, err := p.DeriveCycles()
	if err != nil {
		t.Errorf("Expected no err got %s", err.Error())
	}

	expectedPsalterWeek := 2
	if cycles.PsalterWeek != uint(expectedPsalterWeek) {
		t.Errorf("Expected psalter week %d, got %d", expectedPsalterWeek, cycles.PsalterWeek)

	}
}

func TestGetPropersNoCalendar(t *testing.T) {
	config := model.Config{}
	m := mockFileReaderFailure{}
	p := NewProperBuilder(m, config)
	_, err := p.GetPropers("foo.json")
	if err.Error() != fileDoesNotExistError {
		t.Errorf("Expected %s got %s", fileDoesNotExistError, err.Error())
	}
}

type mockFileReaderGetPropersInvalidJSON struct{}

func (m mockFileReaderGetPropersInvalidJSON) ReadFile(filename string) ([]byte, error) {
	return []byte(invalidCalendar), nil

}
func TestGetPropersInvalidJSON(t *testing.T) {
	config := model.Config{}
	m := mockFileReaderGetPropersInvalidJSON{}
	p := NewProperBuilder(m, config)
	_, err := p.GetPropers("foo.json")
	expectedError := "json: cannot unmarshal array into Go value of type model.calendar"
	if err.Error() != expectedError {
		t.Errorf("Expected %s error, got %s", expectedError, err.Error())
	}
}

type mockFileReaderGetPropersValid struct{}

func (m mockFileReaderGetPropersValid) ReadFile(filename string) ([]byte, error) {
	return []byte(validCalendar), nil
}
func TestGetPropersDayDoesNotExist(t *testing.T) {
	config := model.Config{}
	m := mockFileReaderGetPropersValid{}
	p := NewProperBuilder(m, config)
	p.properDay = "c-ordinary-21"
	_, err := p.GetPropers("foo.json")
	expectedError := fmt.Sprintf("could not find propers for %s", p.properDay)
	if err.Error() != expectedError {
		t.Errorf("Expected %s error, got %s", expectedError, err.Error())
	}
}

func TestGetPropersSuccess(t *testing.T) {
	config := model.Config{}
	m := mockFileReaderGetPropersValid{}
	p := NewProperBuilder(m, config)
	p.properDay = "c-ordinary-22"
	propers, err := p.GetPropers("foo.json")
	if err != nil {
		t.Errorf("Expected no err, got %s", err.Error())
	}
	expectedOutput := `{"entranceAntiphon":"Have mercy on me, O Lord, for I cry to you all the day long. O Lord, you are good and forgiving, full of mercy to all who call to you.","firstReading":"Sirach 3:17-18, 20, 28-29","responsorialPsalm":"In your goodness, O God, you provided a home for the poor.","secondReading":"Hebrews 12:18-19, 22-24a","gospelAcclamation":"Alleluia, alleluia. Take my yoke upon you, says the Lord, and learn from me, for I am gentle and lowly in heart. Alleluia.","gospel":"Luke 14:1, 7-14","communionAntiphon":"How great is the goodness, Lord, that you keep for those who fear you."}`
	output, err := json.Marshal(propers)
	if err != nil {
		t.Errorf("Expected no error, got %s", err.Error())
	}

	if string(output) != expectedOutput {
		t.Errorf("Expected %s, got %s", expectedOutput, string(output))
	}
}
