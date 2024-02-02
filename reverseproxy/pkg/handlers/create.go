package handlers

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"github.com/google/uuid"
)

func (controller *SessionController) CreateKubernetesPodHandler(w http.ResponseWriter, r *http.Request) {
	config, err := rest.InClusterConfig()
	if err != nil {
		fmt.Fprint(w, "Failed to create Kubernetes pod")
		return
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Fprint(w, "Failed to create Kubernetes pod")
		return
	}
	
	podName := fmt.Sprintf("backend-pod-%s", uuid.New().String())

	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: podName,
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
	_, err = clientset.CoreV1().Pods("default").Create(context.TODO(), pod, metav1.CreateOptions{})
	if err != nil {
		fmt.Fprint(w, "Failed to create Kubernetes pod")
		return
	}

	fmt.Fprint(w, "Kubernetes pod created successfully")
}
