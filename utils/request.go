package utils

import (
	"errors"
	"reflect"
)

type RequestDataValidator struct {
	Null 	bool
	Type 	reflect.Kind
	Max	 	int
	Min 	int
	Email 	bool
}

func RequestValidator(field string, dataTemp map[string]interface{}, validator RequestDataValidator) (interface{}, error) {
	data := dataTemp[field]

	// null validation
	if !validator.Null && data == nil {
		return nil, errors.New("field " + field + " harus ada")
	}else if validator.Null && data == nil{
		return nil, nil
	}

	// data type validation
	if validator.Type == reflect.Int{
		validator.Type = reflect.Float64
	}
	if reflect.TypeOf(data).Kind() != validator.Type{
		return nil, errors.New("field " + field + " harusnya bertipe " + validator.Type.String())
	}



	return data, nil
}
