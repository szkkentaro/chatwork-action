FROM golang:1.13-alpine

RUN go get github.com/szkkentaro/chatwork-action/cmd/chatwork

LABEL "com.github.actions.name"="Actions for Chatwork"
LABEL "com.github.actions.description"="Send a message to Chatwork"
LABEL "com.github.actions.icon"="message-circle"
LABEL "com.github.actions.color"="white"

ENTRYPOINT ["chatwork"]
