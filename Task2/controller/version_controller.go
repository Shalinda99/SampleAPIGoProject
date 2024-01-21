package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"task2/service"
)

type VersionController struct {
	versionService *service.VersionService
}

func NewVersionController(initialVersion string) *VersionController {
	return &VersionController{
		versionService: service.NewVersionService(initialVersion),
	}
}

func (vc *VersionController) HandleRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		vc.handleGetVersion(w)
	case http.MethodPut:
		vc.handlePutVersion(w, r)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func (vc *VersionController) handleGetVersion(w http.ResponseWriter) {
	currentVersion := vc.versionService.GetCurrentVersion()
	fmt.Fprintf(w, "Current service version: %s", currentVersion)
}

func (vc *VersionController) handlePutVersion(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	newVersion := string(body)
	if newVersion == "" {
		http.Error(w, "Missing 'new_version' parameter in the request body", http.StatusBadRequest)
		return
	}

	vc.versionService.UpdateVersion(newVersion)
	fmt.Fprintf(w, "Service version updated to: %s", newVersion)
}
