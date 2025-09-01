package service

import (
	"errors"
	"testing"
)

const exampleConfigValid = `{
    "date": "2025-08-31",
    "yearAndPsalterRef": true,
    "euchPrayerOption": true,
    "entranceAntiphon": true,
    "communionAntiphon": true,
    "hymns": [
        {
            "position": "PROCESSIONAL",
            "friendlyName": "Opening Hymn",
            "hymnBook": "LAU",
            "hymnNum": 123,
            "verses": 4,
            "columns": true
        },
        {
            "position": "OFFERTORY",
            "friendlyName": "Presentation of gifts",
            "hymnBook": "LAU",
            "hymnNum": 234,
            "verses": 4,
            "columns": true
        },
        {
            "position": "COMMUNION",
            "friendlyName": "Communion Hymn",
            "hymnBook": "LAU",
            "hymnNum": 456,
            "verses": 4,
            "columns": true
        },
        {
            "position": "RECESSIONAL",
            "friendlyName": "Concluding Hymn",
            "hymnBook": "LAU",
            "hymnNum": 567,
            "verses": 4,
            "columns": true
        }
    ]
}`

const exampleConfigInvalid = `
{
    "date": "31st August 2025",
    "yearAndPsalterRef": true,
    "euchPrayerOption": true,
    "entranceAntiphon": true,
    "communionAntiphon": true,
    "hymns": [
        {
            "position": "PROCESSIONAL",
            "friendlyName": "Opening Hymn",
            "hymnBook": "LAU",
            "hymnNum": 123,
            "verses": 4,
            "columns": true
        }
    ]
}
`

type mockFileReaderFailure struct{}

// ReadFile implements FileReader.
func (m mockFileReaderFailure) ReadFile(filename string) ([]byte, error) {
	return nil, errors.New("open NONSENSE_FILE.json: no such file or directory")
}

type mockFileReaderGetConfigSuccess struct{}

// ReadFile implements FileReader.
func (m mockFileReaderGetConfigSuccess) ReadFile(filename string) ([]byte, error) {
	return []byte(exampleConfigValid), nil
}

type MockFileReaderGetConfigInvalidJSON struct{}

// ReadFile implements FileReader.
func (m MockFileReaderGetConfigInvalidJSON) ReadFile(filename string) ([]byte, error) {
	return []byte(exampleConfigInvalid), nil
}

func TestReadConfigFile_FileDoesntExist(t *testing.T) {
	m := mockFileReaderFailure{}
	parser := NewConfigParser(m)
	err := parser.ReadConfigFile("NONSENSE_FILE.json")
	expectedError := "open NONSENSE_FILE.json: no such file or directory"
	if err.Error() != expectedError {
		t.Errorf("Expected %s, got %s", expectedError, err)
	}
}

func TestReadConfigFileSuccess(t *testing.T) {
	m := mockFileReaderGetConfigSuccess{}
	parser := NewConfigParser(m)
	err := parser.ReadConfigFile("foo.json")
	if err != nil {
		t.Errorf("Expected no err, got %s", err)
	}
}

func TestReadConfigFileInvalidJson(t *testing.T) {
	m := MockFileReaderGetConfigInvalidJSON{}
	parser := NewConfigParser(m)
	err := parser.ReadConfigFile("foo.json")
	expectedError := `parsing time "31st August 2025" as "2006-01-02": cannot parse "31st August 2025" as "2006"`
	if err.Error() != expectedError {
		t.Errorf("Expected %s, got %s", expectedError, err)
	}
}
