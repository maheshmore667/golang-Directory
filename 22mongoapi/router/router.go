package router

import (
	"github.com/gorilla/mux"
	"github.com/maheshmore667/mongoapi/controllers"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/movie", controllers.GetPlayList).Methods("GET")
	router.HandleFunc("/api/movie", controllers.CreateOneMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controllers.MarkWatched).Methods("PUT")
	router.HandleFunc("/api/movie/deleteMovie/{id}", controllers.DeleteOneMovie).Methods("DELETE")
	router.HandleFunc("/api/movie/deletePlaylist", controllers.DeleteAllMovie).Methods("DELETE")
	return router
}
