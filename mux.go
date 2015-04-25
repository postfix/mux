package mux

import "github.com/go-ndn/ndn"

type Mux struct {
	*Router
	mw Handler
}

func New() *Mux {
	r := NewRouter()
	return &Mux{Router: r, mw: r}
}

func (mux *Mux) Use(m Middleware) {
	mux.mw = m(mux.mw)
}

func (mux *Mux) Run(w Sender, ch <-chan *ndn.Interest) {
	for i := range ch {
		go mux.mw.ServeNDN(w, i)
	}
}