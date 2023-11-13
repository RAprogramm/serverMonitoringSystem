// Package handlers provides HTTP handlers for various functionalities.
package handlers

import (
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/pusher/pusher-http-go"
)

// PusherAuth method handles authentication for Pusher presence channels.
func (repo *DBRepo) PusherAuth(w http.ResponseWriter, r *http.Request) {
	// Get the user ID from the session
	userID := repo.App.Session.GetInt(r.Context(), "userID")

	// Get user information from the database
	u, _ := repo.DB.GetUserById(userID)

	// Read the request body
	params, _ := io.ReadAll(r.Body)

	// Create presence data for the Pusher member
	presenceData := pusher.MemberData{
		UserID: strconv.Itoa(userID),
		UserInfo: map[string]string{
			"name": u.FirstName,
			"id":   strconv.Itoa(userID),
		},
	}

	// Authenticate presence channel with Pusher
	response, err := app.WsClient.AuthenticatePresenceChannel(params, presenceData)
	if err != nil {
		log.Println("Error in response")
		log.Println(err)
		return
	}

	// Set the response header and write the response
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(response)
}
