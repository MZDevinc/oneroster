package orServices

import (
	"encoding/csv"
    "fmt"
	"os"
	"strconv"
	// "github.com/MZDevinc/oneroster/models"
) 


func ProccessFiles(filePath string)  string{
	fmt.Println("ProccessFiles -> filePath: ",filePath)
	return "ProccessFiles -> filePath: "+filePath;
}

func ReadUsers(filePath string)  string{
	fmt.Println("ProccessFiles -> filePath: ",filePath)
	var message string
	lines, err := ReadCsv(filePath)
    if err != nil {
		message = fmt.Sprintf("Can't read file filePath: %s ",filePath)
	
    }else{
		message = fmt.Sprintf("sucess read file filePath: %s lineNo: %s",filePath, strconv.Itoa( len(lines) ) )
	}

    // Loop through lines & turn into object
   
	fmt.Println(message)
	return message
}

func ReadCsv(filename string) ([][]string, error) {

    // Open CSV file
    f, err := os.Open(filename)
    if err != nil {
        return [][]string{}, err
    }
    defer f.Close()

    // Read File into a Variable
    lines, err := csv.NewReader(f).ReadAll()
    if err != nil {
        return [][]string{}, err
    }

    return lines, nil
}