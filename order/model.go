package order

type Order struct {
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
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
