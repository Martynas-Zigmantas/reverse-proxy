package main

import (
    "net/http"
)

func main() {
    http.ListenAndServe(":80", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        switch r.Host {
        case "clouds.jack-sally.com", "vault.jack-sally.com", "dash.jack-sally.com":
            http.Redirect(w, r, "https://"+r.Host+r.RequestURI, http.StatusMovedPermanently)

        default:
            http.Error(w, "Forbidden", http.StatusForbidden)
        }
    }))
}