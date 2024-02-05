package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func (controller *SessionController) CreateKubernetesPodHandler(w http.ResponseWriter, r *http.Request) {

	PodUuid := uuid.New().String()
	podName := fmt.Sprintf("backend-pod-%s", PodUuid)
	beSvcName := fmt.Sprintf("backend-svc-%s", PodUuid)
	ttydSvcName := fmt.Sprintf("ttyd-svc-%s", PodUuid)
	codeSvcName := fmt.Sprintf("codeeditor-svc-%s", PodUuid)

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
							Name:  "NODE_ENV",
							Value: "development",
						},
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

	_, err := controller.clientset.CoreV1().Services(controller.namespace).Create(context.TODO(), backendService, metav1.CreateOptions{})
	if err != nil {
		fmt.Fprint(w, "Failed to create Kubernetes service")
	}

	_, err = controller.clientset.CoreV1().Services(controller.namespace).Create(context.TODO(), ttydService, metav1.CreateOptions{})
	if err != nil {
		fmt.Fprint(w, "Failed to create Kubernetes service")
	}
	_, err = controller.clientset.CoreV1().Services(controller.namespace).Create(context.TODO(), codeEditorService, metav1.CreateOptions{})
	if err != nil {
		fmt.Fprint(w, "Failed to create Kubernetes service")
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
	fmt.Fprint(w, PodUuid)
}
