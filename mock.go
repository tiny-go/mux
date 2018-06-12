package mux

import "net/http"

var _ Node = (*Mock)(nil)

type Mock struct{}

func (*Mock) Add(string, http.Handler) error { return nil }

func (*Mock) Get(string) (http.Handler, map[string]string, bool) { return nil, nil, false }
