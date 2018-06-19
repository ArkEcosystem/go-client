package types

import (
    "net/http"
    "net/url"
    "sync"
)

type Service struct {
    Client *HttpClient
}

type HttpClient struct {
    clientMu sync.Mutex
    Client   *http.Client

    BaseURL *url.URL
}
