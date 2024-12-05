package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

type IPStackResponse struct {
	IP          string  `json:"ip"`
	CountryName string  `json:"country_name"`
	City        string  `json:"city"`
	RegionName  string  `json:"region_name"`
	Zip         string  `json:"zip"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
}


func main() {
    err := godotenv.Load()
    if err != nil {
        fmt.Printf("Error reading from .env file: %v", err)
    }
    apiKey := os.Getenv("API_KEY")
    if apiKey == "" {
        fmt.Printf("API Key not found in .env file")
    }

    db, err := sql.Open("sqlite3", "./ipinfo.db")
    if err != nil {
        fmt.Printf("Failed to open database: %v", err)
    }
    defer db.Close()
    
    createTableSql := `CREATE TABLE IF NOT EXISTS ipinfo(
        ip TEXT PRIMARY KEY,
        country TEXT,
        city TEXT,
        region TEXT,
        zip TEXT,
        latitude REAL,
        longitude REAL
    );`

    _, err = db.Exec(createTableSql)
    if err != nil {
        fmt.Printf("Failed to create table: %v", err)
    }

    var ip string
    fmt.Println("Enter the IP you want to know the information about: ")
    fmt.Scanf("%s", &ip)

    // Read from DB 
    row := db.QueryRow("SELECT * FROM ipinfo WHERE ip = ?", ip)
    var result IPStackResponse
    err = row.Scan(&result.IP, &result.CountryName, &result.City, &result.RegionName, &result.Zip, &result.Latitude, &result.Longitude)
    
    if err == sql.ErrNoRows {
        url := fmt.Sprintf("http://api.ipstack.com/%s?access_key=%s", ip, apiKey)
        resp, err := http.Get(url)
        if err != nil {
            fmt.Println(err)
        }
        defer resp.Body.Close()
        if resp.StatusCode != http.StatusOK {
	        fmt.Printf("Error: Received HTTP %d", resp.StatusCode)
	    } 
        body, err := io.ReadAll(resp.Body)
        if err != nil {
            fmt.Println(err)
        }
        err = json.Unmarshal(body, &result)
        if err != nil {
            fmt.Println(err)
        }
        _, err = db.Exec("INSERT INTO ipinfo (ip, country, city, region, zip, latitude, longitude) VALUES (?, ?, ?, ?, ?, ?, ?)",result.IP, result.CountryName, result.City, result.RegionName, result.Zip, result.Latitude, result.Longitude)
		if err != nil {
			fmt.Printf("Failed to insert data into database: %v", err)
		} 
    } else if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("IP found in database")
    }
    // Access the values
	fmt.Printf("IP: %s\n", result.IP)
	fmt.Printf("Country: %s\n", result.CountryName)
	fmt.Printf("City: %s\n", result.City)
	fmt.Printf("Latitude: %f, Longitude: %f\n", result.Latitude, result.Longitude)

}

