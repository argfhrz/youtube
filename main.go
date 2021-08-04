package main

import (
	"log"
	"net/http"
	"youtube/apis"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Connected"))
	}).Methods("GET")

	dataApi := apis.DataApi{}
	router.HandleFunc("/data", dataApi.PostData).Methods("POST")
	router.HandleFunc("/data/update", dataApi.UpdateData).Methods("PUT")
	router.HandleFunc("/datas", dataApi.GetData).Methods("GET")
	router.HandleFunc("/data/find", dataApi.Find).Methods("GET")
	router.HandleFunc("/data/delete", dataApi.RemoveDataByID).Methods("DELETE")

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	log.Println("Server listening to 127.0.0.1:8001")
	err := http.ListenAndServe("127.0.0.1:8001", handlers.CORS(headers, methods, origins)(router))
	if err != nil {
		log.Fatal(err)
	}
}
