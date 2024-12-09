package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Product struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Pricing struct {
	ID    string  `json:"id"`
	Price float64 `json:"price"`
}

type Review struct {
	ProductID string `json:"product_id"`
	Rating    int    `json:"rating"`
}

func fetchProduct(id string) (*Product, error) {
	resp, err := http.Get(fmt.Sprintf("http://product-service/products/%s", id))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var product Product
	err = json.NewDecoder(resp.Body).Decode(&product)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func fetchPricing(id string) (*Pricing, error) {
	resp, err := http.Get(fmt.Sprintf("http://pricing-service/price/%s", id))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var pricing Pricing
	err = json.NewDecoder(resp.Body).Decode(&pricing)
	if err != nil {
		return nil, err
	}
	return &pricing, nil
}

func fetchReview(id string) (*Review, error) {
	resp, err := http.Get(fmt.Sprintf("http://review-service/reviews/%s", id))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var review Review
	err = json.NewDecoder(resp.Body).Decode(&review)
	if err != nil {
		return nil, err
	}
	return &review, nil
}

func getProductDetails(w http.ResponseWriter, r *http.Request) {
	productID := r.URL.Query().Get("id")

	product, err := fetchProduct(productID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	pricing, err := fetchPricing(productID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	review, err := fetchReview(productID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := struct {
		Product *Product `json:"product"`
		Pricing *Pricing `json:"pricing"`
		Review  *Review  `json:"review"`
	}{
		Product: product,
		Pricing: pricing,
		Review:  review,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/product", getProductDetails)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
