package crud

import (
	"crud/controller"
	"crud/modelResource"
	"crud/serviceGlobals"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	//	next line just generates some test data records:
	modelResource.Init()

	// next define the endpoints and the functions to be called
	// this fuctionality is the mux (multiplexer)
	// there are mux frameworks (e.g. Gorilla).
	// Usually they are not necessary for microservices
	// However, they are very useful for multiple endpoint APIs
	//
	http.HandleFunc("/resources", controller.ResourcesHandleFunc)
	http.HandleFunc("/resources/", controller.ResourceHandleFunc)
	http.HandleFunc("/", controller.Home)
	http.HandleFunc("/about", about)
}

	func StartWebserver() {

		controller.SetRoutesResource()

		// configuration parameters for webserver
		s := &http.Server{
			Addr:           serviceGlobals.SvcGlob.Port,
			Handler:        nil,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}
		fmt.Println(
			"\nWebserver is started and listening on port",
			serviceGlobals.SvcGlob.Port,
			"\n",
		)
		log.Fatal(s.ListenAndServe())
	}


func about(w http.ResponseWriter, r *http.Request) {
	// response for version information
	fmt.Fprint(w, "KWer's CRUD-REST-API for resources: v0.0.7, 08/06/2019 11:56 MEST.")
}
