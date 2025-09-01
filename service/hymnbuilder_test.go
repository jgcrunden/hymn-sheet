package service

import (
	"testing"

	"github.com/jgcrunden/hymn-sheet/model"
)

func TestFetchHymnsHymDoesNotExist(t *testing.T) {
	m := mockFileReaderFailure{}
	config := model.Config{
		Hymns: []model.Hymn{
			{
				Position:     model.PROCESSIONAL,
				FriendlyName: "Opening Hymn",
				HymnBook:     model.LAU,
				HymnNum:      123,
				Verses:       4,
				Columns:      true,
			},
		},
	}
	hymnBuilder := NewHymnBuilder(m, config)

	_, err := hymnBuilder.GetHymns()
	expectedError := "open NONSENSE_FILE.json: no such file or directory"
	if err == nil || err.Error() != expectedError {
		t.Errorf("Expected %v, got %v", expectedError, err)
	}
}

const hymn = `Diverse in culture, nation, race,
we come together by your grace.
God let us be a meeting ground
where hope and healing love are found. 

[This is a chorus
This is a chorus
This is a chorus
This is a chorus]

God, let us be a bridge of care
connecting people everywhere. 
Help us confront all fear and hate
and lust for power that separate.

When chasms widen, storms arise,
O Holy Spirit, make us wise.
Let us resolve like steel be strong

God, let us be a table spread
with gifts of love and broken bread,
where all find welcome, grace attends
and enemies arise as friends`

type mockFileReaderReturnHymn struct{}

// ReadFile implements FileReader.
func (m mockFileReaderReturnHymn) ReadFile(filename string) ([]byte, error) {
	return []byte(hymn), nil
}

func TestReduceHymn(t *testing.T) {
	m := mockFileReaderReturnHymn{}
	config := model.Config{
		Hymns: []model.Hymn{
			{
				Position:     model.PROCESSIONAL,
				FriendlyName: "Opening Hymn",
				HymnBook:     model.LAU,
				HymnNum:      123,
				Verses:       3,
				Columns:      true,
			},
		},
	}
	expectedResult := `Diverse in culture, nation, race,
we come together by your grace.
God let us be a meeting ground
where hope and healing love are found. 

[This is a chorus
This is a chorus
This is a chorus
This is a chorus]

God, let us be a bridge of care
connecting people everywhere. 
Help us confront all fear and hate
and lust for power that separate.

When chasms widen, storms arise,
O Holy Spirit, make us wise.
Let us resolve like steel be strong`
	hymnBuilder := NewHymnBuilder(m, config)
	res, err := hymnBuilder.GetHymns()
	if err != nil {
		t.Errorf("Expected no error, got %s", err.Error())
	}
	if res[0] != expectedResult {
		t.Errorf("Expected %s, got %s", expectedResult, res[0])
	}
}
