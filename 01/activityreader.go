package main

import (
	"encoding/csv"
	"io"
)

type ActivityReader struct {
	*csv.Reader
	fields []int
}

func NewActivityReader(r io.Reader, fields ...int) *ActivityReader {
	fr := &ActivityReader{
		Reader: csv.NewReader(r),
		fields: fields,
	}

	return fr
}

func (r *ActivityReader) Read() (record []string, err error) {
	rec, err := r.Reader.Read()
	if err != nil {
		return nil, err
	}

	record = make([]string, len(r.fields))
	for i, f := range r.fields {
		record[i] = rec[f]
	}

	return record, nil
}

func (r *ActivityReader) ReadAll() (records [][]string, err error) {
loop:
	for {
		rec, err := r.Read()
		switch err {
		case io.EOF:
			break loop
		case nil:
			records = append(records, rec)
		default:
			return nil, err
		}
	}

	return records, nil
}
