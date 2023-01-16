FROM golang:1.16-buster as build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY ./ ./

RUN go build -o /server_executable ./command_line/*.go

FROM docker:dind

ARG USER=app_user
ARG PASS="qazxsw123$"


RUN adduser -D $USER && echo "$USER:$PASS" | chpasswd
ARG DOCKER_GID=992

RUN addgroup -g $DOCKER_GID docker

RUN addgroup $USER docker

USER app_user

WORKDIR /home/app_user 

COPY --from=build /server_executable ./server_executable

EXPOSE 8376

ENTRYPOINT ["./server_executable"]