package model

import (
	"database/sql"
	"errors"
	"log"
)

type YoutubeData struct {
	BaseData

	ChannelID   string `json:"channelID"`
	Title       string `json:"title"`
	ChannelName string `json:"channelName"`

	PublishAt string `json:"publishAt"`
}

func CreateYoutubeData(db *sql.DB) YoutubeData {
	youtubeData := YoutubeData{}
	youtubeData.DB = db
	return youtubeData
}

func (youtubeData YoutubeData) Migrate() error {

	sqlDropTable := "DROP TABLE IF EXISTS public.data CASCADE"

	sqlCreateTable := `
CREATE TABLE public.data
(
    channel_id character varying(200) NOT NULL,
    title character varying(100) NOT NULL,
	channel_name character varying(100) NOT NULL,
    publish_at character varying(60) NOT NULL,
    PRIMARY KEY(channel_id)
)
	`

	_, err := youtubeData.DB.Exec(sqlDropTable)
	if err != nil {
		return err
	}

	_, err = youtubeData.DB.Exec(sqlCreateTable)
	if err != nil {
		return err
	}

	return nil

}

func (youtubeData YoutubeData) Add(channelID string, title string, channelName string, publishAt string) error {
	sql := `insert into public.data (channel_id, title, channel_name, publish_at) values 
	(?, ?, ?, ?)`

	result, err := youtubeData.Exec(sql, channelID, title, channelName, publishAt)
	if err != nil {
		return err
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	log.Println("YoutubeData.Add.AffectedRows=", affectedRows)
	return nil
}

func (youtubeData YoutubeData) fetchRow(cursor *sql.Rows) (YoutubeData, error) {
	b := YoutubeData{}
	err := cursor.Scan(&b.ChannelID, &b.Title, &b.ChannelName, &b.PublishAt)
	if err != nil {
		return YoutubeData{}, err
	}
	return b, nil
}

func (youtubeData YoutubeData) selectQuery() string {
	sql := `select * from public.data order by publish_at desc`
	return sql
}

func (youtubeData YoutubeData) GetListYoutubeData() ([]YoutubeData, error) {

	sql := youtubeData.selectQuery()
	cursor, err := youtubeData.Query(sql)
	if err != nil {
		return nil, err
	}

	youtubeDatas := []YoutubeData{}

	for cursor.Next() {
		c, err := youtubeData.fetchRow(cursor)
		if err != nil {
			return nil, err
		}
		youtubeDatas = append(youtubeDatas, c)
	}

	return youtubeDatas, nil
}

func (youtubeData YoutubeData) FindYoutubeData() (*YoutubeData, error) {
	sql := `select * from public.data where title like '%minyak pertamina%'`

	cursor, err := youtubeData.Query(sql)
	if err != nil {
		return nil, err
	}

	if cursor.Next() {
		c, err := youtubeData.fetchRow(cursor)
		if err != nil {
			return nil, err
		}
		return &c, nil
	}

	return nil, errors.New("data_not_found")
}

func (youtubeData YoutubeData) RemoveByID(channelID string) error {

	sql := "delete from public.data where channel_id=?"
	result, err := youtubeData.Exec(sql, channelID)
	if err != nil {
		return err
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	log.Println("YoutubeData.Remove.AffectedRows=", affectedRows)
	return nil
}

func (youtubeData YoutubeData) Truncate() error {

	sql := "delete from public.data"
	result, err := youtubeData.Exec(sql)
	if err != nil {
		return err
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	log.Println("YoutubeData.Truncate.AffectedRows=", affectedRows)
	return nil
}

func (youtubeData YoutubeData) Update(title string, channelName string, publishAt string, channelID string) error {

	sql := `update public.data
	set title=?,
	channel_name=?,
	publish_at=?
	

	where channel_id=?
	`
	result, err := youtubeData.Exec(sql, title, channelName, publishAt, channelID)
	if err != nil {
		return err
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	log.Println("YoutubeData.Update.AffectedRows=", affectedRows)

	return nil

}
