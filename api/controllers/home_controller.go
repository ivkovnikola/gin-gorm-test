package controllers

import (
	"net/http"

	"github.com/ivkovnikola/testgormgin/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome")
}
