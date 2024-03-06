package wsession

import (
	"net/http"

	"github.com/gorilla/sessions"
)

type SessionData struct {
	UserId        int
	Authenticated bool
	DarkMode      bool
}

func Setup(session *sessions.Session, r *http.Request, w http.ResponseWriter) error {
	if !session.IsNew {
		return nil
	}
	session.Values["userID"] = -1
	session.Values["darkMode"] = false
	session.Values["authenticated"] = false
	err := session.Save(r, w)
	if err != nil {
		return err
	}
	return nil
}

func GetSessionData(session *sessions.Session) SessionData {
	return SessionData{
		UserId:        session.Values["userID"].(int),
		Authenticated: session.Values["authenticated"].(bool),
		DarkMode:      session.Values["darkMode"].(bool),
	}
}
