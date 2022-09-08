package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type booklogInfo struct {
	Tana struct {
		Account  string `json:"account"`
		Name     string `json:"name"`
		ImageUrl string `json:"image_url"`
	}
	Category []string `json:"category"`
	Books    []struct {
		Url     string `json:"url"`
		Title   string `json:"title"`
		Image   string `json:"image"`
		Catalog string `json:"catalog"`
	}
}

const baseUrl = "http://api.booklog.jp/v2/json/"

var (
	bookGage = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "booklog",
			Name:      "read_books_total",
			Help:      "Read books",
		}, []string{"user"},
	)
)

func main() {
	prometheus.Register(bookGage)

	http.Handle("/metrics", promhttp.Handler())

	go func() {
		setValue()
	}()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func setValue() {
	n := len(getBooklogInfo().Books)
	f := float64(n)
	labels := prometheus.Labels{
		"user": getBooklogInfo().Tana.Name,
	}
	bookGage.With(labels).Set(f)
	time.Sleep(86400 * time.Second) // 1日間隔
}

func loadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Printf("Can't load file: %v", err)
	}
}

func getBooklogInfo() booklogInfo {
	loadEnv()

	accountName := os.Getenv("ACCOUNT_ID")
	url := baseUrl + accountName

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	jsonpStr := string(body)

	var info booklogInfo
	json.Unmarshal([]byte(jsonpStr), &info)
	return info
}
