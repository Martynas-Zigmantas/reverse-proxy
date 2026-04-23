package main

import (
    "crypto/tls"
    "log"
    "net/http"
    "net/http/httputil"
    "net/url"
)

func main() {
    // Define backends ONCE
    cloudBackend, _ := url.Parse("http://127.0.0.1:35547")
    vaultBackend, _ := url.Parse("http://127.0.0.1:55577")
    metabaseBackend, _ := url.Parse("http://127.0.0.1:60123")

    // Create proxies ONCE
    cloudProxy := httputil.NewSingleHostReverseProxy(cloudBackend)
    vaultProxy := httputil.NewSingleHostReverseProxy(vaultBackend)
    metabaseProxy := httputil.NewSingleHostReverseProxy(metabaseBackend)

    // Server
    server := &http.Server{
        Addr: ":443",
        Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

            switch r.Host {
            case "cloud.jack-sally.com":
                cloudProxy.ServeHTTP(w, r)

            case "vault.jack-sally.com":
                vaultProxy.ServeHTTP(w, r)

            case "dash.jack-sally.com":
                metabaseProxy.ServeHTTP(w, r)

            default:
                http.Error(w, "Forbidden", http.StatusForbidden)
            }
        }),
        TLSConfig: &tls.Config{
            MinVersion: tls.VersionTLS12,
        },
    }

    log.Println("Starting HTTPS proxy on :443")
    err := server.ListenAndServeTLS(
        "/etc/letsencrypt/live/cloud.jack-sally.com/fullchain.pem",
        "/etc/letsencrypt/live/cloud.jack-sally.com/privkey.pem",
    )
    if err != nil {
        log.Fatal(err)
    }
}