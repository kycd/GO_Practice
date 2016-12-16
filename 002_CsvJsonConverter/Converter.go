package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func main() {
    // Check parameter
    if !isValidParameter() {
        fmt.Println(usage())
        return
    }

    file, err := os.Open(os.Args[2])
    if err != nil {
        fmt.Println(err)
        return
    }

    parseResult := ""
    switch os.Args[1] {
    case "-c":
        parseResult, err = csvToJson(file)
    case "-t":
        parseResult, err = jsonToCsv(file)
    }

    if err != nil {
        fmt.Println(err)
        return
    }

    if len(os.Args) == 3 {
        // output to stdout
        fmt.Println(parseResult)
    } else {
        // output to file
        outputToFile(parseResult, os.Args[3])
    }
}

//
func outputToFile(parseResult, filename string) {
    outFile, err := os.Create(filename)
    if err != nil {
        fmt.Println(err)
        return
    }
    outFile.WriteString(parseResult)
}

// converter. from csv to json
func csvToJson(f *os.File) (string, error) {
    readerCsv := csv.NewReader(bufio.NewReader(f))
    cells, err := readerCsv.ReadAll()

    if err != nil {
        return "", err
    }

    recordFullLine := make([]string, 0)
    for i := 1; i < len(cells); i++ {
        recordSingleLine := make([]string, 0)
        for j := 0; j < len(cells[i]); j++ {
            recordSingleLine = append(recordSingleLine, fmt.Sprintf("\"%s\": \"%s\"", cells[0][j], cells[i][j]))
        }
        recordFullLine = append(recordFullLine, fmt.Sprintf("{%s}", strings.Join(recordSingleLine, ", ")))
    }

    return fmt.Sprintf("[%s]", strings.Join(recordFullLine, ", ")), nil
}

// converter. from json to csv
func jsonToCsv(f *os.File) (string, error) {
    return "", nil
}

// valider of os args
func isValidParameter() bool{
    if len(os.Args) < 3 || len(os.Args) > 4 {
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
