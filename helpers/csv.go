package helpers

import (
	"encoding/csv"
	"log"
	"os"
	"reza/scrapper-test/model"
	"reza/scrapper-test/model/constant"
)

func WriteCsv(req model.CreateRequest) error {
	_, str := GetNow()
	filename := constant.DefaultCsvFileName + str
	file, err := os.Create("resources/scrapper/" + filename)
	defer file.Close()
	if err != nil {
		log.Fatalln("failed to open file", err)
		return err
	}
	w := csv.NewWriter(file)
	defer w.Flush()

	row := []string{req.Name, req.Description, req.Price, req.Rating, req.MerchantName}
	if err := w.Write(row); err != nil {
		log.Fatalln("error writing record to file", err)
		return err
	}

	return nil
}

func WriteCsvBulk(req []model.CreateRequest) error {
	_, str := GetNow()
	filename := constant.DefaultCsvFileName + str
	file, err := os.Create("resources/scrapper/" + filename)
	defer file.Close()
	if err != nil {
		log.Fatalln("failed to open file", err)
		return err
	}
	w := csv.NewWriter(file)
	defer w.Flush()

	for _, record := range req {
		row := []string{record.Name, record.Description, record.Price, record.Rating, record.MerchantName}
		if err := w.Write(row); err != nil {
			log.Fatalln("error writing record to file", err)
			return err
		}
	}

	return nil
}
