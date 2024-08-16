package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"restApiSkladProject/pkg/models"

	"github.com/google/uuid"
)

// AddGood Add good
//
//	@Summary		Add good
//	@Description	Add good
//	@Tags			Goods
//	@Accept			json
//	@Produce		json
//	@Param			data	body		models.Goods	true	"Goods data"
//	@Success		200		{string}	string
//	@Router			/goods [post]
func (h handler) AddGood(w http.ResponseWriter, r *http.Request) {
	// Read to request body
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
		w.WriteHeader(500)
		return
	}
	var good models.Goods
	json.Unmarshal(body, &good)

	good.Id = (uuid.New()).String()
	queryStmt := `INSERT INTO goodsDB (id,name,provider,price,quantity) VALUES ($1, $2, $3, $4, $5) RETURNING id;`
	err = h.DB.QueryRow(queryStmt, &good.Id, &good.Name, &good.Provider, &good.Price, &good.Quantity).Scan(&good.Id)
	if err != nil {
		log.Println("failed to execute query", err)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")

}
