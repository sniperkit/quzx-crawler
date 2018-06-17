package services

import (
	"log"
	"net/url"

	"github.com/ChimeraCoder/anaconda"

	"github.com/sniperkit/quzx-crawler/pkg/quzxutil"
)

type TwitterService struct {
}

func (s *TwitterService) GetFavoritesTwits(name string) ([]anaconda.Tweet, error) {

	// TODO: refactore (move to function)
	consumer_key := quzxutil.GetParameter("TWICONKEY")
	consumer_secret := quzxutil.GetParameter("TWICONSEC")
	access_token := quzxutil.GetParameter("TWIACCTOK")
	access_token_secret := quzxutil.GetParameter("TWIACCTOKSEC")

	anaconda.SetConsumerKey(consumer_key)
	anaconda.SetConsumerSecret(consumer_secret)
	api := anaconda.NewTwitterApi(access_token, access_token_secret)

	v := url.Values{}
	v.Set("screen_name", name)

	tweets, err := api.GetFavorites(v)
	if err != nil {
		log.Println("Get twitter favorites returned error: %s", err.Error())
	}

	return tweets, err
}

func (s *TwitterService) DestroyFavorites(id int64) {

	consumer_key := quzxutil.GetParameter("TWICONKEY")
	consumer_secret := quzxutil.GetParameter("TWICONSEC")
	access_token := quzxutil.GetParameter("TWIACCTOK")
	access_token_secret := quzxutil.GetParameter("TWIACCTOKSEC")

	anaconda.SetConsumerKey(consumer_key)
	anaconda.SetConsumerSecret(consumer_secret)
	api := anaconda.NewTwitterApi(access_token, access_token_secret)

	_, err := api.Unfavorite(id)
	if err != nil {
		log.Println(err)
	}
}
