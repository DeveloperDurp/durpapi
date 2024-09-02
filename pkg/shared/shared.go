package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"gitlab.com/developerdurp/logger"
	"net/http"
	"reflect"
	"strings"
)

func GetParams(r *http.Request, data interface{}) (interface{}, error) {

	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		// Fallback to headers if body is not available or decoding fails
		v := reflect.ValueOf(data).Elem()
		t := v.Type()
		for i := 0; i < v.NumField(); i++ {
			fieldName := t.Field(i).Name
			queryParms := r.URL.Query()
			if queryParms == nil {
				return nil, errors.New("No parameters found in body or headers")
			}
			if queryParms != nil {
				fieldValue := v.Field(i)
				switch fieldValue.Kind() {
				case reflect.String:
					fieldValue.SetString(queryParms.Get(strings.ToLower(fieldName)))
					//	req.IP = queryParams.Get("IP")
				// Add cases here for other types as needed, e.g., int, bool
				default:
					logger.LogError(fmt.Sprintf(
						"Unsupported field type: %v for field %v", fieldValue.Kind(), fieldName),
					)
				}
			}
		}
	}
	return data, nil
}
