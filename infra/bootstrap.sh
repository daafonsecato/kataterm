#!/bin/bash

# Install Docker
sudo apt-get update
sudo apt-get install -y docker.io

# Add ubuntu user to docker group
sudo usermod -aG docker ubuntu

# Install Containerd
sudo apt-get install -y containerd

# Install Kubernetes
sudo apt-get update && sudo apt-get install -y apt-transport-https curl
curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add -
echo "deb https://apt.kubernetes.io/ kubernetes-xenial main" | sudo tee /etc/apt/sources.list.d/kubernetes.list
sudo apt-get update
sudo apt-get install -y kubelet kubeadm kubectl

# Initialize Kubernetes cluster
sudo kubeadm init --apiserver-advertise-address=10.0.0.37 --pod-network-cidr=10.32.0.0/12

# Set up kubeconfig for non-root user
mkdir -p $HOME/.kube
sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config

# Install network add-on for pod networking
kubectl apply -f https://github.com/weaveworks/weave/releases/download/v2.8.1/weave-daemonset-k8s.yaml

# Allow scheduling pods on the master node (optional)
kubectl taint nodes --all node-role.kubernetes.io/master-

# Print join command for worker nodes
echo "To join worker nodes to this cluster, run the following command:"
sudo kubeadm token create --print-join-command


sudo modprobe br_netfilter
sudo sysctl net.bridge.bridge-nf-call-iptables=1
echo "net.bridge.bridge-nf-call-iptables=1" | sudo tee -a /etc/sysctl.conf
sudo sysctl net.ipv4.ip_forward=1
echo "net.ipv4.ip_forward=1" | sudo tee -a /etc/sysctl.conf
sudo sysctl -p
