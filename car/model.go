package car

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/boantp/go-api-ecomm/config"
)

type Car struct {
	CarId        string
	CarName      string
	CarYear      string
	DefaultPrice float32
	CarStatus    int
}

//if we want to give json reponses
type responseData struct {
	RespCode string
	RespDesc string
	Data     []Car
}

func AllCars() ([]Car, error) {
	rows, err := config.DB.Query("SELECT * FROM car")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cars := make([]Car, 0)
	for rows.Next() {
		car := Car{}
		err := rows.Scan(&car.CarId, &car.CarName, &car.CarYear, &car.DefaultPrice, &car.CarStatus) // order matters
		if err != nil {
			return nil, err
		}
		cars = append(cars, car)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return cars, nil
}

func OneCar(r *http.Request) (Car, error) {
	car := Car{}
	carid := r.FormValue("carid")
	if carid == "" {
		return car, errors.New("400. Bad Request.")
	}

	row := config.DB.QueryRow("SELECT * FROM car WHERE car_id=?", carid)

	err := row.Scan(&car.CarId, &car.CarName, &car.CarYear, &car.DefaultPrice, &car.CarStatus)
	if err != nil {
		return car, err
	}

	return car, nil
}

func PutCar(r *http.Request) (Car, error) {
	// get form values
	car := Car{}
	car.CarId = r.FormValue("carid")
	car.CarName = r.FormValue("carname")
	car.CarYear = r.FormValue("caryear")
	p := r.FormValue("defaultprice")

	// validate form values
	if car.CarId == "" || car.CarName == "" || car.CarYear == "" || p == "" {
		return car, errors.New("400. Bad request. All fields must be complete.")
	}

	// convert form values
	f64, err := strconv.ParseFloat(p, 32)
	if err != nil {
		return car, errors.New("406. Not Acceptable. Price must be a number.")
	}
	car.DefaultPrice = float32(f64)

	// insert values
	sqlStr := "INSERT INTO car(car_id, car_name, car_year, default_price, car_status) VALUES (?, ?, ?, ?, ?)"
	//prepare the statement
	stmt, _ := config.DB.Prepare(sqlStr)

	//format all vals at once
	_, err = stmt.Exec(car.CarId, car.CarName, car.CarYear, p, 1)
	if err != nil {
		return car, errors.New("500. Internal Server Error." + err.Error())
	}
	return car, nil
}

func UpdateCar(r *http.Request) (Car, error) {
	// get form values
	car := Car{}
	car.CarId = r.FormValue("carid")
	car.CarName = r.FormValue("carname")
	car.CarYear = r.FormValue("caryear")
	p := r.FormValue("defaultprice")

	if car.CarId == "" || car.CarName == "" || car.CarYear == "" || p == "" {
		return car, errors.New("400. Bad Request. Fields can't be empty.")
	}

	// convert form values
	f64, err := strconv.ParseFloat(p, 32)
	if err != nil {
		return car, errors.New("406. Not Acceptable. Enter number for price.")
	}
	car.DefaultPrice = float32(f64)

	// insert values
	sqlStr := "update car set car_id=?, car_name=?, car_year=?, default_price=? WHERE car_id=?"
	//prepare the statement
	stmt, dberr := config.DB.Prepare(sqlStr)
	checkErr(dberr)

	//format all vals at once
	_, err = stmt.Exec(car.CarId, car.CarName, car.CarYear, p, car.CarId)
	checkErr(err)
	if err != nil {
		return car, err
	}
	return car, nil
}

func DeleteCar(r *http.Request) error {
	carid := r.FormValue("carid")
	if carid == "" {
		return errors.New("400. Bad Request.")
	}

	stmt, err := config.DB.Prepare("delete from car where car_id=?")
	checkErr(err)

	_, err = stmt.Exec(carid)
	checkErr(err)
	if err != nil {
		return errors.New("500. Internal Server Error")
	}
	return nil
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
