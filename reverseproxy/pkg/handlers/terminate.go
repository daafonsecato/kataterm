package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
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

	err := controller.deletePodAndServicesbySessionID(w, req.SessionID)
	if err != nil {
		log.Printf("Failed to delete pod and services: %v", err)
		return
	}

}
func (controller *SessionController) deletePodAndServicesbySessionID(w http.ResponseWriter, sessionID string) error {

	// Get the PodName using the GetPodNameBySessionId function from the models package
	podName, err := controller.sessionStore.GetPodNameBySessionId(sessionID)
	if err != nil {
		log.Printf("Failed to get PodName: %v", err)
		http.Error(w, fmt.Sprintf("Error getting PodName: %v", err), http.StatusInternalServerError)
		return err
	}

	// Remove the pod by PodName
	err = controller.clientset.CoreV1().Pods(controller.namespace).Delete(context.TODO(), podName, metav1.DeleteOptions{})
	if err != nil {
		log.Printf("Failed to delete pod: %v", err)
		http.Error(w, fmt.Sprintf("Error deleting pod: %v", err), http.StatusInternalServerError)
		return err
	}

	beSvcName := fmt.Sprintf("backend-svc-%s", sessionID)
	ttydSvcName := fmt.Sprintf("ttyd-svc-%s", sessionID)
	codeSvcName := fmt.Sprintf("codeeditor-svc-%s", sessionID)

	// Delete the backend service
	err = controller.clientset.CoreV1().Services(controller.namespace).Delete(context.TODO(), beSvcName, metav1.DeleteOptions{})
	if err != nil {
		log.Printf("Failed to delete backend service: %v", err)
		http.Error(w, fmt.Sprintf("Error deleting backend service: %v", err), http.StatusInternalServerError)
		return err
	}

	// Delete the ttyd service
	err = controller.clientset.CoreV1().Services(controller.namespace).Delete(context.TODO(), ttydSvcName, metav1.DeleteOptions{})
	if err != nil {
		log.Printf("Failed to delete ttyd service: %v", err)
		http.Error(w, fmt.Sprintf("Error deleting ttyd service: %v", err), http.StatusInternalServerError)
		return err
	}

	// Delete the code editor service
	err = controller.clientset.CoreV1().Services(controller.namespace).Delete(context.TODO(), codeSvcName, metav1.DeleteOptions{})
	if err != nil {
		log.Printf("Failed to delete code editor service: %v", err)
		http.Error(w, fmt.Sprintf("Error deleting code editor service: %v", err), http.StatusInternalServerError)
		return err
	}

	// Delete the IngressRoute
	ingressRouteName := fmt.Sprintf("ingressroute-%s", sessionID)
	// Define the GroupVersionResource (GVR)
	gvr := schema.GroupVersionResource{
		Group:    "traefik.containo.us",
		Version:  "v1alpha1",
		Resource: "ingressroutes",
	}
	err = controller.dynamiccclient.Resource(gvr).Namespace(controller.namespace).Delete(context.TODO(), ingressRouteName, metav1.DeleteOptions{})
	if err != nil {
		log.Printf("Failed to delete IngressRoute: %v", err)
		http.Error(w, fmt.Sprintf("Error deleting IngressRoute: %v", err), http.StatusInternalServerError)
		return err
	}

	// Send a success response
	fmt.Fprintf(w, "Pod deleted successfully for session ID: %s", sessionID)
	return nil
}
func (controller *SessionController) TerminateMultipleMachinesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	type MultipleTerminationRequest struct {
		SessionIDs string `json:"session_ids"`
	}

	var req MultipleTerminationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	if req.SessionIDs == "" {
		http.Error(w, "Session IDs are required", http.StatusBadRequest)
		return
	}

	sessionIDs := strings.Split(req.SessionIDs, " ")

	for _, sessionID := range sessionIDs {
		err := controller.deletePodAndServicesbySessionID(w, sessionID)
		if err != nil {
			log.Printf("Failed to delete pod and services for session ID %s: %v", sessionID, err)
			return
		}
	}

	// Send a success response
	fmt.Fprintf(w, "Pods deleted successfully for session IDs: %s", req.SessionIDs)
}
