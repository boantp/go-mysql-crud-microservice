package main

import (
	"net/http"

	"github.com/boantp/go-api-ecomm/car"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/cars", car.Index)
	http.HandleFunc("/cars/show", car.Show)
	http.HandleFunc("/cars/create", car.Create)
	http.HandleFunc("/cars/create/process", car.CreateProcess)
	http.HandleFunc("/cars/update", car.Update)
	http.HandleFunc("/cars/update/process", car.UpdateProcess)
	http.HandleFunc("/cars/delete/process", car.DeleteProcess)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/cars", http.StatusSeeOther)
}
