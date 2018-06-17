package main

import ( "log"
	"net/http"
	"flag"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	staticDir := flag.String("static-dir", "./ui/static", "Path to access static files")

	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", Home)
	mux.HandleFunc("/snippet", ShowSnippet)
	mux.HandleFunc("/snippet/new", NewSnippet)

	fileServer := http.FileServer(http.Dir(*staticDir))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Printf("Starting server on %s", *addr)

	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}