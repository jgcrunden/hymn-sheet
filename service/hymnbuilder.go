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

func (h HymnBuilder) GetHymns() ([]model.Hymn, error) {
	for i, v := range h.config.Hymns {
		filename := fmt.Sprintf("./resources/%s/%d.txt", v.HymnBook, v.HymnNum)
		content, err := h.reader.ReadFile(filename)
		if err != nil {
			return nil, err
		}
		reduced := reduceVerses(string(content), v.Verses)
		h.config.Hymns[i].Lyrics = reduced
		// Add latex markup to each verse/chorus
	}
	return h.config.Hymns, nil
}

func (h HymnBuilder) TagHymnVerses(input []model.Hymn) ([]model.Hymn, error) {
	for i, v := range input {
		versesInput := strings.Split(v.Lyrics, "\n\n")
		versesOutput := make([]string, 0)
		for _, w := range versesInput {
			isChorus := strings.Contains(w, "[")
			hymnsTmp := w
			if isChorus {
				hymnsTmp = strings.ReplaceAll(hymnsTmp, "[", "")
				hymnsTmp = strings.ReplaceAll(hymnsTmp, "]", "")
			}
			fmt.Println(hymnsTmp)
			hymnsTmp = strings.ReplaceAll(hymnsTmp, "\n", "\\\\*\n")
			hymnsTmp = fmt.Sprintf("%s\\\\*", hymnsTmp)
			if isChorus {
				hymnsTmp = fmt.Sprintf("\\textit{%s}", hymnsTmp)
			} else {
				hymnsTmp = fmt.Sprintf("\\flagverse{\\printcount.} %s", hymnsTmp)
			}
			hymnsTmp = fmt.Sprintf("\\begin{verse}\n%s\n\\end{verse}", hymnsTmp)
			versesOutput = append(versesOutput, hymnsTmp)
		}
		input[i].Lyrics = fmt.Sprintf("\\setcounter{count}{0}\n%s", strings.Join(versesOutput, "\n\n"))
	}
	return input, nil
}
