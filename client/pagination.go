package client

import (
	"net/url"
	"strings"
)

type Pagination struct {
	Page   int `url:"page"`
	Limit  int `url:"limit"`
	Offset int `url:"offset,omitempty"`
}

type paginator interface {
	applyDefaults()
}

func (p *Pagination) applyDefaults() {
	if p == nil {
		return
	}

	if p.Page == 0 {
		p.Page = 1
	}
}

type CommaSeparated []string

func (c CommaSeparated) EncodeValues(key string, v *url.Values) error {
	if len(c) == 0 {
		return nil
	}

	v.Set(key, strings.Join(c, ","))

	return nil
}
