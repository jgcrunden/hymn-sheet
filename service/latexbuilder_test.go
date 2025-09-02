package service

import (
	"testing"
	"time"

	"github.com/jgcrunden/hymn-sheet/model"
)

func TestBuildLatex(t *testing.T) {

	d := model.Date{}
	d.Time, _ = time.Parse(model.DateLayout, "2025-08-31")
	config := model.Config{
		Date:              d,
		YearAndPsalterRef: true,
		EuchPrayerOption:  2,
		EntranceAntiphon:  true,
		CommunionAntiphon: true,
		Hymns: []model.Hymn{
			{
				Position:     model.PROCESSIONAL,
				FriendlyName: "Opening Hymn",
				HymnBook:     model.LAU,
				HymnNum:      123,
				Verses:       3,
				Columns:      true,
				Lyrics: `\begin{verse}
\flagverse{\printcount.} Diverse in culture, nation, race,\\*
we come together by your grace.\\*
God let us be a meeting ground\\*
where hope and healing love are found.\\*
\end{verse}

\begin{verse}
\textit{This is a chorus\\*
This is a chorus\\*
This is a chorus\\*
This is a chorus\\*}
\end{verse}

\begin{verse}
\flagverse{\printcount.} God, let us be a bridge of care\\*
connecting people everywhere.\\*
Help us confront all fear and hate\\*
and lust for power that separate.\\*
\end{verse}

\begin{verse}
\flagverse{\printcount.} When chasms widen, storms arise,\\*
O Holy Spirit, make us wise.\\*
Let us resolve like steel be strong\\*
\end{verse}`,
			},
			{
				Position:     model.OFFERTORY,
				FriendlyName: "Preparation of gifts",
				HymnBook:     model.LAU,
				HymnNum:      123,
				Verses:       3,
				Columns:      true,
				Lyrics: `\begin{verse}
\flagverse{\printcount.} Diverse in culture, nation, race,\\*
we come together by your grace.\\*
God let us be a meeting ground\\*
where hope and healing love are found.\\*
\end{verse}

\begin{verse}
\textit{This is a chorus\\*
This is a chorus\\*
This is a chorus\\*
This is a chorus\\*}
\end{verse}

\begin{verse}
\flagverse{\printcount.} God, let us be a bridge of care\\*
connecting people everywhere.\\*
Help us confront all fear and hate\\*
and lust for power that separate.\\*
\end{verse}

\begin{verse}
\flagverse{\printcount.} When chasms widen, storms arise,\\*
O Holy Spirit, make us wise.\\*
Let us resolve like steel be strong\\*
\end{verse}`,
			},
			{
				Position:     model.COMMUNION,
				FriendlyName: "Holy Communion",
				HymnBook:     model.LAU,
				HymnNum:      123,
				Verses:       3,
				Columns:      true,
				Lyrics: `\begin{verse}
\flagverse{\printcount.} Diverse in culture, nation, race,\\*
we come together by your grace.\\*
God let us be a meeting ground\\*
where hope and healing love are found.\\*
\end{verse}

\begin{verse}
\textit{This is a chorus\\*
This is a chorus\\*
This is a chorus\\*
This is a chorus\\*}
\end{verse}

\begin{verse}
\flagverse{\printcount.} God, let us be a bridge of care\\*
connecting people everywhere.\\*
Help us confront all fear and hate\\*
and lust for power that separate.\\*
\end{verse}

\begin{verse}
\flagverse{\printcount.} When chasms widen, storms arise,\\*
O Holy Spirit, make us wise.\\*
Let us resolve like steel be strong\\*
\end{verse}`,
			},
			{
				Position:     model.RECESSIONAL,
				FriendlyName: "Concluding Hymn",
				HymnBook:     model.LAU,
				HymnNum:      123,
				Verses:       3,
				Columns:      true,
				Lyrics: `\begin{verse}
\flagverse{\printcount.} Diverse in culture, nation, race,\\*
we come together by your grace.\\*
God let us be a meeting ground\\*
where hope and healing love are found.\\*
\end{verse}

\begin{verse}
\textit{This is a chorus\\*
This is a chorus\\*
This is a chorus\\*
This is a chorus\\*}
\end{verse}

\begin{verse}
\flagverse{\printcount.} God, let us be a bridge of care\\*
connecting people everywhere.\\*
Help us confront all fear and hate\\*
and lust for power that separate.\\*
\end{verse}

\begin{verse}
\flagverse{\printcount.} When chasms widen, storms arise,\\*
O Holy Spirit, make us wise.\\*
Let us resolve like steel be strong\\*
\end{verse}`,
			},
		},
	}

	propers := model.Propers{
		ProperDay:         "22nd Sunday of Ordinary Time",
		EntranceAntiphon:  "This is the entrance antiphon",
		FirstReading:      "This is the first reading",
		ResponsorialPsalm: "This is the responsorial psalm",
		SecondReading:     "This is the second reading",
		GospelAcclamation: "This is the gospel acclamation",
		Gospel:            "This is the gospel",
		CommunionAntiphon: "This is the communion antiphon",
	}
	cycles := model.Cycles{
		LectionaryYearSunday:  model.C,
		LectionaryYearWeekday: 1,
		PsalterWeek:           2,
	}

	GenerateLatex(config, propers, cycles)
}
