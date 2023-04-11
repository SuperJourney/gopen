FROM golang:1-bullseye

WORKDIR /www

# RUN --mount=type=ssh  mkdir ~/.ssh &&  ssh-keyscan github.com >> ~/.ssh/known_hosts && git clone git@github.com:SuperJourney/CircleCI-Test.git
EXPOSE 8080


COPY go.mod /www
RUN  go mod download

COPY .  /www

RUN  GOOS=linux GOARCH=amd64 go build -o app

RUN  GOOS=linux GOARCH=amd64 go install 

ENTRYPOINT ["gopen"]
 