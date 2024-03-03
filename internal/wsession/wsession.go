package wsession

import (
	"net/http"

	"github.com/gorilla/sessions"
)

type SessionData struct {
	DarkMode bool
}

func Setup(session *sessions.Session, r *http.Request, w http.ResponseWriter) error {
	if !session.IsNew {
		return nil
	}
	session.Values["darkMode"] = false
	err := session.Save(r, w)
	if err != nil {
		return err
	}
	return nil
}

func GetSessionData(session *sessions.Session) SessionData {
	return SessionData{
		DarkMode: session.Values["darkMode"].(bool),
	}
}
