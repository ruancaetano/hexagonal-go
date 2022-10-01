package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ruancaetano/hexagonal-go/adapters/dto"
	"github.com/ruancaetano/hexagonal-go/application"
	"github.com/urfave/negroni"
)

func MakeProductHandlers(r *mux.Router, n *negroni.Negroni, service application.ProductServiceInterface) {
	r.Handle("/products/{id}", n.With(
		negroni.Wrap(getProduct(service)),
	)).Methods("GET", "OPTIONS")

	r.Handle("/products", n.With(
		negroni.Wrap(createProduct(service)),
	)).Methods("POST", "OPTIONS")

	r.Handle("/products/{id}/enable", n.With(
		negroni.Wrap(enableProduct(service)),
	)).Methods("PATCH", "OPTIONS")

	r.Handle("/products/{id}/disable", n.With(
		negroni.Wrap(disableProduct(service)),
	)).Methods("PATCH", "OPTIONS")
}

func getProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		product, err := getProductByRequestParam(r, service)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write(jsonError(err.Error()))
			return
		}

		encondeProductResponse(w, product)
	})
}

func createProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var productDto dto.ProductDto

		err := json.NewDecoder(r.Body).Decode(&productDto)
		if err != nil {
			formatErrorResponse(w, err, http.StatusBadRequest)
			return
		}

		product, err := service.Create(productDto.Name, productDto.Price)
		if err != nil {
			formatErrorResponse(w, err, http.StatusBadRequest)
			return
		}

		encondeProductResponse(w, product)
	})
}

func enableProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		product, err := getProductByRequestParam(r, service)
		if err != nil {
			formatErrorResponse(w, err, http.StatusNotFound)
			return
		}

		result, err := service.Enable(product)
		if err != nil {
			formatErrorResponse(w, err, http.StatusBadRequest)
			return
		}

		encondeProductResponse(w, result)
	})
}

func disableProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		product, err := getProductByRequestParam(r, service)
		if err != nil {
			formatErrorResponse(w, err, http.StatusNotFound)
			return
		}

		result, err := service.Disable(product)
		if err != nil {
			formatErrorResponse(w, err, http.StatusBadRequest)
			return
		}

		encondeProductResponse(w, result)
	})
}

func getProductByRequestParam(r *http.Request, service application.ProductServiceInterface) (application.ProductInterface, error) {
	vars := mux.Vars(r)
	id := vars["id"]
	product, err := service.Get(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func encondeProductResponse(w http.ResponseWriter, product application.ProductInterface) {
	err := json.NewEncoder(w).Encode(product)
	if err != nil {
		formatErrorResponse(w, err, http.StatusInternalServerError)
		return
	}
}

func formatErrorResponse(w http.ResponseWriter, err error, status int) {
	w.WriteHeader(status)
	w.Write(jsonError(err.Error()))
}
