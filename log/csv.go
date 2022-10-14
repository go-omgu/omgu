package log

import (
	"encoding/csv"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"strings"
)

type CSVLogger[T any] struct {
	w       *csv.Writer
	headers []*headerItem
}

type headerItem struct {
	tag   string
	field int
}

func NewCSV[T any](w io.Writer) *CSVLogger[T] {
	log := new(CSVLogger[T])
	log.w = csv.NewWriter(w)
	return log
}

func NewCSVLike[T any](w io.Writer, val T) (*CSVLogger[T], error) {
	log := NewCSV[T](w)
	h, err := digHeaders(val)
	if err != nil {
		return nil, err
	}
	log.headers = h
	log.printHeaders()
	return log, nil
}

func (log *CSVLogger[T]) printHeaders() error {
	hs := make([]string, 0, len(log.headers))
	for _, i := range log.headers {
		hs = append(hs, i.tag)
	}
	return log.w.Write(hs)
}

func (log *CSVLogger[T]) AppendMany(vals []T) error {
	if len(vals) == 0 {
		return nil
	}
	if len(log.headers) == 0 {
		h, err := digHeaders(vals[0])

		if err != nil {
			return err
		}

		log.headers = h
		log.printHeaders()
	}

	var err error
	for _, val := range vals {
		s, err := digValues(val, log.headers)
		if err != nil {
			return err
		}

		err = log.w.Write(s)
	}
	log.w.Flush()
	return err
}

func (log *CSVLogger[T]) Append(v T) error {
	return log.AppendMany([]T{v})
}

func digHeaders(item any) ([]*headerItem, error) {
	val := reflect.ValueOf(item)
	if val.Kind() == reflect.Pointer {
		val = reflect.Indirect(val)
	}
	tp := val.Type()
	headers := make([]*headerItem, tp.NumField())

	var count int
	for i := 0; i < tp.NumField(); i++ {
		f := tp.Field(i)
		t := f.Tag.Get("csv")
		if t == "" {
			continue
		}

		// Format: csv:"name,position"
		params := strings.Split(t, ",")
		if len(params) != 2 {
			return nil, fmt.Errorf("Please specific a position for \"%s\" field.", params[0])
		}

		pos, err := strconv.Atoi(params[1])
		if err != nil {
			return nil, err
		}

		headers[pos] = &headerItem{params[0], i}
		count++
	}

	return headers[:count], nil
}

func digValues(v any, headers []*headerItem) ([]string, error) {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Pointer {
		val = reflect.Indirect(val)
	}
	strs := make([]string, 0, len(headers))
	for _, k := range headers {
		strs = append(strs, stringify(val.Field(k.field).Interface()))
	}

	return strs, nil
}

func stringify(v any) string {
	switch v.(type) {
	case fmt.Stringer:
		return fmt.Sprintf("%s", v)
	case fmt.GoStringer:
		return fmt.Sprintf("%#v", v)
	default:
		return fmt.Sprintf("%v", v)
	}
}
