package client

import (
	"context"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func testQueryParam(t *testing.T, got url.Values, key string, want string) {
	t.Helper()

	if v := got.Get(key); v != want {
		t.Errorf("query param %q = %q, want %q", key, v, want)
	}
}

func testQueryParamAbsent(t *testing.T, got url.Values, key string) {
	t.Helper()

	if got.Has(key) {
		t.Errorf("query param %q = %q, want absent", key, got.Get(key))
	}
}

func TestQueryEncoding_BlocksQuery(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	var got url.Values
	mux.HandleFunc("/blocks", func(w http.ResponseWriter, r *http.Request) {
		got = r.URL.Query()
		w.Write([]byte(`{"data":[]}`))
	})

	query := &BlocksQuery{
		Hash:       "abc123",
		Number:     10,
		NumberFrom: 5,
		NumberTo:   15,
	}
	_, _, err := client.Blocks.List(context.Background(), query)
	testGeneralError(t, "Blocks.List", err)

	testQueryParam(t, got, "hash", "abc123")
	testQueryParam(t, got, "number", "10")
	testQueryParam(t, got, "number.from", "5")
	testQueryParam(t, got, "number.to", "15")
	testQueryParamAbsent(t, got, "height")
	testQueryParamAbsent(t, got, "height.from")
	testQueryParamAbsent(t, got, "height.to")
	testQueryParamAbsent(t, got, "id")
}

func TestQueryEncoding_PaginationDefaults(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	var got url.Values
	mux.HandleFunc("/wallets", func(w http.ResponseWriter, r *http.Request) {
		got = r.URL.Query()
		w.Write([]byte(`{"data":[]}`))
	})

	_, _, err := client.Wallets.List(context.Background(), &WalletsQuery{})
	testGeneralError(t, "Wallets.List", err)
	testQueryParam(t, got, "page", "1")

	var gotNested url.Values
	mux.HandleFunc("/tokens/transfers", func(w http.ResponseWriter, r *http.Request) {
		gotNested = r.URL.Query()
		w.Write([]byte(`{"data":[]}`))
	})

	_, _, err = client.Tokens.Transfers(context.Background(), &TokenTransfersQuery{})
	testGeneralError(t, "Tokens.Transfers", err)
	testQueryParam(t, gotNested, "page", "1")
}

func TestQueryEncoding_CommaSeparated(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	var got url.Values
	var raw string
	mux.HandleFunc("/wallets/tokens", func(w http.ResponseWriter, r *http.Request) {
		got = r.URL.Query()
		raw = r.URL.RawQuery
		w.Write([]byte(`{"data":[]}`))
	})

	query := &WalletTokensQuery{
		Addresses: CommaSeparated{"0xabc", "0xdef", "0x123"},
	}
	_, _, err := client.Wallets.Tokens(context.Background(), query)
	testGeneralError(t, "Wallets.Tokens", err)

	testQueryParam(t, got, "addresses", "0xabc,0xdef,0x123")
	if strings.Count(raw, "addresses=") != 1 {
		t.Errorf("expected a single addresses= param, got raw query %q", raw)
	}
}

func TestQueryEncoding_WeiScaleStringFields(t *testing.T) {
	client, mux, _, teardown := setupTest()
	defer teardown()

	hugeBalance := "934028115287775973804882" // exceeds int64 max (~9.2e18) by ~5 orders of magnitude

	var got url.Values
	mux.HandleFunc("/wallets", func(w http.ResponseWriter, r *http.Request) {
		got = r.URL.Query()
		w.Write([]byte(`{"data":[]}`))
	})

	query := &WalletsQuery{Balance: hugeBalance}
	_, _, err := client.Wallets.List(context.Background(), query)
	testGeneralError(t, "Wallets.List", err)
	testQueryParam(t, got, "balance", hugeBalance)

	var gotVal url.Values
	mux.HandleFunc("/validators", func(w http.ResponseWriter, r *http.Request) {
		gotVal = r.URL.Query()
		w.Write([]byte(`{"data":[]}`))
	})

	valQuery := &ValidatorsQuery{ForgedFeesFrom: hugeBalance, Votes: hugeBalance}
	_, _, err = client.Validators.List(context.Background(), valQuery)
	testGeneralError(t, "Validators.List", err)
	testQueryParam(t, gotVal, "forged.fees.from", hugeBalance)
	testQueryParam(t, gotVal, "votes", hugeBalance)
}
