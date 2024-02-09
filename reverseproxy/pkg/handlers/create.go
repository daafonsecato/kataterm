package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func (controller *SessionController) CreateKubernetesPodHandler(w http.ResponseWriter, r *http.Request) {

	PodUuid := uuid.New().String()
	podName := fmt.Sprintf("backend-pod-%s", PodUuid)
	beSvcName := fmt.Sprintf("backend-svc-%s", PodUuid)
	ttydSvcName := fmt.Sprintf("ttyd-svc-%s", PodUuid)
	codeSvcName := fmt.Sprintf("codeeditor-svc-%s", PodUuid)
	ingressRouteName := fmt.Sprintf("ingressroute-%s", PodUuid)
	beHostName := fmt.Sprintf("Host(`%s.backend.terminal.kataterm.com`) && PathPrefix(`/`)", PodUuid)
	ttydHostName := fmt.Sprintf("Host(`%s.ttyd.terminal.kataterm.com`) && PathPrefix(`/`)", PodUuid)
	codeeditorHostName := fmt.Sprintf("Host(`%s.codeeditor.terminal.kataterm.com`) && PathPrefix(`/`)", PodUuid)

	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: podName,
			Labels: map[string]string{
				"app": podName,
			},
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  "backend",
					Image: "daafonsecato/kataterm-backend:v2",
					Ports: []corev1.ContainerPort{
						{
							ContainerPort: 8000,
						},
					},
					Env: []corev1.EnvVar{
						{
							Name: "DB_HOST",
							ValueFrom: &corev1.EnvVarSource{
								SecretKeyRef: &corev1.SecretKeySelector{
									LocalObjectReference: corev1.LocalObjectReference{
										Name: "my-secret",
									},
									Key: "DB_HOST",
								},
							},
						},
						{
							Name: "VALIDATOR_HOST",
							ValueFrom: &corev1.EnvVarSource{
								SecretKeyRef: &corev1.SecretKeySelector{
									LocalObjectReference: corev1.LocalObjectReference{
										Name: "my-secret",
									},
									Key: "VALIDATOR_HOST",
								},
							},
						},
						{
							Name: "GITKATAS_HOST",
							ValueFrom: &corev1.EnvVarSource{
								SecretKeyRef: &corev1.SecretKeySelector{
									LocalObjectReference: corev1.LocalObjectReference{
										Name: "my-secret",
									},
									Key: "GITKATAS_HOST",
								},
							},
						},
					},
				},
				{
					Name:  "gitkatas",
					Image: "daafonsecato/kataterm-gitkatas:v2",
					SecurityContext: &corev1.SecurityContext{
						RunAsUser:  func() *int64 { i := int64(1000); return &i }(),
						RunAsGroup: func() *int64 { i := int64(1000); return &i }(),
					},
					Ports: []corev1.ContainerPort{
						{
							ContainerPort: 7681,
						},
						{
							ContainerPort: 8080,
						},
						{
							ContainerPort: 8095,
						},
					},
					Env: []corev1.EnvVar{
						{
							Name:  "DB_HOST",
							Value: "postgres",
						},
					},
					VolumeMounts: []corev1.VolumeMount{
						{
							Name:      "gitkatas-exercise-volume",
							MountPath: "/home/git-katas-user/exercise",
						},
					},
				},
				{
					Name:  "validator",
					Image: "daafonsecato/kataterm-validator:v2",
					SecurityContext: &corev1.SecurityContext{
						RunAsUser:  func() *int64 { i := int64(1000); return &i }(),
						RunAsGroup: func() *int64 { i := int64(1000); return &i }(),
					},
					Ports: []corev1.ContainerPort{
						{
							ContainerPort: 8096,
						},
					},
					Env: []corev1.EnvVar{
						{
							Name:  "DB_HOST",
							Value: "postgres",
						},
					},
				},
			},
			Volumes: []corev1.Volume{
				{
					Name: "gitkatas-exercise-volume",
					VolumeSource: corev1.VolumeSource{
						EmptyDir: &corev1.EmptyDirVolumeSource{},
					},
				},
			},
		},
	}
	backendService := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: beSvcName,
		},
		Spec: corev1.ServiceSpec{
			Selector: map[string]string{
				"app": podName,
			},
			Ports: []corev1.ServicePort{
				{
					Name:       "port-8000",
					Protocol:   corev1.ProtocolTCP,
					Port:       8000,
					TargetPort: intstr.FromInt(8000),
				},
			},
		},
	}

	ttydService := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: ttydSvcName,
		},
		Spec: corev1.ServiceSpec{
			Selector: map[string]string{
				"app": podName,
			},
			Ports: []corev1.ServicePort{
				{
					Name:       "port-7681",
					Protocol:   corev1.ProtocolTCP,
					Port:       7681,
					TargetPort: intstr.FromInt(7681),
				},
			},
		},
	}

	codeEditorService := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: codeSvcName,
		},
		Spec: corev1.ServiceSpec{
			Selector: map[string]string{
				"app": podName,
			},
			Ports: []corev1.ServicePort{
				{
					Name:       "port-8080",
					Protocol:   corev1.ProtocolTCP,
					Port:       8080,
					TargetPort: intstr.FromInt(8080),
				},
			},
		},
	}

	// Define the GroupVersionResource (GVR)
	gvr := schema.GroupVersionResource{
		Group:    "traefik.containo.us",
		Version:  "v1alpha1",
		Resource: "ingressroutes",
	}

	ingressRoute := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "traefik.containo.us/v1alpha1",
			"kind":       "IngressRoute",
			"metadata": map[string]interface{}{
				"name":      ingressRouteName,
				"namespace": "default",
			},
			"spec": map[string]interface{}{
				"entryPoints": []string{"web"},
				"routes": []map[string]interface{}{
					{
						"match": beHostName,
						"kind":  "Rule",
						"middlewares": []map[string]interface{}{
							{
								"name":      "cors",
								"namespace": "default",
							},
						},
						"services": []map[string]interface{}{
							{
								"name": beSvcName,
								"port": 8000,
							},
						},
					},
					{
						"match": ttydHostName,
						"kind":  "Rule",
						"middlewares": []map[string]interface{}{
							{
								"name":      "cors",
								"namespace": "default",
							},
						},
						"services": []map[string]interface{}{
							{
								"name": ttydSvcName,
								"port": 7681,
							},
						},
					},
					{
						"match": codeeditorHostName,
						"kind":  "Rule",
						"middlewares": []map[string]interface{}{
							{
								"name":      "cors",
								"namespace": "default",
							},
						},
						"services": []map[string]interface{}{
							{
								"name": codeSvcName,
								"port": 8080,
							},
						},
					},
				},
			},
		},
	}

	// Create the IngressRoute
	_, err := controller.dynamiccclient.Resource(gvr).Namespace(controller.namespace).Create(context.TODO(), ingressRoute, metav1.CreateOptions{})
	if err != nil {
		panic(err.Error())
	}

	_, err = controller.clientset.CoreV1().Services(controller.namespace).Create(context.TODO(), backendService, metav1.CreateOptions{})
	if err != nil {
		fmt.Fprint(w, "Failed to create Kubernetes Backend service")
	}
	_, err = controller.clientset.CoreV1().Services(controller.namespace).Create(context.TODO(), ttydService, metav1.CreateOptions{})
	if err != nil {
		fmt.Fprint(w, "Failed to create Kubernetes TTYD service")
	}
	_, err = controller.clientset.CoreV1().Services(controller.namespace).Create(context.TODO(), codeEditorService, metav1.CreateOptions{})
	if err != nil {
		fmt.Fprint(w, "Failed to create Kubernetes CodeEditor service")
	}

	_, err = controller.clientset.CoreV1().Pods(controller.namespace).Create(context.TODO(), pod, metav1.CreateOptions{})
	if err != nil {
		fmt.Fprint(w, "Failed to create Kubernetes pod")
		return
	}
	// Store session ID and pod details in the database
	err = controller.sessionStore.StoreMachineAndSession(podName, PodUuid, PodUuid)
	if err != nil {
		fmt.Fprintf(w, "Failed to store session in the database: %s", err)
		return
	}
	newUUID := struct {
		PodUUID string `json:"newUUID"`
	}{
		PodUUID: PodUuid,
	}

	json.NewEncoder(w).Encode(newUUID)
}
