package main

import (
	"github.com/lithdew/youtube"
	"errors"
	"net/http"
	"io"
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
		return "", "", errors.New("Failed to load video ID")
	}
	stream, ok := player.SourceFormats().AudioOnly().BestAudio()
	if !ok {
		return "", player.Title(), errors.New(player.Title() + " does not have an audio-only stream")
	}
	url, err := player.ResolveURL(stream)
	if err != nil {
		return "", player.Title(), errors.New("Failed to get audio URL for " + player.Title())
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
