package main

import (
	"layeh.com/gumble/gumble"
	"log"
	"strings"
	"strconv"
)

func sackBotMessageHandler(e *gumble.TextMessageEvent) {
	if strings.HasPrefix(e.Message, "!") {
		splitstring := strings.Split(e.Message, " ")
		switch splitstring[0] {
			case "!add":
				if len(splitstring) < 2 {
					client.Self.Channel.Send("You need to provide an ID", false)
				}
				videoID := extractVideoIDFromCommand(splitstring)
				title, err := playbackSongHandler(videoID)
				if err != nil {
					client.Self.Channel.Send("Error adding requested ID \"" + videoID +"\": " + err.Error(), false)
					log.Println("Error adding requested ID \"" + videoID +"\": " + err.Error())
				} else {
					client.Self.Channel.Send("Added " + title + " to playback queue", false)
					log.Println("Added " + title + " to playback queue")
				}
			case "!loop":
				loopCurrentSong = !loopCurrentSong
				if loopCurrentSong {
					client.Self.Channel.Send("Enabled looping title", false)
					log.Println("Enabled looping title")
				} else {
					client.Self.Channel.Send("Disabled looping title", false)
					log.Println("Disabled looping title")
				}
			case "!info":
				client.Self.Channel.Send("SackBot by UFeindschiff\nSource available under github.com/UFeindschiff/sackbot", false)
			case "!pause":
				if activesong == nil {
					client.Self.Channel.Send("Currently no song being played", false)
					break;
				}
				activesong.Pause()
				client.Self.Channel.Send("Paused playback", false)
				log.Println("Paused playback")
			case "!resume":
				if activesong == nil {
					client.Self.Channel.Send("Currently no song being paused", false)
					break;
				}
				activesong.Play()
				client.Self.Channel.Send("Resumed playback", false)
				log.Println("Resumed playback")
			case "!skip":
				if activesong == nil {
					client.Self.Channel.Send("Currently no song being played", false)
					break;
				}
				activesong.Stop()
				client.Self.Channel.Send("Skipped playback", false)
				log.Println("Skipped playback")
			case "!help":
				client.Self.Channel.Send("SackBot usage: \n!add <Youtube-ID> adds a song to the queue\n!loop toggles looping the current song\n!info displays info about the bot\n!pause pauses playback\n!resume resumes playback\n!skip skips the current song\n!help prints this help message\n!quit exits the bot", false)
			case "!volume":
				if len(splitstring) < 2 {
					client.Self.Channel.Send("You need to provide a volume level", false)
				}
				newvolume, err := strconv.ParseFloat(splitstring[1], 32)
				if err != nil {
					client.Self.Channel.Send("Failed to parse argument. Make sure it is a valid number", false)
				} else if newvolume > 2.0 {
					client.Self.Channel.Send("Volume must be 2.0 or lower (you won't make people go deaf here)", false)
				} else {
					if activesong != nil {
						activesong.Volume = float32(newvolume)
					}
					targetvolume = float32(newvolume)
					client.Self.Channel.Send("Setting volume to " + splitstring[1], false)
					log.Println("Setting volume to " + splitstring[1])
				}
				
			case "!quit":
				client.Self.Channel.Send("Quitting... Have a nice day :)", false)
				log.Println("Got command to exit from user. Exiting...")
				terminationChan <- nil
			default:
				client.Self.Channel.Send("Unknown command", false)
		}
	}
}

//This is needed as the Mumble Client supports HTML rendering and by default puts URLs as an HTML hyperlink, so I have to check whether the user just sent an ID or a hyperlink.
func extractVideoIDFromCommand(splitcommand []string) string {
	if len(splitcommand) < 2 {
		//Nothing but the command prefix provided, return empty string
		return ""
	}
	if splitcommand[1] != "<a" {
		//We don't have a hyperlink, so just use what we got as videoID
		return splitcommand[1]
	}
	if len(splitcommand) < 3 {
		//There is no actual link, but just <a was provided. Without this check, the bot would panic
		return ""
	}
	splitstring := strings.Split(splitcommand[2], ">")
	if len(splitstring) < 2 {
		return ""
	}
	return strings.TrimSuffix(splitstring[1], "</a")
}
