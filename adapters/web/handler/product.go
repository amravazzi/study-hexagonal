package handler

import (
	"encoding/json"
	"net/http"

	"github.com/amravazzi/study-hexagonal/adapters/dto"
	"github.com/amravazzi/study-hexagonal/application"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func MakeProductHandlers(r *mux.Router, n *negroni.Negroni, service application.ProductServiceInterface) {
	r.Handle("/product/{id}", n.With(
		negroni.Wrap(getProduct(service)),
	)).Methods("GET", "OPTIONS")

	r.Handle("/product", n.With(
		negroni.Wrap(createProduct(service)),
	)).Methods("POST", "OPTIONS")

	r.Handle("/product/{id}/enable", n.With(
		negroni.Wrap(enableProduct(service)),
	)).Methods("GET", "OPTIONS")

	r.Handle("/product/{id}/disable", n.With(
		negroni.Wrap(disableProduct(service)),
	)).Methods("GET", "OPTIONS")
}

func getProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(req)
		id := vars["id"]
		product, err := service.Get(id)
		if err != nil {
			writer.WriteHeader(http.StatusNotFound)
			return
		}
		err = json.NewEncoder(writer).Encode(product)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}

func createProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		var productDto dto.Product
		err := json.NewDecoder(req.Body).Decode(&productDto)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write(jsonError(err.Error()))
			return
		}

		product, err := service.Create(productDto.Name, productDto.Price)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write(jsonError(err.Error()))
			return
		}

		err = json.NewEncoder(writer).Encode(product)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write(jsonError(err.Error()))
			return
		}
	})
}

func enableProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(req)
		id := vars["id"]
		product, err := service.Get(id)
		if err != nil {
			writer.WriteHeader(http.StatusNotFound)
			return
		}

		result, err := service.Enable(product)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write(jsonError(err.Error()))
			return
		}

		err = json.NewEncoder(writer).Encode(result)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}

func disableProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(req)
		id := vars["id"]
		product, err := service.Get(id)
		if err != nil {
			writer.WriteHeader(http.StatusNotFound)
			return
		}

		result, err := service.Disable(product)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write(jsonError(err.Error()))
			return
		}

		err = json.NewEncoder(writer).Encode(result)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}
