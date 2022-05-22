package utils

import (
	"admin/models"
	"net/http"
	"sync"
	"time"
)

const (
	cookieName    = "go-session"
	cookieExpires = 12 * time.Hour
)

var Sessions = struct {
	context map[string]interface{}
	m       map[string]models.User
	sync.RWMutex
}{m: make(map[string]models.User), context: make(map[string]interface{})}

func SetContext(r *http.Request, name string, context interface{}) {
	Sessions.Lock()
	defer Sessions.Unlock()
	uuid := getValCookie(r)
	if _, ok := Sessions.m[uuid]; ok {
		Sessions.m[uuid].Context[name] = context
	}
}

func GetContext(r *http.Request, name string) interface{} {
	usr := GetUser(r)
	return usr.Context[name]
}

func GetFullContext(r *http.Request) map[string]interface{} {
	usr := GetUser(r)
	return usr.Context
}

func SetSession(user models.User, w http.ResponseWriter) {
	Sessions.Lock()
	defer Sessions.Unlock()

	uuid := GenerateJWT(user)
	user.Context = make(map[string]interface{})

	Sessions.m[uuid] = user

	cookie := &http.Cookie{
		Name:    cookieName,
		Value:   uuid,
		Path:    "/",
		Expires: time.Now().Add(cookieExpires),
	}
	// Sessions.m[uuid].Context = make(map[string]interface{})
	Sessions.m[uuid].Context["User"] = user
	t := time.Now()
	Sessions.m[uuid].Context["Today"] = t.Format("02-01-2006")
	Sessions.m[uuid].Context["Error"] = nil

	http.SetCookie(w, cookie)
}

func GetUser(r *http.Request) models.User {
	Sessions.Lock()
	defer Sessions.Unlock()
	uuid := getValCookie(r)
	if user, ok := Sessions.m[uuid]; ok {
		return user
	}
	return models.User{}
}

func DeleteSession(w http.ResponseWriter, r *http.Request) {
	Sessions.Lock()
	defer Sessions.Unlock()

	delete(Sessions.m, getValCookie(r))

	cookie := &http.Cookie{
		Name:   cookieName,
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
}

func getValCookie(r *http.Request) string {
	if cookie, err := r.Cookie(cookieName); err == nil {
		return cookie.Value
	}
	return ""
}

func IsAuthenticated(r *http.Request) bool {
	return getValCookie(r) != ""
}

func IsAdmin(r *http.Request) bool {
	user := GetUser(r)
	if user.Role == "admin" {
		return true
	} else {
		return false
	}
}
