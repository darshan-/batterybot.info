package goapp

import (
	"net/http"
)

const faqURL = "http://code.google.com/p/battery-indicator/wiki/FAQ?tm=6"
const changelogURL = "http://code.google.com/p/battery-indicator/wiki/Changelog?tm=6"
const sourceURL = "http://code.google.com/p/battery-indicator/"

func init() {
	http.HandleFunc("/faq", makeRedirect(faqURL))
	http.HandleFunc("/changelog", makeRedirect(changelogURL))
	http.HandleFunc("/sourcecode", makeRedirect(sourceURL))
}

func makeRedirect(url string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, url, http.StatusFound)
	}
}
