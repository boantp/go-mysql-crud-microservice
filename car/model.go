package car

import (
	"github.com/boantp/go-api-ecomm/config"
)

type Car struct {
	CarId        int
	CarName      string
	CarYear      string
	DefaultPrice float32
	CarStatus    int
}

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
