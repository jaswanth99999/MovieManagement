package main

import(
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jaswanth99999/MovieManagement/pkg/routes"
)

func main(){
	r := mux.NewRouter()
	routes.RegisterMoviesRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9090",r))
}

