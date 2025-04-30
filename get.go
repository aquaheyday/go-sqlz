package sqlz

import (
	"database/sql"
	"errors"
	"reflect"
)

// Get scans a single row into a struct.
// Returns sql.ErrNoRows if no row was found.
func Get(rows *sql.Rows, dest interface{}) error {
	defer rows.Close()

	destVal := reflect.ValueOf(dest)
	if destVal.Kind() != reflect.Ptr || destVal.Elem().Kind() != reflect.Struct {
		return errors.New("dest must be pointer to a struct")
	}

	elemType := destVal.Elem().Type()
	result, err := scanStructs(rows, elemType)
	if err != nil {
		return err
	}

	if result.Len() == 0 {
		return sql.ErrNoRows
	}

	destVal.Elem().Set(result.Index(0))
	return nil
}
