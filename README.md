# IP Information Fetcher

A Go application that retrieves information about a given IP address using the IPStack API and stores the data in a SQLite database. If the IP information already exists in the database, it fetches the data from there instead of calling the API.

## Features

- Fetch IP information such as country, city, region, zip code, latitude, and longitude.
- Cache IP information in a SQLite database to minimize API calls.
- Reads the API key securely from a `.env` file.

## Prerequisites

- Go (1.19 or later)
- SQLite
- [IPStack API Key](https://ipstack.com/)
- `godotenv` Go package for loading `.env` files.

## Installation

1. Clone this repository:
   ```bash
   git clone https://github.com/abolinjast/ipinfo.git
   cd ip-info-fetcher

    Install dependencies:

go mod tidy

Create a .env file in the project root and add your IPStack API key:

API_KEY=your_api_key_here

Run the application:

    go run main.go

Usage

    Run the application:

go run main.go

Enter the IP address you want to fetch information about when prompted:

Enter the IP you want to know the information about: 134.201.250.155

The application will:

    Check if the IP exists in the SQLite database.
    If found, fetch the data from the database.
    If not, call the IPStack API, store the data in the database, and display it.

Output example:

    IP: 134.201.250.155
    Country: United States
    City: Los Angeles
    Region: California
    Zip: 90001
    Latitude: 34.052235, Longitude: -118.243683

Project Structure

.
├── main.go         # Main application logic
├── go.mod          # Go module file
├── go.sum          # Dependencies checksum
├── ipinfo.db       # SQLite database (auto-created if not present)
└── .env            # Environment variables file (not included in the repo)

Dependencies

    godotenv - Load environment variables from .env.
    sqlite3 - SQLite driver for Go.
