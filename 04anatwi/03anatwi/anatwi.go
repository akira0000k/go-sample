package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/user"

	"github.com/labstack/echo"
	"github.com/ChimeraCoder/anaconda"
)

func main(){

	e := echo.New()
	e.POST("/tweet", search)
	e.Logger.Fatal(e.Start(":1323"))
}

func search(c echo.Context) error {
	keyword := c.FormValue("keyword")
	fmt.Println("keyword=", keyword)
	
	api := connectTwitterApi()

	searchResult, _ := api.GetSearch(`"` +keyword+ `"`, nil)

	tweets := make([]*Tweet, 0)

	for idx, data := range searchResult.Statuses {
		tweet := new(Tweet)
		tweet.Text = data.FullText
		tweet.User = data.User.Name
		if idx == 0 {
			//fmt.Printf("%#v\n", data)
			fmt.Printf("%#v\n", tweet)
			return c.JSON(http.StatusOK, data)
		}

		tweets = append(tweets, tweet)
	}

	return c.JSON(http.StatusOK, tweets)
}

func connectTwitterApi() *anaconda.TwitterApi {
	usr, _ := user.Current()
	raw, error := ioutil.ReadFile(usr.HomeDir + "/twitter/twitterAccount.json")

	if error != nil {
		fmt.Println(error.Error())
		return nil
	}

	var twitterAccount TwitterAccount
	json.Unmarshal(raw, &twitterAccount)

	return anaconda.NewTwitterApiWithCredentials(
		twitterAccount.AccessToken,
		twitterAccount.AccessTokenSecret,
		twitterAccount.ConsumerKey,
		twitterAccount.ConsumerSecret)

}



type TwitterAccount struct {
	AccessToken string `json:"accessToken"`
	AccessTokenSecret string `json:"accessTokenSecret"`
	ConsumerKey string `json:"consumerKey"`
	ConsumerSecret string `json:"consumerSecret"`
}

type Tweet struct {
	User string `json:"user"`
	Text string `json:"text"`
}

//type Tweets *[]Tweet
