package pechanger

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"net/mail"
	"net/smtp"
	"net/textproto"
	"strconv"
	"time"

	"github.com/asvvvad/exchange"
	"github.com/jordan-wright/email"
)

const defString = "NOT IMPLEMENTED"

var (
	errBadQueryRequest = errors.New("All query params must be specify")
	errBadBodyRequest  = errors.New("All body must be specify with right data")
	errBadEmail        = errors.New("Specify a correct email address")
	errInternalServer  = errors.New("Internal server error")
	errMinYear2006     = errors.New("The minimum year can only be 2006")
	errWrongDateRange  = errors.New("Start date must be less than end date")
)

// Author: Roman Kozlov
func (s *Server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}

}

// Author: Roman Kozlov
func (s *Server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

// Author: Roman Kozlov
func (s *Server) handleGetApi() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.logger.Logf("INFO Request from:%s\n", r.Host)
		// s.logger.Logf("INFO %v", r.URL.Query())

		s.respond(w, r, http.StatusOK, defString)
	}
}

// Author: Sergey Shakotko
// Handle basic converstion
func (s *Server) handleGetExchange() http.HandlerFunc {
	type respose struct {
		Result *big.Float `json:"res"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		amountStr := r.URL.Query()["amount"][0]
		if amountStr == "" {
			s.error(w, r, http.StatusBadRequest, errBadQueryRequest)
			return
		}

		from := r.URL.Query()["from"][0]
		if from == "" {
			s.error(w, r, http.StatusBadRequest, errBadQueryRequest)
			return
		}

		to := r.URL.Query()["to"][0]
		if to == "" {
			s.error(w, r, http.StatusBadRequest, errBadQueryRequest)
			return
		}

		// create new exchange object
		ex := exchange.New(from)
		if ex == nil {
			s.error(w, r, http.StatusInternalServerError, errInternalServer)
			return
		}

		amount, err := strconv.Atoi(amountStr)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, errInternalServer)
			return
		}

		// convert 'from' to specific 'to'
		res, err := ex.ConvertTo(to, amount)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, errInternalServer)
			return
		}

		s.respond(w, r, http.StatusOK, &respose{Result: res})
	}
}

// Author: Sergey Shakotko
// Handle specific date converstion
func (s *Server) handleGetExchangeAt() http.HandlerFunc {
	type respose struct {
		Result *big.Float `json:"res"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		amountStr := r.URL.Query()["amount"][0]
		if amountStr == "" {
			s.error(w, r, http.StatusBadRequest, errBadQueryRequest)
			return
		}

		from := r.URL.Query()["from"][0]
		if from == "" {
			s.error(w, r, http.StatusBadRequest, errBadQueryRequest)
			return
		}

		to := r.URL.Query()["to"][0]
		if to == "" {
			s.error(w, r, http.StatusBadRequest, errBadQueryRequest)
			return
		}

		// date must be in the format YYYY-MM-DD
		date := r.URL.Query()["date"][0]
		if date == "" {
			s.error(w, r, http.StatusBadRequest, errBadQueryRequest)
			return
		}

		tm, err := time.Parse("2006-01-02", date)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, errInternalServer)
			return
		}

		// API's min available year is 2006
		if tm.Year() < 2006 {
			s.error(w, r, http.StatusInternalServerError, errMinYear2006)
			return
		}

		// create new exchange object
		ex := exchange.New(from)
		if ex == nil {
			s.error(w, r, http.StatusInternalServerError, errInternalServer)
			return
		}

		amount, err := strconv.Atoi(amountStr)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, errInternalServer)
			return
		}

		// convert at specific date
		res, err := ex.ConvertAt(date, to, amount)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, errInternalServer)
			return
		}

		s.respond(w, r, http.StatusOK, &respose{Result: res})
	}
}

// Author: Roman Kozlov
// Handle get currencies between to dates
func (s *Server) handleGetExchangeBetween() http.HandlerFunc {
	type respose struct {
		Result []*big.Float `json:"res"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		from := r.URL.Query()["from"][0]
		if from == "" {
			s.error(w, r, http.StatusBadRequest, errBadQueryRequest)
			return
		}

		to := r.URL.Query()["to"][0]
		if to == "" {
			s.error(w, r, http.StatusBadRequest, errBadQueryRequest)
			return
		}

		// Start Date
		startDate := r.URL.Query()["start_date"][0]
		if startDate == "" {
			s.error(w, r, http.StatusBadRequest, errBadQueryRequest)
			return
		}

		// End Date
		endDate := r.URL.Query()["end_date"][0]
		if endDate == "" {
			s.error(w, r, http.StatusBadRequest, errBadQueryRequest)
			return
		}

		// Time object from
		tmFrom, err := time.Parse("2006-01-02", startDate)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, errInternalServer)
			return
		}

		if tmFrom.Year() < 2006 {
			s.error(w, r, http.StatusInternalServerError, errMinYear2006)
			return
		}

		// Time object to
		tmTo, err := time.Parse("2006-01-02", endDate)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, errInternalServer)
			return
		}

		// API's min available year is 2006
		if tmTo.Year() < 2006 {
			s.error(w, r, http.StatusInternalServerError, errMinYear2006)
			return
		}

		// Check for valid range
		if tmFrom.Unix() > tmTo.Unix() {
			s.error(w, r, http.StatusInternalServerError, errWrongDateRange)
			return

		}

		// get slice of currencies between chosen dates
		resSlice, err := getCurrenciesRange(from, to, startDate, endDate)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, errInternalServer)
			return
		}

		s.respond(w, r, http.StatusOK, &respose{Result: resSlice})
	}
}

// Author: Sergey Shakotko
// Handle share by email
func (s *Server) handlePostEmailShare() http.HandlerFunc {
	type respose struct {
		Result string `json:"res"`
	}

	type credentials struct {
		AmountFrom string `json:"amount_from"`
		AmountTo   string `json:"amount_to"`
		FromCur    string `json:"from"`
		ToCur      string `json:"to"`
		Email      string `json:"email"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Body == nil {
			s.error(w, r, http.StatusBadRequest, errBadBodyRequest)
			return
		}

		var cred credentials

		// Decode request body
		err := json.NewDecoder(r.Body).Decode(&cred)
		if err != nil {
			s.logger.Logf("ERROR %v", err)
			s.error(w, r, http.StatusInternalServerError, errInternalServer)
			return
		}

		// Simple validation
		if cred.AmountFrom == "" || cred.AmountTo == "" || cred.Email == "" || cred.FromCur == "" || cred.ToCur == "" {
			s.error(w, r, http.StatusBadRequest, errBadBodyRequest)
			return
		}

		// Parse email with default validation
		addr, err := mail.ParseAddress(cred.Email)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, errBadEmail)
			return
		}

		e := &email.Email{
			To:      []string{addr.Address},
			From:    fmt.Sprintf("Pechanger <%s>", s.config.EmailFrom),
			Subject: fmt.Sprintf("New exchange was complete"),
			HTML:    []byte(fmt.Sprintf("<h1>Your last exchange was: </h1><h2>%s %s to %s %s</h2>", cred.AmountFrom, cred.FromCur, cred.AmountTo, cred.ToCur)),
			Headers: textproto.MIMEHeader{},
		}

		// Create tls.config (required by mailer)
		tlsconf := &tls.Config{
			InsecureSkipVerify: true,
			ServerName:         s.config.GetAddress(),
		}

		// Send Email
		err = e.SendWithStartTLS(s.config.GetAddress(), smtp.PlainAuth("", s.config.EmailFrom, s.config.EmailPass, s.config.EmailSMTP), tlsconf)
		if err != nil {
			s.logger.Logf("ERROR %v\n", err)
			s.error(w, r, http.StatusInternalServerError, errInternalServer)
			return
		}

		s.respond(w, r, http.StatusOK, &respose{Result: "Email successfully send"})
	}
}
