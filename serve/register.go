package serve

import "net/http"

func Register()  {
	http.HandleFunc("/", Template)
	http.HandleFunc("/api", api)
}


