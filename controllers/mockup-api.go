package controllers

import (
	"encoding/json"
	"log"
	"main/helpers"
	"main/models"
	"net/http"
	"strings"
	"time"

	"github.com/asdine/storm/v3/q"
	faker "github.com/brianvoe/gofakeit/v5"
)

type any = interface{}
type array = []any
type object = map[string]any

func genFakeData(fakeType any, fakeExtra any) any {
	rawFakeData := func() any {
		switch fakeType {
		case "boolean":
			return faker.Bool()
		case "int":
			return faker.Int64()
		case "float":
			return faker.Float64()
		case "string":
			return faker.Sentence(5)
		case "date":
			return faker.Date().Format(time.RFC3339)
		default:
			return ""
		}
	}()

	return formatFakeData(rawFakeData, fakeExtra)
}

func formatFakeData(rawFakeData any, fakeExtra any) any {
	if fakeExtra == nil {
		return rawFakeData
	}

	// min := (fakeExtra.(object)["min"].(float64))
	// max := (fakeExtra.(object)["max"].(float64))
	// format := fakeExtra.(object)["max"].(string)

	formatInt := func() int {
		return rawFakeData.(int)
	}

	formatString := func() string {
		return rawFakeData.(string)
	}

	switch rawFakeData.(type) {
	case int:
		return formatInt()
	case string:
		return formatString()
	default:
		return rawFakeData
	}
}

func parseDataRecursive(dataModel object) object {
	data := make(object)
	faker.Seed(0)

	for key, value := range dataModel {
		switch value.(type) {
		case object:
			if strings.Contains(key, ".count") || strings.Contains(key, ".format") {
				break
			}
			data[key] = parseDataRecursive(value.(object))
		case array:
			extra := dataModel[key+".count"]
			var count int
			if extra != nil {
				min := int(extra.(object)["min"].(float64))
				max := int(extra.(object)["max"].(float64))
				count = faker.Number(min, max)
			} else {
				count = 1
			}

			data[key] = make(array, count)
			item := value.(array)[0].(object)

			for index := range data[key].(array) {
				data[key].(array)[index] = parseDataRecursive(item)
			}
		default:
			extra := dataModel[key+".format"]
			data[key] = genFakeData(value, extra)
		}
	}

	return data
}

// MockupAPIHandler ...
/** TODO: phase 2: consistent return
- stored fake data into db for next use
*/
func MockupAPIHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := helpers.OpenDB()
	defer helpers.CloseDB()

	var mockupConfig models.MockupConfig
	query := db.Select(
		q.And(q.Eq("URL", r.URL.String()), q.Eq("Method", r.Method)),
	)
	err := query.First(&mockupConfig)
	if err != nil {
		log.Fatal(err)
	}

	var dataModel object
	err = json.Unmarshal([]byte(mockupConfig.DataModel), &dataModel)
	if err != nil {
		log.Fatal(err)
	}

	data := parseDataRecursive(dataModel)
	json.NewEncoder(w).Encode(data)
}
