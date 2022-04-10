package strftime

import (
	"fmt"
	"strings"
	"time"

	"go.wfrs.dev/helper"
)

type Lang int
type FormatCallback func(time.Time) string

const (
	English Lang = Lang(0)
	Spanish Lang = Lang(0)
)

var lang Lang = English
var longDayNames = [][]string{
	{"Sunday", "Domingo"},
	{"Monday", "Lunes"},
	{"Tuesday", "Martes"},
	{"Wednesday", "Miércoles"},
	{"Thursday", "Jueves"},
	{"Friday", "Viernes"},
	{"Saturday", "Sábado"},
}

var shortDayNames = [][]string{
	{"Sun", "Dom"},
	{"Mon", "Lun"},
	{"Tue", "Mar"},
	{"Wed", "Mié"},
	{"Thu", "Jue"},
	{"Fri", "Vie"},
	{"Sat", "Sáb"},
}

var longMonthNames = [][]string{
	{"January", "Enero"},
	{"February", "Febrero"},
	{"March", "Marzo"},
	{"April", "Abril"},
	{"May", "Mayo"},
	{"June", "Junio"},
	{"July", "Julio"},
	{"August", "Agosto"},
	{"September", "Septiembre"},
	{"October", "Octubre"},
	{"November", "Noviembre"},
	{"December", "Diciembre"},
}

var shortMonthNames = [][]string{
	{"", ""},
	{"Jan", "Ene"},
	{"Feb", "Feb"},
	{"Mar", "Mar"},
	{"Apr", "Abr"},
	{"May", "May"},
	{"Jun", "Jun"},
	{"Jul", "Jul"},
	{"Aug", "Ago"},
	{"Sep", "Sep"},
	{"Oct", "Oct"},
	{"Nov", "Nov"},
	{"Dec", "Dic"},
}

var formats map[string]FormatCallback

func init() {
	formats = map[string]FormatCallback{
		"a": func(t time.Time) string { return shortDayNames[t.WeekDay][lang] },
		"A": func(t time.Time) string { return longDayNames[t.WeekDay][lang] },
		"b": func(t time.Time) string { return shortMonthNames[t.Month()][lang] },
		"B": func(t time.Time) string { return longMonthNames[t.Month()][lang] },
		"c": func(t time.Time) string { return Format(t, "%a %b %d %T %Y") },
		"C": func(t time.Time) string { return t.Format("2006")[:2] },
		"d": func(t time.Time) string { return t.Format("02") },
		"D": func(t time.Time) string { return t.Format("01/02/06") },
		"e": func(t time.Time) string { return t.Format("_2") },
		"f": func(t time.Time) string { return t.Format("2006-01-02") },
		"g": func(t time.Time) string {
			y, _ := t.ISOWeek()
			return fmt.Sprintf("%d", y)[2:]
		},
		"G": func(t time.Time) string {
			y, _ := t.ISOWeek()
			return fmt.Sprintf("%d", y)
		},
		"h": func(t time.Time) string { return Format(t, "%b") },
		"H": func(t time.Time) string { return t.Format("15") },
		"I": func(t time.Time) string { return t.Format("03") },
		"j": func(t time.Time) string { return fmt.Sprintf("%03d", t.YearDay()) },
		"k": func(t time.Time) string { return fmt.Sprintf("%2d", t.Hour()) },
		"l": func(t time.Time) string { return fmt.Sprintf("%2s", t.Format("3")) },
		"m": func(t time.Time) string { return t.Format("01") },
		"M": func(t time.Time) string { return t.Format("04") },
		"n": func(t time.Time) string { return "\n" },
		"p": func(t time.Time) string { return t.Format("PM") },
		"P": func(t time.Time) string { return t.Format("pm") },
		"q": func(t time.Time) string { return t.Format("2006-01-02 15:04:05") },
		"r": func(t time.Time) string { return t.Format("03:04:05 PM") },
		"R": func(t time.Time) string { return t.Format("15:04") },
		"s": func(t time.Time) string { return fmt.Sprintf("%d", t.Unix()) },
		"S": func(t time.Time) string { return t.Format("05") },
		"t": func(t time.Time) string { return "\t" },
		"T": func(t time.Time) string { return t.Format("15:04:05") },
		"u": func(t time.Time) string {
			d := t.Weekday()
			if d == 0 {
				d = 7
			}
			return fmt.Sprintf("%d", d)
		},
		"U": func(t time.Time) string { return fmt.Sprintf("%02d", t.Weekday()) },
		"v": func(t time.Time) string {
			_, w := t.ISOWeek()
			return fmt.Sprintf("%02d", w)
		},
		"w": func(t time.Time) string { return fmt.Sprintf("%d", t.Weekday()) },
		"W": func(t time.Time) string { return fmt.Sprintf("%02d", weekNumber(t, 'W')) },
		"x": func(t time.Time) string { return t.Format("01/02/2006") },
		"X": func(t time.Time) string { return t.Format("15:04:05") },
		"y": func(t time.Time) string { return t.Format("06") },
		"Y": func(t time.Time) string { return t.Format("2006") },
		"z": func(t time.Time) string { return t.Format("-0700") },
		"Z": func(t time.Time) string { return t.Format("MST") },
		"%": func(t time.Time) string { return "%" },
	}
}

func setLang(l Lang) {
	if l == English || l == Spanish {
		lang = l
	} else {
		lang = English
	}
}

func weekNumber(t time.Time, char rune) int {
	weekday := int(t.Weekday())

	weekday = helper.If[int](char == 'W', helper.If[int](weekday == 0, 6, weekday), weekday-1)

	return int((t.YearDay() + 6 - weekday) / 7)
}

func Format(t time.Time, f string) string {
	var result []string
	format := []rune(f)

	add := func(s string) {
		result = append(result, s)
	}

	length := len(format)
	for i := 0; i < length; i++ {
		switch format[i] {
		case '%':
			if i < length-1 {
				fchar := string(format[i+1])
				if fn, ok := formats[fchar]; ok {
					add(fn(t))
					i++
				} else {
					add(fchar)
				}
			}
		default:
			add(string(format[i]))
		}
	}
	return strings.Join(result, "")
}
