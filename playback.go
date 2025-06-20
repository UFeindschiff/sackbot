package main

import (
// 	"layeh.com/gumble/gumble"
	"layeh.com/gumble/gumbleffmpeg"
	"io"
	"log"
	// "strings"
)

var activesong *gumbleffmpeg.Stream
var playbackChan chan *io.ReadCloser
var loopCurrentSong bool
var targetvolume float32

func playbackThread() {
	playbackChan = make(chan *io.ReadCloser, 4096)
	for song := range playbackChan {
		for ok := true; ok; ok = loopCurrentSong {
			source := gumbleffmpeg.SourceReader(*song)
			stream := gumbleffmpeg.New(client, source)
			stream.Volume = targetvolume
			activesong = stream
			err := activesong.Play()
			if err != nil {
				log.Println(err.Error())
			}
			activesong.Wait()
		}
		activesong = nil
	}
}

func playbackSongHandler(id string) (string, error) {
	audiodata, _, title, err := getAudioDataByID(id)
	if err != nil {
		return title, err
	}
	playbackChan <- audiodata
	return title, nil
}
