package handlers

import (
	"crypto/rand"
	"log"
	"math/big"
	"net/http"
	"time"
)

var timeWeight []int

var defaultTime = time.Millisecond * 20

var squareDeviation = time.Millisecond * 10

func MainHandler(w http.ResponseWriter, r *http.Request) {

	randomBigInt, err := rand.Int(rand.Reader, big.NewInt(int64(len(timeWeight))))
	if err != nil {
		panic(err)
	}

	var weight = timeWeight[randomBigInt.Int64()]

	log.Println("sleep weight:", weight)
	if weight == 4 {
		log.Println("\t\tLONG REQUEST!")
	}

	// like go to db
	time.Sleep(defaultTime + squareDeviation*time.Duration(weight))

	w.WriteHeader(200)
	w.Write([]byte(examplePage))
}

// 68
// 26
// 4
// 0.2
func SetTimeWeight() {
	for i := 0; i < 640; i++ {
		timeWeight = append(timeWeight, 1)
	}
	for i := 0; i < 260; i++ {
		timeWeight = append(timeWeight, 2)
	}
	for i := 0; i < 40; i++ {
		timeWeight = append(timeWeight, 3)
	}
	for i := 0; i < 2; i++ {
		timeWeight = append(timeWeight, 4)
	}
}
