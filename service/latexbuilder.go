package service

import (
	"bytes"
	"fmt"
	"os"

	"github.com/jgcrunden/hymn-sheet/model"
	"github.com/jgcrunden/hymn-sheet/utils"
)

const preamble = `\documentclass[12pt]{article}
\usepackage[parfill]{parskip} % for supressing first line indent
\usepackage[margin={1.5cm, 1cm}]{geometry}
\usepackage{multicol} % creating columns
\usepackage{verse} % rendering verses
\usepackage[none]{hyphenat} % prevent splitting words over two lines
\usepackage{helvet} % font
\usepackage{fontspec}
\usepackage[super]{nth}
\setmainfont{Arial}
\pagestyle{empty} % suppress page numbers
\setlength{\vleftskip}{12pt} % decrease gap between verse numbers and verses
\newcounter{count}
\newcommand\printcount{\addtocounter{count}{1}\thecount}

`

func wrapHymn(friendlyName string, lyrics string) string {
	return fmt.Sprintf("\\begin{center}\n\\subsection*{%s}\n\\end{center}\n\\begin{multicols}{2}\n\\setcounter{count}{0}\n%s\n\\end{multicols}\n\n", friendlyName, lyrics)
}

func GenerateLatex(config model.Config, propers model.Propers, cycles model.Cycles) (string, error) {

	var byteBuffer bytes.Buffer
	// preamble
	if _, err := byteBuffer.WriteString(preamble); err != nil {
		return "", err
	}

	prettifiedProperDay, err := utils.PrettifyProperDay(propers.ProperDay)
	if err != nil {
		return "", err
	}

	prettifiedDate := fmt.Sprintf("\\nth{%d} %s", config.Date.Day(), config.Date.Month().String())

	// Title and info about day/week/year
	if _, err := byteBuffer.WriteString(fmt.Sprintf("\\begin{document}\n\\begin{center}\n\\section*{%s\\\\%s}\n\n\\begin{tabular}{ |l|l| }\n\\hline\n\\textbf{Readings} & Sunday Mass Cycle %s : Weekday Mass Year %d : Psalter %d\\\\\n\\textbf{Eucharistic Prayer} & %d (or at priest’s choice)\\\\\n\\hline\n\\end{tabular}\n\\end{center}\n\n", prettifiedProperDay, prettifiedDate, cycles.LectionaryYearSunday, cycles.LectionaryYearWeekday, cycles.PsalterWeek, config.EuchPrayerOption)); err != nil {
		return "", err
	}

	// opening hymn
	var openingHymn *model.Hymn
	for _, v := range config.Hymns {
		if v.Position == model.PROCESSIONAL {
			byteBuffer.WriteString(wrapHymn(v.FriendlyName, v.Lyrics))
		}
	}
	if openingHymn != nil {
	}

	// table with liturgy of the word details
	byteBuffer.WriteString(fmt.Sprintf("\\begin{center}\n\\subsection*{Liturgy of the Word}\n\\begin{tabular}{ |p{0.25\\textwidth}|p{0.5\\textwidth}| }\n\\hline\n\\textbf{Entrance Antiphon} & %s\\\\\n\\hline\n\\textbf{First Reading} & %s\\\\\n\\hline\n\\textbf{Psalm Response} & %s\\\\\n\\hline\n\\textbf{Second Reading} & %s\\\\\n\\hline\n\\textbf{Gospel Acclamation} & %s\\\\\n\\hline\n\\textbf{Gospel} & %s\\\\\n\\hline\n\\textbf{Communion\\newline Antiphon} & %s\\\\\n\\hline\n\\end{tabular}\\end{center}", propers.EntranceAntiphon, propers.FirstReading, propers.ResponsorialPsalm, propers.SecondReading, propers.GospelAcclamation, propers.Gospel, propers.CommunionAntiphon))

	// offertory hymn
	for _, v := range config.Hymns {
		if v.Position == model.OFFERTORY {
			byteBuffer.WriteString(wrapHymn(v.FriendlyName, v.Lyrics))
		}
	}

	byteBuffer.WriteString("\\newpage\n\n")

	// communion hymn(s)
	for _, v := range config.Hymns {
		if v.Position == model.COMMUNION {
			byteBuffer.WriteString(wrapHymn(v.FriendlyName, v.Lyrics))
		}
	}

	// recessional hymn
	for _, v := range config.Hymns {
		if v.Position == model.RECESSIONAL {
			byteBuffer.WriteString(wrapHymn(v.FriendlyName, v.Lyrics))
		}
	}

	byteBuffer.WriteString("\\end{document}")
	latexFileName := fmt.Sprintf("%s.tex", config.Date.String())
	os.WriteFile(latexFileName, byteBuffer.Bytes(), 0644)
	return latexFileName, nil
}
