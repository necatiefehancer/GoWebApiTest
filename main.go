package main

import (
	"api/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

func SingleProductHandler(writer http.ResponseWriter, request *http.Request) {

	params := mux.Vars(request)
	id := params["id"]
	product := loadSingleProduct(id)
	template, err := template.ParseFiles("./templates/singleProduct.html")
	if err != nil {
		log.Fatal("Template Parse Error ", err.Error())
	}
	if template.Execute(writer, product) != nil {
		log.Fatal("Execute Error ", err.Error())
	}

}

func AllProductHandler(writer http.ResponseWriter, request *http.Request) {
	products := loadAllProducts()
	template, err := template.ParseFiles("./templates/allProducts.html")
	if err != nil {
		log.Fatal("Template Parse Error ", err.Error())
	}
	if template.Execute(writer, products) != nil {
		log.Fatal("Execute Error ", err.Error())
	}
}

func loadSingleProduct(ProductCode string) models.Product {
	bytesData, err := ioutil.ReadFile("./data/Products.json")
	if err != nil {
		log.Fatal("Read Data byte err")
	}
	var products models.Products
	if json.Unmarshal(bytesData, &products) != nil {
		log.Fatal("UnMarshal err")
	}

	for _, value := range products.ProductsData {
		if value.ProductCode == ProductCode {
			return value
		}
	}

	return models.Product{}
}

func loadAllProducts() models.Products {
	bytesData, err := ioutil.ReadFile("./data/Products.json")
	if err != nil {
		log.Fatal("Convert data bytes err ", err.Error())
	}
	var productsData models.Products
	if json.Unmarshal(bytesData, &productsData) != nil {
		log.Fatal("Json Unmarshal error ", err.Error())
		return models.Products{}
	}
	return productsData
}

func main() {

	var router = mux.NewRouter()
	router.HandleFunc("/api/product/{id}", SingleProductHandler)
	router.HandleFunc("/api/products", AllProductHandler)
	http.Handle("/", router)
	http.ListenAndServe(":9000", nil)

	// fmt.Println(loadAllProducts())

}
