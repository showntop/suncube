package routes

import (
	"github.com/julienschmidt/httprouter"

	"github.com/showntop/suncube/app/handlers"
)

func Instrument() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", handlers.IndexHomeHandler)
	router.GET("/videos", handlers.IndexVideoHandler)
	router.GET("/videos/:id", handlers.ShowVideoHandler)
	return router
}
