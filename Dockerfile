FROM golang:1-bullseye as build
WORKDIR /www

ARG GOPROXY
ENV GOPROXY=${GOPROXY} 

COPY go.mod /www
RUN  go mod download

COPY . /www
ENV CGO_ENABLED=0
RUN go install github.com/swaggo/swag/cmd/swag@latest && swag init
RUN  go run cmd/main.go  migration init --dbfile /data/db
RUN  go build -o server


FROM alpine:3.12
EXPOSE 8080

COPY --from=build /www/server /bin
CMD ["/bin/server"]