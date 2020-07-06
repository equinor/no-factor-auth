package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"sort"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestToken(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	claims := make(map[string]interface{})
	claims["ttt"] = "TTT"
	if !assert.NoError(t, Token(claims)(c)) {
		return
	}

	if !assert.Equal(t, http.StatusOK, rec.Code) {
		return
	}

}

func contains(s []string, searchterm string) bool {
	sort.Strings(s)
	i := sort.SearchStrings(s, searchterm)
	return i < len(s) && s[i] == searchterm
}

func extractArrayElement(result map[string]interface{}) []string {
	var items []string

	object := reflect.ValueOf(result["arr_elem_key"])
	for i := 0; i < object.Len(); i++ {
		s := fmt.Sprint(object.Index(i).Interface())
		items = append(items, s)
	}

	return items
}
