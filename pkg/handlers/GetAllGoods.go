package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"restApiSkladProject/pkg/models"
)

// GetAllGoods Get all goods
//
//	@Summary		Get all goods
//	@Description	Get all goods
//	@Tags			Goods
//	@Accept			json
//	@Produce		json
//	@Success		200
//	@Router			/goods [get]
func (h handler) GetAllGoods(w http.ResponseWriter, r *http.Request) {

	results, err := h.DB.Query("SELECT * FROM goodsDB;")
	if err != nil {
		log.Println("failed to execute query", err)
		w.WriteHeader(500)
		return
	}

	var goods = make([]models.Goods, 0)
	for results.Next() {
		var good models.Goods
		err = results.Scan(&good.Id, &good.Name, &good.Provider, &good.Price, &good.Quantity)
		if err != nil {
			log.Println("failed to scan", err)
			w.WriteHeader(500)
			return
		}

		goods = append(goods, good)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(goods)
}
