package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

// A global Employeess array to simulate a database

type Employee struct {
    Fname string `json:"fname"`
    Lname string `json:"lname"`
    Email string `json:"email"`
}

type Employees []Employee

func returnAllEmployees(w http.ResponseWriter, r *http.Request){
    employees := Employees{
	Employee{Fname:"Basem", Lname:"Sousan", Email:"basemsousan@gmail.com"},
    }

    fmt.Println("Endpoint Hit: All Employees Endpoint")
    json.NewEncoder(w).Encode(employees)

}

func handleRequests() {
    http.HandleFunc("/employees", returnAllEmployees)
    http.Handle("/metrics", promhttp.Handler())
    log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
    handleRequests()
}
