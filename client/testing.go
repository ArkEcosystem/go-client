// This file is part of Ark Go Client.
//
// (c) Ark Ecosystem <info@ark.io>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"reflect"
	"strings"
	"testing"
)

const (
	// baseURLPath is a non-empty Client.BaseURL path to use during tests,
	// to ensure relative URLs are used for all endpoints. See issue #752.
	baseURLPath = "/api"
)

func setupTest() (client *Client, mux *http.ServeMux, serverURL string, teardown func()) {
	// mux is the HTTP request multiplexer used with the test server.
	mux = http.NewServeMux()

	// We want to ensure that tests catch mistakes where the endpoint URL is
	// specified as absolute rather than relative. It only makes a difference
	// when there's a non-empty base URL path. So, use that. See issue #752.
	apiHandler := http.NewServeMux()
	apiHandler.Handle(baseURLPath+"/", http.StripPrefix(baseURLPath, mux))
	apiHandler.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(os.Stderr, "FAIL: Client.BaseURL path prefix is not preserved in the request URL:")
		fmt.Fprintln(os.Stderr)
		fmt.Fprintln(os.Stderr, "\t"+req.URL.String())
		fmt.Fprintln(os.Stderr)
		http.Error(w, "Client.BaseURL path prefix is not preserved in the request URL.", http.StatusInternalServerError)
	})

	// server is a test HTTP server used to provide mock API responses.
	server := httptest.NewServer(apiHandler)

	// client is the Ark client being tested and is
	// configured to use test server.
	client = NewClient(nil)
	url, _ := url.Parse(server.URL + baseURLPath + "/")
	client.BaseURL = url

	return client, mux, server.URL, server.Close
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

func testGeneralError(t *testing.T, method string, err error) {
	if err != nil {
		t.Errorf("[%v] returned error: %v", method, err)
	}
}

func testResponseUrl(t *testing.T, method string, r *http.Response, want string) {
	if !strings.Contains(r.Request.URL.String(), want) {
		t.Errorf("[%+v][URL] got %+v, want %+v", method, r.Request.URL.String(), want)
	}
}

func testResponseStruct(t *testing.T, method string, got interface{}, want interface{}) {
	if !reflect.DeepEqual(got, want) {
		var gotType reflect.Type = reflect.TypeOf(got)
		var wantType reflect.Type = reflect.TypeOf(want)
		t.Errorf("[%+v][Response] got (%+v) %+v, want (%+v) %+v", method, gotType.String(), got, wantType.String(), want)
	}
}

type values map[string]int

type dummyJsonPayload struct {
	Limit int `json:"limit,omitempty"`
}

func testJsonPayload(t *testing.T, r *http.Request, values values) {
	want := make(map[string]int)
	for k, v := range values {
		want[k] = v
	}

	decoder := json.NewDecoder(r.Body)
	var got dummyJsonPayload
	_ = decoder.Decode(&got)

	gotBuffer, _ := json.Marshal(got)
	wantBuffer, _ := json.Marshal(want)

	if !reflect.DeepEqual(gotBuffer, wantBuffer) {
		t.Errorf("Request parameters: %v, want %v", string(gotBuffer), string(wantBuffer))
	}
}
