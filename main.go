package main

import (
	"net/http"

	_ "github.com/boantp/go-api-ecomm/car"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/cars", car.Index)
	// http.HandleFunc("/cars/show", car.Show)
	// http.HandleFunc("/order/create", order.Create)
	// http.HandleFunc("/order/create/process", order.CreateProcess)
	// http.HandleFunc("/cars/update", car.Update)
	// http.HandleFunc("/order/update/process", order.UpdateProcess)
	// http.HandleFunc("/order/delete/process", order.DeleteProcess)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/cars", http.StatusSeeOther)
}
