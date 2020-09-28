package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
)

type HeightWeight struct {
	Height float64 `json:"height"`
	Weight float64 `json:"weight"`
}

type Bmis struct {
	Bmi float64 `json:"bmi"`
}

func calc_bmi(hw HeightWeight) Bmis {
	var bmi Bmis
	bmi.Bmi = hw.Weight / math.Pow(hw.Height/100, 2)
	return bmi
}

func main() {
	// JSON形式の文字列を定義
	jsontest := `{"height": 170, "weight": 50}`
	// HeightWeight型の変数の定義
	var heightweight HeightWeight
	// jsontestを読み取りheightweightの変数に値を書き込む
	if err := json.Unmarshal([]byte(jsontest), &heightweight); err != nil {
		log.Fatal(err)
	}
	bmi := calc_bmi(heightweight)
	jsonBytes, err := json.MarshalIndent(bmi, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("HeightWeight構造体")
	fmt.Println(heightweight)
	fmt.Println("\nBMIのJSON")
	fmt.Println(string(jsonBytes))
}
