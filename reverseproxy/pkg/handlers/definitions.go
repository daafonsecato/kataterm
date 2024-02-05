package handlers

import (
	db "github.com/daafonsecato/kataterm-reverseproxy/internal/database"
	"github.com/daafonsecato/kataterm-reverseproxy/pkg/models"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type SessionController struct {
	sessionStore *models.SessionStore
	namespace    string
	clientset    *kubernetes.Clientset
}

func NewSessionController() *SessionController {
	db, err := db.InitDB()
	if err != nil {
		panic("Error initializing DB")
	}

	sessionStore := models.NewSessionStore(db)
	namespace := "default"

	// Create the rest.Config from controller.kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", "/app/rest.config")
	if err != nil {
		panic("Error building kubeConfig from flags: " + err.Error())
	}

	// Create the Kubernetes clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic("Error creating Kubernetes clientset")
	}

	return &SessionController{
		sessionStore: sessionStore,
		namespace:    namespace,
		clientset:    clientset,
	}
}
