package main

import (
    "bufio"
    "encoding/json"
    "fmt"
    "io"
    "os"
)

func main() {
    employees, err := ReadJSONL("samples.jsonl")
    if err != nil {
        fmt.Println(err.Error())
    } else {
        fmt.Printf("%v", employees)
    }
}

func ReadJSONL(filename string) ([]Employee, error) {
    file, err := os.Open(filename)
    if err != nil { return nil, err }
    defer closeSilently(file)
    scanner := bufio.NewScanner(file)
    lineCounter := 0
    employees := make([]Employee, 0)
    for scanner.Scan() {
        var employee Employee
        err := json.Unmarshal(scanner.Bytes(), &employee)
        if err != nil {
            fmt.Printf("Error: failed to read line #%d: %s", lineCounter, err.Error())
        } else {
            employees = append(employees, employee)
        }
        lineCounter += 1
    }
    return employees, nil
}

func closeSilently(closer io.Closer) {
    _ = closer.Close()
}

type Employee struct {
    ID int             `json:"ID"`
    Name string        `json:"name"`
    Email string       `json:"email"`
    Languages []string `json:"lang"`
    Position struct{
        Title string
        Department string
    } `json:"position"`
}
