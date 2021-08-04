package model

import (
	"testing"
	"youtube/config"
	"youtube/connection"
)

func TestData(t *testing.T) {
	db := connection.OpenConnection(config.DEV)
	defer db.Close()

	youtubeDataModel := CreateYoutubeData(db)
	// err := youtubeDataModel.Migrate()
	// if err != nil {
	// 	t.Fatal(err.Error())
	// }

	// channelID := "02"
	// title := "video2"
	// channelName := "channel2"
	// publishAt := time.Now().UTC().String()

	//Add
	// err := youtubeDataModel.Add(channelID, title, channelName, publishAt)
	// if err != nil {
	// 	t.Fatal(err.Error())
	// }

	//Get
	// youtubeData, err := youtubeDataModel.GetListYoutubeData()
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// log.Println(helpers.ToJson(youtubeData))

	//Find
	// data, err := youtubeDataModel.FindYoutubeData()
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// log.Println(helpers.ToJson(data))

	channelID := "02"
	// title := "video3-3"
	// channelName := "channel3-3"
	// publishAt := time.Now().UTC().String()

	// err := youtubeDataModel.Update(title, channelName, publishAt, channelID)
	// if err != nil {
	// 	t.Fatal(err.Error())
	// }

	err := youtubeDataModel.RemoveByID(channelID)
	if err != nil {
		t.Fatal(err.Error())
	}

}
