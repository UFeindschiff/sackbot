package main

import (
	"github.com/kkdai/youtube/v2"
	"errors"
	"io"
	"log"
)

var ytclient *youtube.Client = &youtube.Client{}

func getAudioDataByID(id string) (*io.ReadCloser, int64, string, error) {
	videoStats, err := ytclient.GetVideo(id)
	if err != nil {
		return nil, 0, "", err
		log.Println("Error fetching video: " + err.Error())
	}
	formats := videoStats.Formats.Type("audio/mp4")
	if len(formats) < 1 {
		return nil, 0, videoStats.Title, errors.New("Specified video contains no MP4 Audio Stream")
	}
	stream, size, err := ytclient.GetStream(videoStats, &formats[0])
	if err != nil {
		log.Println("Error fetching audio from video: " + err.Error())
	}
	return &stream, size, videoStats.Title, err
}
