package net_http

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

type MyCookie struct {
	Cookie http.Cookie
	Expire time.Time
}

func (mc MyCookie) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// method 1: use http.SetCookie
	if mc.Expire.String() == "0001-01-01 00:00:00 +0000 UTC" {
		http.SetCookie(w, &mc.Cookie)
		w.Write([]byte("Cookie demo :" + mc.Expire.String()))
	} else {
		// method 2: append the formatted cookie string to the header of response writer
		// You should set the date format to unix date, otherwise the Browser will not recognized
		// the default time format '2023-10-16T14:13:45.000Z'
		expire_date := mc.Expire.UTC().Format(time.UnixDate)
		cookie_string := fmt.Sprintf("site=example.com; Expires=%s", expire_date)
		w.Header().Add("Set-Cookie", cookie_string)
		w.Write([]byte("Cookie demo :" + cookie_string))
	}
}

func Test_Set_Cookie(t *testing.T) {
	mycookie := MyCookie{
		Cookie: http.Cookie{
			Name:    "site",
			Value:   "example.com",
			Expires: time.Now().AddDate(0, 0, 21)}, // For method 1
		Expire: time.Now().AddDate(0, 0, 21), // For method 2
	}

	http.Handle("/set", &mycookie)
	t.Fatal(http.ListenAndServe(":8080", nil))
}

func Test_Date_Nil(t *testing.T) {
	var nil_date time.Time
	t.Logf("nil date: %v", nil_date.Format(time.UnixDate))
}
