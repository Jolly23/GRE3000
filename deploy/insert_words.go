package main

import (
	"GRE3000/database"
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	// read csv file in deploy/words.csv
	// insert words into database
	file, err := os.Open("deploy/再要你命3000电子版.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	db := database.NewDatabase()
	// read csv file
	f := csv.NewReader(file)
	records, err := f.ReadAll()
	if err != nil {
		panic(err)
	}
	for _, record := range records {
		// insert into database
		// insert into vocabulary(word, mean) values (record[0], record[1])
		fmt.Println(record[0], record[1])
		db.InsertWord(record[0], record[1])
	}
}
