package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"restApiSkladProject/pkg/models"

	"github.com/gorilla/mux"
)

// GetGood Get goods by id
//
//	@Summary		Get goods by id
//	@Description	Get goods by id
//	@Tags			Goods
//	@Accept			json
//	@Produce		json
//	@Param			Id	path		string	true	"Goods id"
//	@Success		200	{string}	string
//	@Router			/goods/{Id} [get]
func (h handler) GetGood(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	queryStmt := `SELECT * FROM goodsDB WHERE id = $1 ;`
	results, err := h.DB.Query(queryStmt, id)
	if err != nil {
		log.Println("failed to execute query", err)
		w.WriteHeader(500)
		return
	}

	var good models.Goods
	for results.Next() {
		err = results.Scan(&good.Id, &good.Name, &good.Provider, &good.Price, &good.Quantity)
		if err != nil {
			log.Println("failed to scan", err)
			w.WriteHeader(500)
			return
		}
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(good)
}
