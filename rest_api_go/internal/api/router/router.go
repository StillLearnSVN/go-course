package router

import (
	"net/http"
)

func MainRouter() *http.ServeMux {

	tRouter := teachersRouter()
	sRouter := studentsRouter()

	sRouter.Handle("/", ExecsRouter())
	tRouter.Handle("/", sRouter)
	return tRouter
}
