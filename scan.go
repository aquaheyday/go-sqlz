package sqlz

import (
	"database/sql"
	"errors"
	"reflect"
)

func scanStructs(rows *sql.Rows, elemType reflect.Type) (reflect.Value, error) {
	columns, err := rows.Columns()
	if err != nil {
		return reflect.Value{}, err
	}

	isPtr := false
	realType := elemType

	if elemType.Kind() == reflect.Ptr && elemType.Elem().Kind() == reflect.Struct {
		isPtr = true
		realType = elemType.Elem()
	} else if elemType.Kind() != reflect.Struct {
		return reflect.Value{}, errors.New("element type must be struct or pointer to struct")
	}

	result := reflect.MakeSlice(reflect.SliceOf(elemType), 0, 0)
	fieldMap := buildFieldMap(realType)

	for rows.Next() {
		elemPtr := reflect.New(realType)
		elem := elemPtr.Elem()

		scanTargets := make([]interface{}, len(columns))
		for i, col := range columns {
			if idx, ok := fieldMap[col]; ok {
				field := elem.Field(idx)
				if field.CanSet() && field.Addr().CanInterface() {
					scanTargets[i] = field.Addr().Interface()
					continue
				}
			}
			var ignored sql.RawBytes
			scanTargets[i] = &ignored
		}

		if err := rows.Scan(scanTargets...); err != nil {
			return reflect.Value{}, err
		}

		// If the field is a pointer type ([]*T), store the pointer itself.
		if isPtr {
			result = reflect.Append(result, elemPtr)
		} else {
			result = reflect.Append(result, elem)
		}
	}

	if err := rows.Err(); err != nil {
		return reflect.Value{}, err
	}

	return result, nil
}
