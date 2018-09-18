package service

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

type opFunc func(a, b float64) float64

type handlerBinaryOp struct {
	a, b float64
	op   opFunc
}

func NewHandlerBinaryOp(op opFunc) *handlerBinaryOp {
	return &handlerBinaryOp{op: op}
}

func (h *handlerBinaryOp) parseRequest(r *http.Request) error {
	q := r.URL.Query()
	a, b := q["a"], q["b"]

	if len(a) < 1 || len(b) < 1 {
		return errors.New("Need 'a' and 'b' parameters")
	}
	var aErr, bErr error
	h.a, aErr = strconv.ParseFloat(a[0], 64)
	h.b, bErr = strconv.ParseFloat(b[0], 64)
	if aErr != nil || bErr != nil {
		return errors.New("Parameter parse error")
	}
	return nil
}

func (h handlerBinaryOp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h.parseRequest(r); err != nil {
		w.WriteHeader(400)
		fmt.Fprintln(w, err.Error())
	}
	result := h.op(h.a, h.b)
	fmt.Fprintf(w, strconv.FormatFloat(result, 'f', -1, 64)+"\n")
}
