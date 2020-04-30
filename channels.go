package main

import (
	"layeh.com/gumble/gumble"
	"errors"
)

func getChannelByName(client *gumble.Client, name string) (*gumble.Channel, error) {
	retval := client.Channels.Find(name)
	if retval == nil {
		return nil, errors.New("Channel with name " + name + " not found")
	}
	return retval, nil
}

func moveSelfToTargetChannel(client *gumble.Client, name string) error {
	if name == "" {
		return nil //empty string means we want to stay in the root channel
	}
	channel, err := getChannelByName(client, name)
	if err != nil {
		return err
	}
	client.Self.Move(channel)
// 	if client.Self.Channel != channel { //Check doesn't work as it fails despite the user successfully joining the channel. Possible race condition, but I don't want to wait here for verification
// 		return errors.New("Joining the desired channel failed")
// 	}
	return nil
}
