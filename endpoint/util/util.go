package util

import (
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/javinc/graham/model"
)

// Output response
func Output(c *gin.Context, data interface{}, err error) {
	if err != nil {
		OutputError(c, err)

		return
	}

	// okay
	c.JSON(http.StatusOK, data)
}

// OutputError response
func OutputError(c *gin.Context, err error) {
	stat := http.StatusBadRequest

	// format generic error
	if _, ok := err.(*model.Error); !ok {
		err = &model.Error{
			Name:    "GENERIC",
			Message: err.Error(),
		}
	}

	e := err.(*model.Error)
	// check for internal errors
	if e.Panic {
		stat = http.StatusInternalServerError
	}

	c.JSON(stat, err)
}

// OutputNotFound response with 404
func OutputNotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, &model.Error{
		Name:    "NOT_FOUND",
		Message: "404 not found",
	})
}

// ParsePayload parse payload with error handling
func ParsePayload(c *gin.Context, p interface{}) error {
	// check method if needs a payload to read
	m := http.MethodPost + http.MethodPatch
	if strings.Contains(m, c.Request.Method) {
		err := c.BindJSON(p)
		if err != nil {
			// modify error for more info
			err = &model.Error{
				Name:    "INVALID_JSON",
				Message: err.Error(),
			}

			return err
		}
	}

	return nil
}

// ParseFilter get filter option
func ParseFilter(c *gin.Context, object interface{}) map[string]interface{} {
	prefix := "filter."
	params := c.Request.URL.Query()

	m := map[string]interface{}{}
	o := reflect.TypeOf(object)
	for i := 0; i < o.NumField(); i++ {
		field := o.Field(i).Tag.Get("json")
		kind := reflect.ValueOf(object).Field(i).Kind()

		// get first element because its the field name
		fieldName := strings.Split(field, ",")[0]

		// prefixed filter field name
		filterName := prefix + fieldName
		if v, ok := params[filterName]; ok {
			switch kind {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				num, _ := strconv.Atoi(v[0])
				m[fieldName] = num
			case reflect.String:
				m[fieldName] = v[0]
			case reflect.Bool:
				b, _ := strconv.ParseBool(v[0])
				m[fieldName] = b
			case reflect.Ptr:
				s := reflect.Indirect(reflect.ValueOf(v[0]))
				b, _ := strconv.ParseBool(s.String())
				m[fieldName] = b
			}
		}
	}

	return m
}
