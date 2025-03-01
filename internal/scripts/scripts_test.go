package scripts

import (
	"fmt"
	"statd/config"
	"strings"
	"testing"
	"time"

	"github.com/mshafiee/jalali"
)

func setupCfg() {
	config.Cfg = &config.DefaultConfig
}

func TestDate(t *testing.T) {
	// TODO: test for colors as well as icons (yeah sure)
	setupCfg()
	config.Cfg.Date.Colorize = false
	config.Cfg.Date.Format = "{{.gdate}} --- {{.jdate}}"
	config.Cfg.Date.GDate = "a=%a A=%A d=%d b=%b B=%B m=%m y=%y Y=%Y H=%H I=%I p=%p M=%M S=%M"
	config.Cfg.Date.JDate = "a=%a A=%A d=%d b=%b B=%B m=%m y=%y Y=%Y I=%I p=%p"

	now := time.Now()
	got, err := Date()
	if err != nil {
		t.Errorf("Date() produced error. %v", err)
	}
	gotSlice := strings.Split(got, "---")
	if len(gotSlice) != 2 {
		t.Errorf("templating gone wrong. %s", got)
	}
	gGot, jGot := strings.TrimSpace(gotSlice[0]), strings.TrimSpace(gotSlice[1])
	gParts := strings.Split(gGot, " ")
	if len(gParts) != 13 {
		t.Errorf("Formatting Gregorian date didn't produce the expected %d number. gGot=%s", 13, gGot)
	}
	checkG := func(ndx int, format string) {
		if val := now.Format(format); gParts[ndx] != val {
			target := fmt.Sprintf("%%%c", format[0])
			t.Errorf("Gregorian:: %s didn't work. got=%s != want=%s", target, gParts[ndx], val)
		}
	}
	checkG(0, "a=Mon")
	checkG(1, "A=Monday")
	checkG(2, "d=02")
	checkG(3, "b=Jan")
	checkG(4, "B=January")
	checkG(5, "m=01")
	checkG(6, "y=06")
	checkG(7, "Y=2006")
	checkG(8, "H=15")
	checkG(9, "I=03")
	checkG(10, "p=PM")
	checkG(11, "M=04")
	// checkG(12, "S=05") skipping seconds for obvious complications

	jNow := jalali.JalaliFromTime(now)
	jParts := strings.Split(jGot, " ")
	if len(jParts) != 10 {
		t.Errorf("Formatting Jalali date didn't produce the expected %d number. jGot=%s", 10, jGot)
	}
	checkJ := func(ndx int, format string, val string) {
		if val == "" {
			val = jNow.Format(format)
		} else {
			val = fmt.Sprintf("%c=%s", format[0], val)
		}
		if jParts[ndx] != val {
			target := fmt.Sprintf("%%%c", format[0])
			t.Errorf("Jalali:: %s didn't work. got=%s != want=%s", target, jParts[ndx], val)
		}
	}
	checkJ(0, "a", jNow.Weekday().String()[:3])
	checkJ(1, "A", jNow.Weekday().String())
	checkJ(2, "d=%d", "")
	checkJ(3, "b", jNow.Month().String()[:3])
	checkJ(4, "B", jNow.Month().String())
	checkJ(5, "m=%m", "")
	checkJ(6, "y=%y", "")
	checkJ(7, "Y=%Y", "")
	checkJ(8, "I", now.Format("03"))
	checkJ(9, "p", now.Format("PM"))
}
