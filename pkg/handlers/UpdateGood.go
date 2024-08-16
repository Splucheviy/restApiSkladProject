package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"	
	"restApiSkladProject/pkg/models"

	"github.com/gorilla/mux"
)

// AddGood Add good
//
//	@Summary		Add good
//	@Description	Add good
//	@Tags			Goods
//	@Accept			json
//	@Produce		json
//	@Param			Id			path		string	true	"Goods id to update"
//	@Param			Name		body		string	true	"Goods name to update"
//	@Param			Provider	body		integer	true	"Goods provider to update"
//	@Param			Price		body		integer	true	"Goods price to update"
//	@Param			Quantity	body		integer	true	"Goods quantity to update"
//	@Success		200			{string}	string
//	@Router			/goods/{id} [put]
func (h handler) UpdateGoods(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    // Read request body
    defer r.Body.Close()
    body, err := io.ReadAll(r.Body)

    if err != nil {
        log.Fatalln(err)
    }

    var updatedGood models.Goods
    json.Unmarshal(body, &updatedGood)

    queryStmt := `UPDATE goodsDB SET title = $2, description = $3, content = $4 WHERE id = $1 RETURNING id;`
    err = h.DB.QueryRow(queryStmt, &id, &updatedGood.Id, &updatedGood.Name, &updatedGood.Provider, &updatedGood.Price, &updatedGood.Quantity).Scan(&id)
    if err != nil {
        log.Println("failed to execute query", err)
        w.WriteHeader(500)
        return
    }

    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode("Updated")

}
