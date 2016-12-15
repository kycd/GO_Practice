package main

import(
    "fmt"
    "os"
    // "strings"
    // "io"
    // "encoding/csv"
)

func main() {
    // Check parameter
    if(!isValidParameter()) {
        fmt.Println(usage())
        return
    }

    switch os.Args[1] {
    case "-c":
        toJson()
    case "-t":
        toCsv()
    }
}

// converter. from csv to json
func toJson() {

}

// converter. from json to csv
func toCsv() {

}

// valider of os args
func isValidParameter() bool{
    if(len(os.Args) < 3 || len(os.Args) > 4) {
        return false
    }

    return true
}

// usage for this program
func usage() string {
    str := "[Usage]\n"
    str += "002_CsvJsonConvert -c <csv_file_name> [<json_file_name>]\n"
    str += "002_CsvJsonConvert -j <json_file_name> [<csv_file_name>]"
    return str
}
