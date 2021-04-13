package faaslib_go_handler

import (
	"bytes"
	"encoding/json"
	"github.com/kpango/glg"
	"github.com/rs/xid"
	"net/http"
	"time"
)

var url = "https://postman-echo.com/post"

func InjectCallback(key string, start time.Time, xid xid.ID) {
	timeTaken := time.Since(start)
	callbackHomeCompleteEvent(key, timeTaken, xid)
}

func CallHomeStart(key string) xid.ID {
	uniqueToken := xid.New() //Token for starting + ending, to link the 'Events'
	go callbackHomeStartEvent(key, uniqueToken, "STARTED")
	_ = glg.Infof("calling home start")
	return uniqueToken
}

func callbackHomeStartEvent(key string, xid xid.ID, event string) {
	postBody, _ := json.Marshal(map[string]string{
		"key":   key,
		"token": xid.String(),
		"event": event,
	})

	http.Post(url, "application/json", bytes.NewBuffer(postBody))
}

func callbackHomeCompleteEvent(key string, timeTaken time.Duration, xid xid.ID) {
	_ = glg.Infof("key: %s took %s for token %s", key, timeTaken, xid.String())
}
