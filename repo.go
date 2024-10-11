package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Item struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Price      float32 `json:"price"`
	Quantity   int     `json:"quantity"`
	Category   string  `json:"category"`
	Created_at string  `json:"created_at"`
}

func get(w http.ResponseWriter, r *http.Request) {
	ItemName := chi.URLParam(r, "ItemName")

	var item Item
	if err := DB.Get(&item, "SELECT id, name, price, quantity, category, created_at FROM public.items WHERE name = $1;", ItemName); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("error querying item from public.items table with name of %v %v", ItemName, err)
		return
	}

	j, err := json.Marshal(item)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("error marshalling items into json %v", err)
		return
	}

	w.Write(j)
}
