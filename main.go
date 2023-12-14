package main

import (
	"os"
	"os/signal"
	"reflect"

	"github.com/paralin/go-dota2"
	"github.com/paralin/go-steam"
	"github.com/paralin/go-steam/steamid"
	"github.com/sirupsen/logrus"
	"github.com/paralin/go-dota2/protocol"
)

func main() {
	// Retrieve Steam credentials from environment variables
	steamUsername := os.Getenv("STEAM_USERNAME")
	steamPassword := os.Getenv("STEAM_PASSWORD")

	if steamUsername == "" || steamPassword == "" {
		logrus.Fatal("Steam credentials are not set as environment variables yo")
	}

	// Create a Logrus logger
	logger := logrus.New()
	logger.Out = os.Stdout
	logger.Level = logrus.DebugLevel
	logger.Formatter = &logrus.TextFormatter{
		FullTimestamp: true,
	}

	// Initialize the Steam client
	client := steam.NewClient()
	dota2Client := dota2.New(client, logger)

	// Connect to Steam
	// ☢️☢️☢️ Here's where things get hairy, and I'm getting timeouts ☢️☢️☢️
	client.Connect()
	
	// Event handling
	go func() {
		for event := range client.Events() {
			switch e := event.(type) {
			case *steam.ConnectedEvent:
				logrus.Println("Connected to Steam")
				client.Auth.LogOn(&steam.LogOnDetails{
					Username: steamUsername,
					Password: steamPassword,
				})
			case *steam.LoggedOnEvent:
				logrus.Println("Logged on to Steam")
				// Example: Creating a lobby after logging in
				createLobby(dota2Client)
			case *steam.DisconnectedEvent:
				logrus.Println("Disconnected from Steam")
				return
			case steam.FatalErrorEvent:
				logrus.Errorf("Fatal error: %v", e)
				return
			default:
				logrus.Printf("Received event: %s\n", reflect.TypeOf(e).String())
			}
		}
	}()

	// Channel to block main goroutine and handle graceful shutdown
	stopChan := make(chan struct{})
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	go func() {
		<-signalChan // Wait for an OS signal
		close(stopChan) // Close the stopChan to unblock the main goroutine
	}()

	<-stopChan // Block here until stopChan is closed
	// Perform any cleanup if necessary
}

// createLobby creates a Dota 2 lobby
func createLobby(dota2Client *dota2.Dota2) {
	gameName := "NADCL test"
	details := &protocol.CMsgPracticeLobbySetDetails{
		GameName: &gameName, // Set the address of the string variable
		// Set other lobby details as needed
	}
	dota2Client.CreateLobby(details)
	logrus.Println("Lobby created")
}

// inviteToLobby invites a player to the current lobby
func inviteToLobby(dota2Client *dota2.Dota2, steamID string) {
	sid, err := steamid.NewId(steamID)
	if err != nil {
		logrus.Errorf("Invalid SteamID: %v", err)
		return
	}
	dota2Client.InviteLobbyMember(sid)
	logrus.Printf("Invited user %s to the lobby\n", steamID)
}
