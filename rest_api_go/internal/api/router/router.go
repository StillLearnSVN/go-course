package router

import (
	"net/http"
)

func MainRouter() *http.ServeMux {

	tRouter := TeacherRouter()
	sRouter := StudentRouter()

	tRouter.Handle("/", sRouter)
	return tRouter
	// mux := http.NewServeMux()
	// mux.HandleFunc("/", handlers.RootHandler)
	// mux.HandleFunc("GET /execs/", handlers.ExecsHandler)
	// return mux
}
