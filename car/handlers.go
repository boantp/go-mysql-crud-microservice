package car

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	cars, err := AllCars()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	d := responseData{RespCode: http.StatusText(200), RespDesc: "success", Data: cars}

	uj, err := json.Marshal(d)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
}

func Show(w http.ResponseWriter, r *http.Request) {
	//code here
}

func Update(w http.ResponseWriter, r *http.Request) {
	//code here
}

func UpdateProcess(w http.ResponseWriter, r *http.Request) {
	//code here
}
