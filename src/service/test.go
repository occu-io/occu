package service

import (
	"encoding/json"
	"net/http"
	"util/models"
)

func TestEndpoint(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Exception{Message: "hello world"})
}
