package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// AddGood Add good
//
//	@Summary		Add good
//	@Description	Add good
//	@Tags			Goods
//	@Accept			json
//	@Produce		json
//	@Param			Id	path		string	true	"Goods Id for delete"
//	@Success		200	{string}	string
//	@Router			/goods/{id} [delete]
func (h handler) DeleteGoods(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	queryStmt := `DELETE FROM goodsDB WHERE id = $1;`
	_, err := h.DB.Query(queryStmt, &id)
	if err != nil {
		log.Println("failed to execute query", err)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deleted")
}
