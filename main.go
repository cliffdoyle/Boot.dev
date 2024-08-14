package main

import "net/http"

func Check(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","text/plain;charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func main() {
	mux := http.NewServeMux()

	serv := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	mux.HandleFunc("/healthz",Check)
	fs:=http.FileServer(http.Dir("./assets"))
	mux.Handle("/app/",http.StripPrefix("/app",fs))

	serv.ListenAndServe()
}