package main

import (
    "fmt"
    "github.com/k-capehart/go-salesforce"
)

func main() {
    // Credentials ko struct mein daalo
    creds := salesforce.Creds{
        ConsumerKey:    "3MVG9k02hQhyUgQBmbVUDclU6E9MvUXXWR1MH3646RfDG6EjKW21SxAxQPEDBUUTniBqYT1SHzKjA2HQtm2.8",
        ConsumerSecret: "14AE4A83994E9A4285177047367BD85AED8C319A7C4FED8CCB70795C61E58439",
        Username:       "auralwc2024lwc@gmail.com",
        Password:       "",
        SecurityToken:  "rKkqiK4rdisrtoQBQritXqUH",
        Domain:         "https://login.salesforce.com", // Yaha https:// add kiya
    }

    // Salesforce initialize karo
    sf, err := salesforce.Init(creds)
    if err != nil {
        fmt.Println("Connection Error:", err)
        return
    }

    // Account struct define karo jo query result ko hold karega
    type Account struct {
        Id   string `json:"Id"`
        Name string `json:"Name"`
    }

    // Query karo aur result ko slice mein decode karo
    var accounts []Account
    err = sf.Query("SELECT Id, Name FROM Account LIMIT 5", &accounts)
    if err != nil {
        fmt.Println("Query Error:", err)
        return
    }

    // Results print karo
    for _, account := range accounts {
        fmt.Printf("ID: %s, Name: %s\n", account.Id, account.Name)
    }
}