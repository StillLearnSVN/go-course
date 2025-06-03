package router

import (
	"net/http"
)

func MainRouter() *http.ServeMux {

	tRouter := TeacherRouter()
	sRouter := StudentRouter()

	sRouter.Handle("/", ExecsRouter())
	tRouter.Handle("/", sRouter)
	return tRouter
}
