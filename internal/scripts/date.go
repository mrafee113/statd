package scripts

import (
	"bytes"
	"fmt"
	"html/template"
	"statd/config"
	"statd/pkg/utils"
	"strconv"
	"strings"
	"time"

	"github.com/mshafiee/jalali"
)

type ExtendedJalali struct {
	j jalali.JalaliTime
}

func (this ExtendedJalali) Format(layout string) string {
	var lowerP string
	if this.j.Hour() < 12 {
		lowerP = "AM"
	} else {
		lowerP = "PM"
	}
	upperB := this.j.Month().String()
	lowerB := upperB[:3]
	upperA := this.j.Weekday().String()
	lowerA := upperA[:3]
	upperIint := this.j.Hour() % 12
	if upperIint == 0 {
		upperIint = 12
	}
	upperI := strconv.Itoa(upperIint)
	replacer := strings.NewReplacer("%A", upperA, "%a", lowerA, "%I",
		upperI, "%B", upperB, "%b", lowerB, "%p", lowerP)
	f := replacer.Replace(layout)
	return this.j.Format(f)
}

func Date() (string, error) {
	if config.Cfg == nil {
		return "", fmt.Errorf("UninitializedConfigError")
	}

	cfg := config.Cfg
	gdate := time.Now()
	jdate := ExtendedJalali{jalali.JalaliFromTime(gdate)}
	dayColor := getDayColor(gdate.Format("Monday"))
	monthColor := getMonthColor(gdate.Format("January"))
	seasonIcon, seasonColor := getSeasonTheme(int(gdate.Month()) % 12)

	strftimeColorMap := map[rune]string{
		'a': ".dayColor",
		'A': ".dayColor",
		'd': ".dayColor",
		'b': ".monthColor",
		'B': ".monthColor",
		'm': ".monthColor",
		'y': ".Colors.Info",
		'Y': ".Colors.Info",
		'H': ".Colors.Info",
		'I': ".Colors.Info",
		'p': ".Colors.Info",
		'M': ".Colors.Info",
		'S': ".Colors.Info",
	}
	strftimeGMap := map[rune]string{
		'a': "Mon",
		'A': "Monday",
		'd': "02",
		'b': "Jan",
		'B': "January",
		'm': "01",
		'y': "06",
		'Y': "2006",
		'H': "15",
		'I': "03",
		'p': "PM",
		'M': "04",
		'S': "05",
	}

	tmplPattern := `{{.t.Format "%s" | Colorize .colorize %s}}`
	var oldnew []string
	for char, gFormat := range strftimeGMap {
		oldnew = append(oldnew, fmt.Sprintf("%%%c", char),
			fmt.Sprintf(tmplPattern, gFormat, strftimeColorMap[char]))
	}
	gReplacer := strings.NewReplacer(oldnew...)

	oldnew = []string{}
	for char, color := range strftimeColorMap {
		oldnew = append(oldnew, fmt.Sprintf("%%%c", char),
			fmt.Sprintf(tmplPattern, fmt.Sprintf("%%%c", char), color))
	}
	jReplacer := strings.NewReplacer(oldnew...)

	tmplData := map[string]any{
		"t":          gdate,
		"Icons":      cfg.Date.Icons,
		"Colors":     cfg.Colors,
		"dayColor":   dayColor,
		"monthColor": monthColor,
		"seasonIcon": utils.Colorize(cfg.Date.Colorize, seasonColor, seasonIcon),
		"dayOfWeek":  utils.Colorize(cfg.Date.Colorize, dayColor, gdate.Format("Mon")),
		"Colorize":   utils.Colorize,
		"colorize":   cfg.Date.Colorize,
	}

	var err error
	tmplData["gdate"], err = ExecTemplate("gdate", gReplacer.Replace(cfg.Date.GDate), tmplData)
	if err != nil {
		return "", err
	}

	tmplData["time"], err = ExecTemplate("time", gReplacer.Replace(cfg.Date.Time), tmplData)
	if err != nil {
		return "", err
	}

	tmplData["t"] = jdate
	tmplData["jdate"], err = ExecTemplate("jdate", jReplacer.Replace(cfg.Date.JDate), tmplData)
	if err != nil {
		return "", err
	}

	msg, err := ExecTemplate("format", gReplacer.Replace(cfg.Date.Format), tmplData)
	if err != nil {
		return "", nil
	}
	return msg, nil
}

func ExecTemplate(name, pattern string, data map[string]any) (string, error) {
	var buf bytes.Buffer
	tmpl := template.Must(template.New(name).Funcs(template.FuncMap{"Colorize": utils.Colorize}).Parse(pattern))
	if err := tmpl.Execute(&buf, data); err != nil {
		fmt.Println(buf.String())
		return buf.String(), fmt.Errorf("TemplateExecuteError: %w", err)
	}
	return buf.String(), nil
}

func getSeasonTheme(season int) (seasonIcon, seasonColor string) {
	switch {
	case 0 <= season && season <= 2:
		seasonIcon = config.Cfg.Date.Icons.Winter
		seasonColor = config.Cfg.Date.Colors.Winter
	case 3 <= season && season <= 5:
		seasonIcon = config.Cfg.Date.Icons.Spring
		seasonColor = config.Cfg.Date.Colors.Spring
	case 6 <= season && season <= 8:
		seasonIcon = config.Cfg.Date.Icons.Summer
		seasonColor = config.Cfg.Date.Colors.Summer
	case 9 <= season && season <= 11:
		seasonIcon = config.Cfg.Date.Icons.Autumn
		seasonColor = config.Cfg.Date.Colors.Autumn
	}
	return seasonIcon, seasonColor
}

func getDayColor(day string) string {
	colors := config.Cfg.Date.Colors
	switch day {
	case "Saturday":
		return colors.Saturday
	case "Sunday":
		return colors.Sunday
	case "Monday":
		return colors.Monday
	case "Tuesday":
		return colors.Tuesday
	case "Wednesday":
		return colors.Wednesday
	case "Thursday":
		return colors.Thursday
	case "Friday":
		return colors.Friday
	}
	return config.Cfg.Colors.Info
}

func getMonthColor(month string) string {
	colors := config.Cfg.Date.Colors
	switch month {
	case "January":
		return colors.January
	case "February":
		return colors.February
	case "March":
		return colors.March
	case "April":
		return colors.April
	case "May":
		return colors.May
	case "June":
		return colors.June
	case "July":
		return colors.July
	case "August":
		return colors.August
	case "September":
		return colors.September
	case "October":
		return colors.October
	case "November":
		return colors.November
	case "December":
		return colors.December
	}
	return config.Cfg.Colors.Info
}
