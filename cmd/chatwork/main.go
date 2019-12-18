package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/pkg/errors"
)

func main() {
	// Ref: http://developer.chatwork.com/ja/endpoint_rooms.html#POST-rooms-room_id-messages

	// data
	data := url.Values{}
	body := os.Getenv("CHATWORK_MESSAGE_BODY")
	data.Set("body", body)

	selfUnread := "0"
	if os.Getenv("CHATWORK_SELF_UNREAD") == "true" {
		selfUnread = "1"
	}
	data.Set("self_unread", selfUnread)

	// url
	roomID := os.Getenv("CHATWORK_ROOM_ID")
	url := fmt.Sprintf("https://api.chatwork.com/v2/rooms/%s/messages", roomID)

	// request
	req, err := http.NewRequest("POST", url, strings.NewReader(data.Encode()))
	if err != nil {
		log.Fatal(errors.Wrap(err, "request"))
		return
	}
	token := os.Getenv("CHATWORK_TOKEN")
	req.Header.Add("X-ChatWorkToken", token)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Send message
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(errors.Wrap(err, "request is failed"))
		return
	}
	res.Body.Close()

}
