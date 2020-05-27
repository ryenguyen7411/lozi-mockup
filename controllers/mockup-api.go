package controllers

import (
	"encoding/json"
	"log"
	"main/helpers"
	"main/models"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/asdine/storm/q"
)

type any = interface{}
type array = []any
type object = map[string]any

func genFakeData(fakeType any, fakeExtra any) any {
	rawFakeData := func() any {
		switch fakeType {
		case "boolean":
			return true
		case "int":
			return 0
		case "float":
			return 1.0
		case "string":
			return "Fake string"
		case "date":
			return "29/05/2020"
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
				rand.Seed(time.Now().UnixNano())
				min := int(extra.(object)["min"].(float64))
				max := int(extra.(object)["max"].(float64))
				count = rand.Intn(max-min+1) + min
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
