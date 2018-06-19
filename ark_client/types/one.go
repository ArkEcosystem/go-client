package types

import (
    "context"
    "net/http"
)

type One_AccountsService interface {
    List(context.Context) (*http.Response, error)
    Get(context.Context) (*http.Response, error)
    Count(context.Context) (*http.Response, error)
    Delegate(context.Context) (*http.Response, error)
    DelegateFee(context.Context) (*http.Response, error)
    Balance(context.Context) (*http.Response, error)
    PublicKey(context.Context) (*http.Response, error)
    Top(context.Context) (*http.Response, error)
}
