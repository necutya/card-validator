package controllers

import "net/http"

func (c *Controller) HealthCheck(w http.ResponseWriter, r *http.Request) {
	c.sendResponse(w, r, http.StatusOK, map[string]string{"message": "OK"})
}
