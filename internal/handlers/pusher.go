package handlers

import "net/http"

// PusherAuth method ...
func (repo *DBRepo) PusherAuth(w http.ResponseWriter, r *http.Request) {
    userID := repo.App.Session.GetInt()
}

