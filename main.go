package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
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

func handler(w http.ResponseWriter, r *http.Request) {
	// POST かつ json形式でないなら処理しない
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 長さr.ContentLength の byte配列．
	body := make([]byte, r.ContentLength)
	// bodyにリクエストのbodyを代入
	length, err := r.Body.Read(body)
	if err != nil && err != io.EOF {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var hw HeightWeight
	// bodyをパースしてHeightWeightの構造体hwの変数に代入する
	err = json.Unmarshal(body[:length], &hw)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	bmi := calc_bmi(hw)
	// bmis -> json
	jsonBytes, err := json.MarshalIndent(bmi, "", "    ")
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	// json形式であることを明示
	w.Header().Set("content-type", "application/json")
	// 計算したbmiをjson形式で返す
	fmt.Fprintf(w, string(jsonBytes))
	fmt.Fprintf(w, "\n")
}

func main() {
	// http://localhost:5000 にhandler関数を割り当て
	http.HandleFunc("/", handler)
	// サーバ起動
	http.ListenAndServe(":5000", nil)
}
