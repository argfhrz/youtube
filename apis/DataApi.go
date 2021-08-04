package apis

import (
	"encoding/json"
	"log"
	"net/http"
	"youtube/config"
	"youtube/connection"
	"youtube/model"
)

type DataApi struct {
	BaseApi

	model.YoutubeData
}

func (dataApi DataApi) GetData(w http.ResponseWriter, r *http.Request) {

	db := connection.OpenConnection(config.DEV)
	defer db.Close()

	dataModel := model.CreateYoutubeData(db)
	datas, err := dataModel.GetListYoutubeData()
	if err != nil {
		dataApi.Error(w, err)
		return
	} else {
		dataApi.Json(w, datas, http.StatusOK)
		return
	}

}

func (dataApi DataApi) Find(w http.ResponseWriter, r *http.Request) {

	db := connection.OpenConnection(config.DEV)
	defer db.Close()

	dataModel := model.CreateYoutubeData(db)
	datas, err := dataModel.FindYoutubeData()
	if err != nil {
		dataApi.Error(w, err)
		return
	} else {
		dataApi.Json(w, datas, http.StatusOK)
		return
	}

}

func (dataApi DataApi) RemoveDataByID(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&dataApi)
	if err != nil {
		dataApi.Error(w, err)
		return
	}

	db := connection.OpenConnection(config.DEV)
	defer db.Close()

	channelID := dataApi.QueryParam(r, "channelID")

	dataModel := model.CreateYoutubeData(db)
	err = dataModel.RemoveByID(channelID)
	if err != nil {
		dataApi.Error(w, err)
		return
	} else {
		log.Println("remove_success")
		return
	}

}

func (dataApi DataApi) PostData(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&dataApi)
	if err != nil {
		dataApi.Error(w, err)
		return
	}

	db := connection.OpenConnection(config.DEV)
	defer db.Close()

	dataModel := model.CreateYoutubeData(db)

	err = dataModel.Add(dataApi.ChannelID, dataApi.Title, dataApi.ChannelName, dataApi.PublishAt)
	if err != nil {
		dataApi.Error(w, err)
		return
	} else {
		log.Println("post_sucess")
	}

}

func (dataApi DataApi) UpdateData(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&dataApi)
	if err != nil {
		dataApi.Error(w, err)
		return
	}

	db := connection.OpenConnection(config.DEV)
	defer db.Close()

	dataModel := model.CreateYoutubeData(db)

	err = dataModel.Update(dataApi.ChannelID, dataApi.Title, dataApi.ChannelName, dataApi.PublishAt)
	if err != nil {
		dataApi.Error(w, err)
		return
	} else {
		log.Println("update_sucess")
	}

}
