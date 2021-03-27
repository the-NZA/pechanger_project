package pechanger

import (
	"errors"
	"math/big"

	"github.com/asvvvad/exchange"
)

// Author: Roman Kozlov
func getCurrenciesRange(fromCur, toCur, startDate, endDate string) ([]*big.Float, error) {
	// create new exchange object
	ex := exchange.New(fromCur)
	if ex == nil {
		return nil, errors.New("Can't create exchange object")
	}

	ex.CacheEnabled = false

	// convert from startDate to end_end for 2 currencies (from and to)
	// res, err := ex.TimeseriesMultiple(startDate, endDate, []string{toCur})
	res, err := ex.TimeseriesSingle(startDate, endDate, toCur)
	if err != nil {
		return nil, err
	}

	// Fill slice with all values
	resSlice := make([]*big.Float, 0)
	for _, m := range res {
		for _, v := range m {
			resSlice = append(resSlice, v)
		}
	}

	return resSlice, nil
}
