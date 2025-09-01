package service

import (
	"fmt"
	"strings"

	"github.com/jgcrunden/hymn-sheet/model"
)

type HymnBuilder struct {
	reader FileReader
	config model.Config
}

func NewHymnBuilder(reader FileReader, config model.Config) HymnBuilder {
	return HymnBuilder{
		reader: reader,
		config: config,
	}
}

func reduceVerses(input string, versesLimit uint) string {
	versesTmp := strings.Split(string(input), "\n\n")
	verseCount := 0
	verses := make([]string, 0)
	for _, j := range versesTmp {
		if j[0:1] != "[" {
			verseCount += 1
		}
		if verseCount <= int(versesLimit) {
			verses = append(verses, j)
		} else {
			break
		}
	}
	return strings.Join(verses, "\n\n")

}

func (h HymnBuilder) GetHymns() ([]string, error) {
	res := make([]string, 0)
	for _, v := range h.config.Hymns {
		filename := fmt.Sprintf("./resources/%s/%d.txt", v.HymnBook, v.HymnNum)
		content, err := h.reader.ReadFile(filename)
		if err != nil {
			return nil, err
		}
		reduced := reduceVerses(string(content), v.Verses)

		res = append(res, reduced)

	}
	return res, nil
}
