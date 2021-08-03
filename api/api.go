package api

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/idasilva/npe/internal/pkg"
	"net/http"
	"os"
	"os/signal"
	"sync"
)

type Digital struct {
	Context context.Context
	Router *mux.Router
	Sample  *Sample
	Config  pkg.CfgFile
}

func (d *Digital)Server() error{
	ctx, cancel := context.WithCancel(d.Context)

	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, os.Interrupt)
		<-ch
		cancel()
		d.Sample.Shutdown(ctx)
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer cancel()
		d.Sample.Start(d)
	}()
	wg.Wait()

	return nil
}

func(d *Digital) ServeHTTP(w http.ResponseWriter, r *http.Request){
	q := r.URL.Query()
	q.Add("env", d.Config.Env)
	r.URL.RawQuery = q.Encode()

	fmt.Println("[server http]", d.Config.Env)

	var h http.Handler = d.Router
	h.ServeHTTP(w,r)
}

func NewDigital(config pkg.CfgFile) *Digital{
	fmt.Println("[new digital]",config.Env)
	return &Digital{
		context.Background(),
		mux.NewRouter().StrictSlash(true),
		NewSample(),
		config,
	}
}


