package sqlz

import (
	"database/sql"
	"errors"
	"reflect"
)

// Select scans all rows from *sql.Rows into a slice of structs or struct pointers.
// It supports both []T and []*T.
func Select(rows *sql.Rows, dest interface{}) error {
	defer rows.Close()

	destVal := reflect.ValueOf(dest)
	if destVal.Kind() != reflect.Ptr || destVal.Elem().Kind() != reflect.Slice {
		return errors.New("dest must be pointer to a slice")
	}

	elemType := destVal.Elem().Type().Elem()
	result, err := scanStructs(rows, elemType)
	if err != nil {
		return err
	}

	destVal.Elem().Set(result)
	return nil
}
