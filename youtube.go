package main

import (
	"github.com/UFeindschiff/youtube"
	"errors"
	"net/http"
	"io"
	"log"
)

func getAudioDataByIdWrapper(id string) (*io.ReadCloser, string, error) {
	url, title, err := getAudioURLAndData(id)
	if err != nil {
		return nil, title, err
	}
	retval, err := getAudioFromURL(url)
	if err != nil {
		return nil, title, err
	}
	return retval, title, nil
}

func getAudioURLAndData(id string) (string, string, error) {
	player, err := youtube.Load(youtube.StreamID(id))
	if err != nil {
		//return "", "", errors.New("Failed to load video ID")
		return "", "", err
	}
	streams := player.SourceFormats().AudioOnly().SortByAudioQuality()
	for _, stream := range streams {
		log.Println(stream)
		url, err := player.ResolveURL(stream)
		
		if err == nil {
			return url, player.Title(), nil
		}
	}
	if no_video_fallback {
		return "", player.Title(), errors.New("Failed to get audio URL for " + player.Title() + ": " + err.Error())
	} else {
		return getMuxedVideoURLAndData(id)
	}
}

func getMuxedVideoURLAndData(id string) (string, string, error) {
	player, err := youtube.Load(youtube.StreamID(id))
	if err != nil {
		//return "", "", errors.New("Failed to load video ID")
		return "", "", err
	}
	stream, ok := player.MuxedFormats().BestAudio()
	if !ok {
		return "", player.Title(), errors.New(player.Title() + " does not have any streams")
	}
	url, err := player.ResolveURL(stream)
	if err != nil {
		return "", player.Title(), errors.New("Failed to get muxed video URL for " + player.Title() + ": " + err.Error())
	}
	return url, player.Title(), nil
}

func getAudioFromURL(url string) (*io.ReadCloser, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New("Request succeeded, but received non-OK Response: " + resp.Status)
	}
	return &resp.Body, nil
}
