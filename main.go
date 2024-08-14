package main

import (
	"fmt"
	"net/http"
	// "net/url"
)

type apiConfig struct {
	fileserverHits int
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits++
		fmt.Println(cfg.fileserverHits)
		next.ServeHTTP(w, r)
	})

}

func Check(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain;charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}


func (cfg *apiConfig) Metrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain;charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hits: %d", cfg.fileserverHits)
}


func (cfg *apiConfig) Reset(w http.ResponseWriter, r *http.Request) {
	cfg.fileserverHits = 0
	w.Header().Set("Content-Type", "text/plain;charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func main() {
	// newApp()
	mux := http.NewServeMux()
	apiCfg := apiConfig{}

	// wrapMux := apiConfig.middlewareMetricsInc(mux)

	serv := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	mux.HandleFunc("/healthz", Check)
	mux.HandleFunc("/metrics", apiCfg.Metrics)
	mux.HandleFunc("/reset", apiCfg.Reset)
	fs := http.FileServer(http.Dir("./assets"))

	mux.Handle("/app/", apiCfg.middlewareMetricsInc(http.StripPrefix("/app", fs)))
	serv.ListenAndServe()
}
