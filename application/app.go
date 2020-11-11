package application

import (
	"func-api/application/common"
	"net/http"
	_ "net/http/pprof"
)

func Application(dep common.Dependency) (err error) {
	cfg := dep.Config
	if cfg.Debug != "" {
		go http.ListenAndServe(cfg.Debug, nil)
	}
	return
}
