#!/bin/bash

# npm installs ---------------------------------------------------
npm install -g @go-task/cli

# pip installs ---------------------------------------------------
pip install -r .devcontainer/requirements.txt

# go installs ----------------------------------------------------
go install github.com/charmbracelet/glow@latest
go install github.com/goreleaser/goreleaser@latest
