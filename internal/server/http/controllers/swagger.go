package controllers

import (
	"net/http"
	"os"
)

func (*Controller) Swagger(w http.ResponseWriter, _ *http.Request) {
	fileBytes, err := os.ReadFile("./api/swagger.yml")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "application/octet-stream")

	_, err = w.Write(fileBytes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
