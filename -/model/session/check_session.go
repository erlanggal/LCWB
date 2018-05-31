package session

import (
	"github.com/gorilla/sessions"
	"net/http"
)

func CheckSession(r *http.Request) interface{} {
	var store = sessions.NewCookieStore([]byte("%SESSION%2104%"))
	sessionToken, _ := store.Get(r, "session-token")
	return sessionToken.Values["token"]
}
