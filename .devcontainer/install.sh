#!/bin/bash

# npm installs ---------------------------------------------------
npm install -g @go-task/cli

# pip installs ---------------------------------------------------
pip install -r .devcontainer/requirements.txt

# go installs ----------------------------------------------------
go install github.com/charmbracelet/glow@latest

# kubectl installs -----------------------------------------------
curl -L -o /tmp/kubectl "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
curl -L -o /tmp/kubectl.sha256 "https://dl.k8s.io/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl.sha256"
echo "$(cat /tmp/kubectl.sha256)  /tmp/kubectl" | sha256sum --check
sudo install -o root -g root -m 0755 /tmp/kubectl /usr/local/bin/kubectl
