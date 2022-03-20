# SackBot - a super-simple mumble music bot
SackBot is a very simple mumble music bot born out of necessity as the MumbleDJ bot is unmaintained and no longer functional.
It has very few features compared to the MumbleDJ bot, but also has a few advantages
* does not require any external tool to be installed in order to operate (MumbleDJ needed youtube-dl)
* does not require a YouTube API-Key in order to play back the audio of YouTube videos
The only supported audio source so far is YouTube

## Features
* Playback audio from any YouTube video in a Mumble channel
* Pause and Resume Playback mid-song
* Songs are skippable

## Bugs and issues
* Playback of a song occasionally stops shortly before the song is actually over - need to investigate the cause
* Looping songs is currently broken
* Can currently only join in either the root channel or a direct child channel of the root channel
* No permission handling at all - every user can issue the bot every command

## Usage
```
Usage of ./sackbot:
  -channelname string
        the channel for the bot to join. Will join root channel if not set
  -insecureTLS
        skip verification of the mumble server's TLS certificate
  -legacyFetching
        use the legacy way of fetching audio from a YouTube stream. Tends to fail for many videos and takes longer, but may result in better audio quality
  -no_video_fallback
        do not fall back on grabbing a video in case no audio-only stream is available
  -password string
        server password, should the server require one
  -server string
        the server to join - needs to be formatted like <domain or IP>:<port>
  -username string
        the username for the bot to use (default "SackBot")
```

## Building
Building requires Go 1.12 or newer. Simply clone the repository and run `make`. Go's dependency system will take care of fetching any library dependencies

## License
SackBot is released under the terms of the 2-clause BSD license. It uses the [gumble](https://github.com/layeh/gumble) library which is licensed under the terms of the Mozilla Public License 2.0
