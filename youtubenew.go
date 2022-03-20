package main

import (
	"github.com/kkdai/youtube/v2"
	"errors"
// 	"net/http"
	"io"
	"log"
)

func getNewAudioDataByIdWrapper(id string) (*io.ReadCloser, string, error) {
	client := youtube.Client{}
	video, err := client.GetVideo(id)
	if err != nil {
		return nil, "", errors.New("Failed to get Video for URL " + err.Error())
	}
	
	audioformats := video.Formats.Type("audio")
	audioformats.Sort()
	log.Println(audioformats)
	
	stream, _, err := client.GetStream(video, &audioformats[0])
	if err != nil {
		return nil, video.Title, errors.New("Failed to get audio stream: " + err.Error())
	}
	
	return &stream, video.Title, nil
}
