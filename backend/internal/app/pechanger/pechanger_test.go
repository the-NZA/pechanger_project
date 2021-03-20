package pechanger

import (
	"fmt"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	configPath  = "../../../config/pechanger_config.json"
	testAmount  = 200
	testFrom    = "USD"
	testTo      = "RUB"
	testDate    = "2018-11-05"
	testDateEnd = "2018-11-15"
)

func TestHandleGetExchange(t *testing.T) {

	testCases := []struct {
		wantCode    int
		wantBody    string
		name        string
		queryString string
	}{
		{
			name:        "All specified",
			wantCode:    200,
			wantBody:    "Can't mock it",
			queryString: fmt.Sprintf("?amount=%d&from=%s&to=%s", testAmount, testFrom, testTo),
		},
		{
			name:        "One not specified",
			wantCode:    400,
			wantBody:    "{\"error\":\"All query params must be specify\"}\n",
			queryString: fmt.Sprintf("?amount=&from=%s&to=%s", testFrom, testTo),
		},
		{
			name:        "Two not specified",
			wantCode:    400,
			wantBody:    "{\"error\":\"All query params must be specify\"}\n",
			queryString: fmt.Sprintf("?amount=&from=&to=%s", testTo),
		},
		{
			name:        "All not specified",
			wantCode:    400,
			wantBody:    "{\"error\":\"All query params must be specify\"}\n",
			queryString: fmt.Sprintf("?amount=&from=&to="),
		},
		{
			name:        "Not specify query params at all",
			wantCode:    400,
			wantBody:    "{\"error\":\"All query params must be specify\"}\n",
			queryString: "",
		},
	}

	s := New(NewConfig())

	handler := http.HandlerFunc(s.handleGetExchange())

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", fmt.Sprintf("/api/exchange%s", testCase.queryString), nil)

			handler.ServeHTTP(rec, req)

			if testCase.name != "All specified" {
				assert.Equal(t, testCase.wantBody, rec.Body.String())
			}

			assert.Equal(t, testCase.wantCode, rec.Code)
		})
	}

}

func TestHandleGetExchangeAt(t *testing.T) {
	testCases := []struct {
		wantCode    int
		wantBody    string
		name        string
		queryString string
	}{
		{
			name:        "All specified",
			wantCode:    200,
			wantBody:    "Can't mock it",
			queryString: fmt.Sprintf("?amount=%d&from=%s&to=%s&date=%s", testAmount, testFrom, testTo, testDate),
		},
		{
			name:        "One not specified",
			wantCode:    400,
			wantBody:    "{\"error\":\"All query params must be specify\"}\n",
			queryString: fmt.Sprintf("?amount=&from=%s&to=%s&date=%s", testFrom, testTo, testDate),
		},
		{
			name:        "Two not specified",
			wantCode:    400,
			wantBody:    "{\"error\":\"All query params must be specify\"}\n",
			queryString: fmt.Sprintf("?amount=&from=&to=%s&date=%s", testTo, testDate),
		},
		{
			name:        "Three not specified",
			wantCode:    400,
			wantBody:    "{\"error\":\"All query params must be specify\"}\n",
			queryString: fmt.Sprintf("?amount=&from=&to=&date=%s", testDate),
		},
		{
			name:        "All not specified",
			wantCode:    400,
			wantBody:    "{\"error\":\"All query params must be specify\"}\n",
			queryString: fmt.Sprintf("?amount=&from=&to=&date="),
		},
		{
			name:        "Not specify query params at all",
			wantCode:    400,
			wantBody:    "{\"error\":\"All query params must be specify\"}\n",
			queryString: "",
		},
	}

	s := New(NewConfig())

	handler := http.HandlerFunc(s.handleGetExchange())

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", fmt.Sprintf("/api/exchange_at%s", testCase.queryString), nil)

			handler.ServeHTTP(rec, req)

			if testCase.name != "All specified" {
				assert.Equal(t, testCase.wantBody, rec.Body.String())
			}

			assert.Equal(t, testCase.wantCode, rec.Code)
		})
	}

}

func TestGetCurrenciesRange(t *testing.T) {
	testCases := []struct {
		name      string
		want      []*big.Float
		from      string
		to        string
		startDate string
		endDate   string
	}{
		{
			name:      "All specified",
			want:      []*big.Float{},
			from:      testFrom,
			to:        testTo,
			startDate: testDate,
			endDate:   testDateEnd,
		},
		{
			name:      "All specified wrong",
			want:      nil,
			from:      "XYZ",
			to:        "WYL",
			startDate: testDateEnd,
			endDate:   testDate,
		},
		{
			name:      "End date less than start date",
			want:      nil,
			from:      testFrom,
			to:        testTo,
			startDate: testDateEnd,
			endDate:   testDate,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			sl, _ := getCurrenciesRange(testCase.from, testCase.to, testCase.startDate, testCase.endDate)
			assert.IsType(t, testCase.want, sl)

		})
	}
}

func TestHandleGetExchangeBetween(t *testing.T) {
	testCases := []struct {
		wantCode    int
		wantBody    string
		name        string
		queryString string
	}{
		{
			name:        "All specified",
			wantCode:    200,
			wantBody:    "Can't mock it",
			queryString: fmt.Sprintf("?from=%s&to=%s&start_date=%s&end_date=%s", testFrom, testTo, testDate, testDateEnd),
		},
		{
			name:        "One not specified",
			wantCode:    400,
			wantBody:    "{\"error\":\"All query params must be specify\"}\n",
			queryString: fmt.Sprintf("?from=&to=%s&start_date=%s&end_date=%s", testTo, testDate, testDateEnd),
		},
		{
			name:        "Two not specified",
			wantCode:    400,
			wantBody:    "{\"error\":\"All query params must be specify\"}\n",
			queryString: fmt.Sprintf("?from=&to=&start_date=%s&end_date=%s", testDate, testDateEnd),
		},
		{
			name:        "Three not specified",
			wantCode:    400,
			wantBody:    "{\"error\":\"All query params must be specify\"}\n",
			queryString: fmt.Sprintf("?from=&to=&start_date=&end_date=%s", testDateEnd),
		},
		{
			name:        "All not specified",
			wantCode:    400,
			wantBody:    "{\"error\":\"All query params must be specify\"}\n",
			queryString: fmt.Sprintf("?from=&to=&start_date=&end_date="),
		},
		{
			name:        "Not specify query params at all",
			wantCode:    400,
			wantBody:    "{\"error\":\"All query params must be specify\"}\n",
			queryString: "",
		},
	}

	s := New(NewConfig())

	handler := http.HandlerFunc(s.handleGetExchange())

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", fmt.Sprintf("/api/exchange_between%s", testCase.queryString), nil)

			handler.ServeHTTP(rec, req)

			if testCase.name != "All specified" {
				assert.Equal(t, testCase.wantBody, rec.Body.String())
				assert.Equal(t, testCase.wantCode, rec.Code)
			}

		})
	}

}
