package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/user"

	"github.com/ChimeraCoder/anaconda"
)

func main(){
	usr, _ := user.Current()
	raw, error := ioutil.ReadFile(usr.HomeDir + "/twitter/twitterAccount.json")

	if error != nil {
		fmt.Println(error.Error())
		return
	}
	//fmt.Println("raw:\n", string(raw))
	
	var twitterAccount TwitterAccount
	json.Unmarshal(raw, &twitterAccount)

	//fmt.Println(twitterAccount)
	
	api := anaconda.NewTwitterApiWithCredentials(
		twitterAccount.AccessToken,
		twitterAccount.AccessTokenSecret,
		twitterAccount.ConsumerKey,
		twitterAccount.ConsumerSecret)

	//fmt.Println(api)
	
	searchResult, error := api.GetSearch(`スタートアップ`, nil)
	if error != nil {
		fmt.Println(error.Error())
		return
	}
	
	for i, tweet := range searchResult.Statuses {
		fmt.Println("i=", i)
		fmt.Println(tweet.Text)
	}
}

type TwitterAccount struct {
	AccessToken string `json:"accessToken"`
	AccessTokenSecret string `json:"accessTokenSecret"`
	ConsumerKey string `json:"consumerKey"`
	ConsumerSecret string `json:"consumerSecret"`
}
