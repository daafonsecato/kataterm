package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type TerminationRequest struct {
	SessionID string `json:"session_id"`
}

func (controller *SessionController) TerminateMachineHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req TerminationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	if req.SessionID == "" {
		http.Error(w, "Session ID is required", http.StatusBadRequest)
		return
	}

	// Get the InstanceID using the GetAWSInstanceID function from the models package
	instanceID, err := controller.sessionStore.GetAWSInstanceID(req.SessionID)
	if err != nil {
		log.Printf("Failed to get InstanceID: %v", err)
		http.Error(w, fmt.Sprintf("Error getting InstanceID: %v", err), http.StatusInternalServerError)
		return
	}

	// Call the terminateInstance function from the services package with the obtained InstanceID
	err = controller.AWSService.TerminateInstance(instanceID)
	if err != nil {
		log.Printf("Failed to terminate m achine: %v", err)
		http.Error(w, fmt.Sprintf("Error terminating machine: %v", err), http.StatusInternalServerError)
		return
	}

	// Send a success response
	fmt.Fprintf(w, "Machine terminated successfully for session ID: %s", req.SessionID)
}
