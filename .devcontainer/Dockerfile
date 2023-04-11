FROM mcr.microsoft.com/devcontainers/go:0-1-bullseye


ENV GOPROXY=https://goproxy.cn


RUN go install github.com/swaggo/swag/cmd/swag@latest


RUN git config --global --add safe.directory /workspaces/gopen