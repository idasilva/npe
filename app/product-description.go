package app

import (
	"fmt"
	"github.com/idasilva/npe/api"
	"github.com/idasilva/npe/internal/pkg"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
    fmt.Println("handler....", fmt.Sprintf(" env: %v",r.URL.Query().Get("env")))
}

func Run(config pkg.CfgFile) (*api.Digital, error) {
	fmt.Println("[run]",config.Env)

	api := api.NewDigital(config)
	api.Router.HandleFunc("/product", Handler).Methods(http.MethodGet)
	return api, nil
}
