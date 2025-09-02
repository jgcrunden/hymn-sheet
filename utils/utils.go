package utils

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

var (
	numToOrdinal = map[int]string{
		1:  "First",
		2:  "Second",
		3:  "Third",
		4:  "Fourth",
		5:  "Fifth",
		6:  "Sixth",
		7:  "Seventh",
		8:  "Eighth",
		9:  "Ninth",
		10: "Tenth",
		11: "Eleventh",
		12: "Twelfth",
		13: "Thirteenth",
		14: "Fourteenth",
		15: "Fifteenth",
		16: "Sixteenth",
		17: "Seventeenth",
		18: "Eighteenth",
		19: "Nineteenth",
		20: "Twentieth",
		21: "Twenty First",
		22: "Twenty Second",
		23: "Twenty Third",
		24: "Twenty Fourth",
		25: "Twenty Fifth",
		26: "Twenty Sixth",
		27: "Twenty Seventh",
		28: "Twenty Eighth",
		29: "Twenty Ninth",
		30: "Thirtieth",
		31: "Thirty First",
	}
	seasons = map[string]string{
		"ordinary":  "%s Sunday in Ordinary Time",
		"advent":    "%s Sunday of Advent",
		"christmas": "%s Sunday of Christmas",
		"lent":      "%s Sunday of Lent",
		"easter":    "%s Sunday of Easter",
	}
	functionWords = []string{"to", "the", "of", "and"}
)

func PrettifyProperDay(input string) (string, error) {
	var template *string
	for k, v := range seasons {
		if strings.Contains(input, k) {
			template = &v
			break
		}
	}

	var output string
	if template != nil {
		splitStr := strings.Split(input, "-")
		num, err := strconv.Atoi(splitStr[len(splitStr)-1])
		if err != nil {
			return "", err
		}
		output = fmt.Sprintf(*template, numToOrdinal[num])
	} else {
		splitStr := strings.Split(input, "-")
		if splitStr[0] == "a" || splitStr[0] == "b" || splitStr[0] == "c" {
			splitStr = splitStr[1:]
		}
		for i := range splitStr {
			if i == 0 || !slices.Contains(functionWords, splitStr[i]) {
				splitStr[i] = strings.ToUpper(splitStr[i][:1]) + splitStr[i][1:]
			}
		}
		output = strings.Join(splitStr, " ")
	}
	return output, nil
}
