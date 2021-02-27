package cbr

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)


type Valute struct {
	XMLName  string  `xml:"Valute"`
	NumCode  int64   `xml:"NumCode"`
	CharCode string  `xml:"CharCode"`
	Nominal  int64   `xml:"Nominal"`
	Name     string  `xml:"Name"`
	Value    float64 `xml:"Value"`
}

type ValCurs struct {
	XMLName string    `xml:"ValCurs"`
	Valutes []Valute   `xml:"Valute"`
}

type ValCursJSON struct {
	Code  string  `json:"code"`
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

func Extract(cbrUrl string) error {
	req, err := http.Get(cbrUrl)
	if err != nil {
		log.Println(err)
		return err
	}
	defer req.Body.Close()

	if req.StatusCode != http.StatusOK {
		log.Println(err)
		return err
	}

	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
		return err
	}

	parsedData, err := parseXml(data)
	if err != nil {
		log.Println(err)
		return err
	}

	jsonData := convertXmlToJson(parsedData)

	err = exportToJson(jsonData)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func parseXml(data []byte) (ValCurs, error) {
	var decoded ValCurs

	err := xml.Unmarshal(data, &decoded)
	if err != nil {
		log.Println(err)
		return ValCurs{}, err
	}

	return decoded, err
}

func convertXmlToJson(xmlData ValCurs) []ValCursJSON {
	var jsonData []ValCursJSON
	for _, v := range xmlData.Valutes{
		jsonData = append(jsonData, ValCursJSON{
			Code:  v.CharCode,
			Name:  v.Name,
			Value: v.Value,
		})
	}
	return jsonData
}

func exportToJson(data []ValCursJSON) error {
	file, err := os.Create("currencies.json")
	if err != nil {
		log.Println(err)
		return err
	}

	defer func(c io.Closer) {
		if err := c.Close(); err != nil {
			log.Println(err)
		}
	}(file)


	err = json.NewEncoder(file).Encode(data)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}