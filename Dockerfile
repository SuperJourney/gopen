FROM golang:1-bullseye as build
WORKDIR /www
EXPOSE 8080

ARG GOPROXY
ENV GOPROXY=${GOPROXY} 

COPY go.mod /www
RUN  go mod download

COPY . /www


FROM build as init

ENV CGO_ENABLED=1
RUN go build 

FROM build as deploy
ENV CGO_ENABLED=0
RUN go install github.com/swaggo/swag/cmd/swag@latest && swag init
RUN  go build -o server

FROM alpine:3.12

COPY --from=deploy /www/server /bin
CMD ["/bin/server"]