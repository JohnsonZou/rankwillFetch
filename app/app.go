package app

import (
	"github.com/go-co-op/gocron"
	"log"
	"rankwillFetch/fetching"
	"strconv"
	"time"
)

func AppRun() {
	s := gocron.NewScheduler(time.Local)

	s.Every(1).Week().Saturday().At("23:30:30").Do(BiweekPrepare)
	s.Every(1).Week().Sunday().At("00:10:30").Do(BiweekRun)
	s.Every(1).Week().Sunday().At("10:30:30").Do(WeeklyPrepare)
	s.Every(1).Week().Sunday().At("12:10:30").Do(WeeklyRun)
	s.StartAsync()
}

func GetBiWeekNameByTime() string {
	now := time.Now().Unix()
	bw := ((now - 1676730580) / 1209600) + 98
	res := "biweekly-contest-" + strconv.Itoa(int(bw))
	return res
}

func GetWeeklyByTime() string {
	now := time.Now().Unix()
	w := ((now - 1677378580) / 604800) + 334
	res := "weekly-contest-" + strconv.Itoa(int(w))
	return res
}

var biweekTimer1 bool = false
var biweekTimer2 bool = false
var contestantnum int

func BiweekRun() {
	if biweekTimer1 {
		log.Println(GetBiWeekNameByTime(), "  Running")
		fetching.ChannelStart(GetBiWeekNameByTime(), false)
		biweekTimer1 = false
	} else {
		biweekTimer1 = true
	}
}
func BiweekPrepare() {
	if biweekTimer2 {
		log.Println(GetBiWeekNameByTime(), "  Preparing")
		fetching.ChannelStart(GetBiWeekNameByTime(), true)
		biweekTimer2 = false
	} else {
		biweekTimer2 = true
	}
}

func WeeklyRun() {
	log.Println(GetWeeklyByTime(), "  Running")
	fetching.ChannelStart(GetWeeklyByTime(), false)
}
func WeeklyPrepare() {
	log.Println(GetWeeklyByTime(), "  Preparing")
	fetching.ChannelStart(GetWeeklyByTime(), true)
}
