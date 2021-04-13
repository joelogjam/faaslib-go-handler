package faaslib_go_handler

import (
	"bytes"
	"encoding/json"
	"github.com/kpango/glg"
	"github.com/rs/xid"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

var url = "https://postman-echo.com/post"

func InjectCallback(programName string, version string, start time.Time, xid xid.ID) {
	timeTaken := time.Since(start)
	callbackHome(xid, "COMPLETED")
	_ = glg.Infof("exe: %s v %s took %s for token %s", programName, version, timeTaken, xid.String())
}

func CallHomeStart() xid.ID {
	uniqueToken := xid.New()
	go callbackHome(uniqueToken, "STARTED")
	_ = glg.Infof("calling home start")
	return uniqueToken
}

func callbackHome(xid xid.ID, event string) {
	postBody, _ := json.Marshal(map[string]string{
		"name":  filepath.Base(os.Args[0]),
		"token": xid.String(),
		"event": event,
	})

	http.Post(url, "application/json", bytes.NewBuffer(postBody))
}
