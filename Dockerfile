
# for CI testing
ARG GO_VERSION=1.9
ARG PROJECT_PATH=/go/src/github.com/dustin-decker/featuremill
FROM golang:${GO_VERSION}-alpine AS builder