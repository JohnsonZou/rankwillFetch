package main

import (
	_ "github.com/PuerkitoBio/goquery"
	_ "github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"os"
	"rankwillFetch/common"
	"rankwillFetch/fetching"
)

func init() {
	InitConfig()
	_ = common.InitRedis()
	_ = common.InitDB()
}
func main() {
	//wg := sync.WaitGroup{}
	//wg.Add(1)
	//go func() {
	//	app.AppRun()
	//}()
	//wg.Wait()

	fetching.ChannelStart("biweekly-contest-99", false)
	fetching.ChannelStart("weekly-contest-335", false)
	fetching.ChannelStart("biweekly-contest-98", false)
	fetching.ChannelStart("weekly-contest-334", false)
	fetching.ChannelStart("biweekly-contest-97", false)
	fetching.ChannelStart("weekly-contest-333", false)
	fetching.ChannelStart("biweekly-contest-96", false)
	fetching.ChannelStart("weekly-contest-332", false)
	fetching.ChannelStart("biweekly-contest-95", false)
	fetching.ChannelStart("weekly-contest-331", false)
	fetching.ChannelStart("biweekly-contest-94", false)
	fetching.ChannelStart("weekly-contest-330", false)

}
func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "\\config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
