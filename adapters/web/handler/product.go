package handler

import (
	"encoding/json"
	"net/http"

	"github.com/HmmerHead/go-arquit/adapters/dto"
	app "github.com/HmmerHead/go-arquit/application"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func MakeProductHandlers(r *mux.Router, n *negroni.Negroni, service app.ProductServiceInterface) {
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

func getProduct(service app.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(
		func(rw http.ResponseWriter, r *http.Request) {
			rw.Header().Set("Content-Type", "application/json")
			vars := mux.Vars(r)
			id := vars["id"]

			product, err := service.Get(id)
			if err != nil {
				rw.WriteHeader(http.StatusNotFound)
				rw.Write(jsonError(err.Error()))
				return
			}

			err = json.NewEncoder(rw).Encode(product)
			if err != nil {
				rw.WriteHeader(http.StatusInternalServerError)
				return
			}
		})
}

func createProduct(service app.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(
		func(rw http.ResponseWriter, r *http.Request) {
			rw.Header().Set("Content-Type", "application/json")
			var productDto dto.ProductDto
			err := json.NewDecoder(r.Body).Decode(&productDto)
			if err != nil {
				rw.WriteHeader(http.StatusInternalServerError)
				rw.Write(jsonError(err.Error()))
				return
			}

			product, err := service.Create(productDto.Name, productDto.Price)
			if err != nil {
				rw.WriteHeader(http.StatusInternalServerError)
				rw.Write(jsonError(err.Error()))
				return
			}
			err = json.NewEncoder(rw).Encode(product)
			if err != nil {
				rw.WriteHeader(http.StatusInternalServerError)
				rw.Write(jsonError(err.Error()))
				return
			}
		})
}

func enableProduct(service app.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(
		func(rw http.ResponseWriter, r *http.Request) {
			rw.Header().Set("Content-Type", "application/json")
			vars := mux.Vars(r)
			id := vars["id"]

			product, err := service.Get(id)
			if err != nil {
				rw.WriteHeader(http.StatusNotFound)
				rw.Write(jsonError(err.Error()))
				return
			}
			result, err := service.Enable(product)
			if err != nil {
				rw.WriteHeader(http.StatusInternalServerError)
				rw.Write(jsonError(err.Error()))
				return
			}
			err = json.NewEncoder(rw).Encode(result)
			if err != nil {
				rw.WriteHeader(http.StatusInternalServerError)
				return
			}
		})
}

func disableProduct(service app.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(
		func(rw http.ResponseWriter, r *http.Request) {
			rw.Header().Set("Content-Type", "application/json")
			vars := mux.Vars(r)
			id := vars["id"]

			product, err := service.Get(id)
			if err != nil {
				rw.WriteHeader(http.StatusNotFound)
				rw.Write(jsonError(err.Error()))
				return
			}

			result, err := service.Disable(product)
			if err != nil {
				rw.WriteHeader(http.StatusInternalServerError)
				rw.Write(jsonError(err.Error()))
				return
			}
			err = json.NewEncoder(rw).Encode(result)
			if err != nil {
				rw.WriteHeader(http.StatusInternalServerError)
				return
			}
		})
}

