package main

import (
	"layeh.com/gumble/gumble"
	"layeh.com/gumble/gumbleutil"
	_ "layeh.com/gumble/opus"
	"flag"
	"log"
	"net"
	"crypto/tls"
	"os"
)
var username 	string
var password 	string
var server		string
var channelname	string
var insecureTLS	bool
var client *gumble.Client //client is global so I can access the client from event listener functions without using inline event listners
var terminationChan chan error //writing anything to this will terminate the bot

func init() {
	flag.StringVar(&username, "username", "SackBot", "the username for the bot to use")
	flag.StringVar(&password, "password", "", "server password, should the server require one")
	flag.StringVar(&server, "server", "", "the server to join - needs to be formatted like <domain or IP>:<port>")
	flag.StringVar(&channelname, "channelname", "", "the channel for the bot to join. Will join root channel if not set")
	flag.BoolVar(&insecureTLS, "insecureTLS", false, "skip verification of the mumble server's TLS certificate")
}

func main() {
	flag.Parse()
	terminationChan = make(chan error, 2)
	if server == "" {
		panic("You must specify a server to join")
	}
	config := gumble.NewConfig()
	config.Username = username
	if password != "" {
		config.Password = password
	}
	listener := gumbleutil.Listener{}
	listener.TextMessage = sackBotMessageHandler
	config.Attach(listener)
	var tlsConfig *tls.Config
	tlsConfig = nil
	if insecureTLS {
		tlsConfig = &tls.Config{}
		tlsConfig.InsecureSkipVerify = true
	}
	log.Println("Connecting to server " + server)
	var err error
	client, err = gumble.DialWithDialer(new(net.Dialer), server, config, tlsConfig)
	if err != nil {
		panic(err)
	}
	log.Println("Successfully connected to " + server)
	err = moveSelfToTargetChannel(client, channelname)
	if err != nil {
		log.Println("Failed to join desired channel: " + err.Error())
	}
	targetvolume = 1.0
	go playbackThread()
	for exiterr := range terminationChan {
		if exiterr != nil {
			log.Println("Exiting due to error: " + err.Error())
			os.Exit(1)
		}
		os.Exit(0)
	}
}
