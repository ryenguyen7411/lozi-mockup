package controllers

import (
	"encoding/json"
	"log"
	"main/helpers"
	"main/models"
	"math"
	"net/http"
	"strings"
	"time"

	"github.com/asdine/storm/v3/q"
	faker "github.com/brianvoe/gofakeit/v5"
	"github.com/iancoleman/strcase"
)

type any = interface{}
type array = []any
type object = map[string]any

func genFakeData(fakeType any, fakeExtra any) any {
	// round
	fakeDataInt := func() int {
		min := fakeExtra.(object)["min"]
		max := fakeExtra.(object)["max"]

		res := func() int {
			if min == nil && max == nil {
				return faker.Number(0, math.MaxInt16)
			}
			if min == nil {
				return faker.Number(math.MinInt16, int(max.(float64)))
			}
			if max == nil {
				return faker.Number(int(min.(float64)), math.MaxInt16)
			}
			return faker.Number(int(min.(float64)), int(max.(float64)))
		}()

		return res
	}

	fakeDataFloat := func() float64 {
		min := fakeExtra.(object)["min"]
		max := fakeExtra.(object)["max"]
		precision := fakeExtra.(object)["precision"]

		res := func() float64 {
			if min == nil && max == nil {
				return faker.Float64Range(0, math.MaxInt16)
			}
			if min == nil {
				return faker.Float64Range(math.MinInt16, max.(float64))
			}
			if max == nil {
				return faker.Float64Range(min.(float64), math.MaxInt16)
			}
			return faker.Float64Range(min.(float64), max.(float64))
		}()

		if precision == nil {
			return math.Round(res*math.Pow(10, 3)) / math.Pow(10, 3)
		}
		return math.Round(res*math.Pow(10, precision.(float64))) / math.Pow(10, precision.(float64))
	}

	// enum
	fakeDataString := func() string {
		var res string
		word := fakeExtra.(object)["word"]
		format := fakeExtra.(object)["format"]

		if word == nil {
			res = faker.Sentence(2)
		} else {
			res = faker.Sentence(int(word.(float64)))
		}
		res = strings.TrimSuffix(res, ".")

		if format == nil {
			return res
		}
		if format == "username" {
			return strings.Replace(strings.ToLower(res), " ", "", -1)
		}
		if format == "slug" {
			return strcase.ToKebab(res)
		}
		if format == "imageUrl" {
			return faker.ImageURL(512, 512)
		}
		if format == "phone" {
			return faker.Phone()
		}

		return res
	}

	fakeData := func() any {
		switch fakeType {
		case "boolean":
			return faker.Bool()
		case "int":
			return fakeDataInt()
		case "float":
			return fakeDataFloat()
		case "string":
			return fakeDataString()
		case "date":
			return faker.Date().Format(time.RFC3339)
		default:
			return ""
		}
	}()

	return fakeData
}

func parseDataRecursive(dataModel object) object {
	data := make(object)
	faker.Seed(0)

	for key, value := range dataModel {
		switch value.(type) {
		case object:
			if value.(object)["type"] == "array" {
				len := value.(object)["len"]
				if len == nil {
					len = faker.Number(0, 24)
				} else {
					len = int(len.(float64))
				}

				data[key] = make(array, len.(int))
				item := value.(object)["items"].(object)

				for index := range data[key].(array) {
					data[key].(array)[index] = parseDataRecursive(item)
				}
				break
			}
			if value.(object)["type"] == "object" {
				data[key] = parseDataRecursive(value.(object)["properties"].(object))
				break
			}
			data[key] = genFakeData(value.(object)["type"], value.(object))
		default:
			data[key] = genFakeData(value, nil)
		}
	}

	return data
}

// MockupAPIHandler ...
/** TODO: phase 2:
- failedRatio
- consistent return, stored fake data into db for next use
- POST / PUT / DELETE
- tool for control dataModels
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
		json.NewEncoder(w).Encode("")
		return
	}

	var dataModel object
	err = json.Unmarshal([]byte(mockupConfig.DataModel), &dataModel)
	if err != nil {
		log.Fatal(err)
	}

	data := parseDataRecursive(dataModel)
	json.NewEncoder(w).Encode(data)
}
