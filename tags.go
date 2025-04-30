package sqlz

import "reflect"

// buildFieldMap returns a map of db tag (or field name) to field index.
// Unexported fields are skipped.
func buildFieldMap(t reflect.Type) map[string]int {
	m := make(map[string]int)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.PkgPath != "" {
			continue
		}

		tag := f.Tag.Get("db")
		if tag != "" {
			m[tag] = i
		} else {
			m[f.Name] = i
		}
	}
	return m
}
