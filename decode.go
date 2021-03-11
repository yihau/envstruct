package envstruct

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

func FillIn(data interface{}) error {
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil
	}
	for i := 0; i < v.NumField(); i++ {
		envName := t.Field(i).Tag.Get("env")
		err := fillIn(v.Field(i), os.Getenv(envName))
		if err != nil {
			return fmt.Errorf(
				"fieldName: %v, env tag: %v, env value: %v, err: %v",
				t.Field(i).Name,
				envName,
				os.Getenv(envName),
				err,
			)
		}
	}
	return nil
}

func fillIn(v reflect.Value, envValue string) error {
	switch v.Kind() {
	case reflect.Bool:
		if envValue == "" {
			envValue = "false"
		}
		b, err := strconv.ParseBool(envValue)
		if err != nil {
			return err
		}
		v.SetBool(b)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if envValue == "" {
			envValue = "0"
		}
		i, err := strconv.ParseInt(envValue, 0, 64)
		if err != nil {
			return err
		}
		if v.OverflowInt(i) {
			return errors.New("overflow")
		}
		v.SetInt(i)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if envValue == "" {
			envValue = "0"
		}
		i, err := strconv.ParseUint(envValue, 0, 64)
		if err != nil {
			return err
		}
		if v.OverflowUint(i) {
			return errors.New("overflow")
		}
		v.SetUint(i)
	case reflect.Float32:
		if envValue == "" {
			envValue = "0"
		}
		f, err := strconv.ParseFloat(envValue, 32)
		if err != nil {
			return err
		}
		if v.OverflowFloat(f) {
			return errors.New("overflow")
		}
		v.SetFloat(f)
	case reflect.Float64:
		if envValue == "" {
			envValue = "0"
		}
		f, err := strconv.ParseFloat(envValue, 64)
		if err != nil {
			return err
		}
		if v.OverflowFloat(f) {
			return errors.New("overflow")
		}
		v.SetFloat(f)
	case reflect.String:
		v.SetString(envValue)
	case reflect.Ptr:
		if envValue == "" {
			return nil
		}
		v.Set(reflect.New(v.Type().Elem()))
		v = v.Elem()
		return fillIn(v, envValue)
	default:
		return fmt.Errorf("unsupported type: %v", v.Kind())
	}
	return nil
}
