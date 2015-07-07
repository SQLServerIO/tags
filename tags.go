package tags

import (
	"fmt"
	"reflect"
	"strings"
)

// Tag returns the value of the "tagName" tag for the given struct field name of
// the "data" struct, stripping any tag modifiers such as "omitempty".
// Returns the empty string if there is no tag, and an error if the field
// does not exist in the struct.
func Tag(data interface{}, tagName, fieldName string) (string, error) {
	dataType := reflect.TypeOf(data)
	if dataType.Kind() != reflect.Struct {
		return "", fmt.Errorf("must pass in a struct data type")
	}
	field, found := dataType.FieldByName(fieldName)
	if !found {
		return "", fmt.Errorf("struct does not have a field %v", fieldName)
	}
	tag := field.Tag.Get(tagName)

	// NOTE: this stops us from being able to use commas in the bson field names
	// of our models
	if index := strings.Index(tag, ","); index != -1 {
		tag = tag[:index]
	}
	return tag, nil
}

// MustHaveTag gets the "tagName" struct tag for a field, panicking if
// either the field does not exist or has no tag.
func MustHaveTag(data interface{}, tagName, fieldName string) string {
	tagValue, err := Tag(data, tagName, fieldName)
	if err != nil {
		panic(fmt.Sprintf("error getting %v tag: %v", tagName, err))
	}
	if tagValue == "" {
		panic(fmt.Sprintf("field %v cannot have an empty %v tag", fieldName, tagName))
	}
	return tagValue
}
