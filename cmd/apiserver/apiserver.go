package apiserver

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"restApiSkladProject/pkg/handlers"

	"github.com/gorilla/mux"

	// swaggerFiles "github.com/swaggo/files"
	httpSwagger "github.com/swaggo/http-swagger"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Sklad REST API!")
	fmt.Println("SKLAD REST API")
}

func HandleRequests(DB *sql.DB) {
	h := handlers.New(DB)
	r := mux.NewRouter().StrictSlash(true)
	
	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)

	r.HandleFunc("/", homePage)
	r.HandleFunc("/goods", h.GetAllGoods).Methods(http.MethodGet)
	r.HandleFunc("/goods/{id}", h.GetGood).Methods(http.MethodGet)
	r.HandleFunc("/goods", h.AddGood).Methods(http.MethodPost)
	r.HandleFunc("/goods/{id}", h.UpdateGoods).Methods(http.MethodPut)
	r.HandleFunc("/goods/{id}", h.DeleteGoods).Methods(http.MethodDelete)
	log.Fatal(http.ListenAndServe(":8080", r))
}