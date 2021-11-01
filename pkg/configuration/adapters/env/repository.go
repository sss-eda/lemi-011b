package env

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/sss-eda/lemi-011b/pkg/configuration"
)

type repository struct{}

// NewRepository TODO
func NewRepository() (configuration.Repository, error) {
	return &repository{}, nil
}

func (repo *repository) Configure(
	ctx context.Context,
	config interface{},
) error {
	return nil
}

// ParseENV TODO
func ParseENV(config interface{}) error {
	return parse(config, "_")
}

func parse(
	configStruct interface{},
	seperator string,
	prefixes ...string,
) error {
	if reflect.TypeOf(configStruct).Kind() != reflect.Ptr {
		return fmt.Errorf(
			"environment configuration parse function requires configStruct "+
				"parameter of type pointer, got parameter of type: %v",
			reflect.TypeOf(configStruct).Kind(),
		)
	}

	// fmt.Println(strings.Join(prefixes, seperator))

	val := reflect.ValueOf(configStruct).Elem()
	typ := val.Type()

	// fmt.Printf(
	//     "Type: %v, Kind: %v\n",
	//     reflect.TypeOf(configStruct),
	//     reflect.TypeOf(configStruct).Kind(),
	// )

	switch typ.Kind() {
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			f := val.Field(i).Addr().Interface()

			// fmt.Printf(
			//     "Type: %v, Kind: %v\n",
			//     val.Field(i).Type(),
			//     val.Field(i).Type().Kind(),
			// )

			if tag, ok := typ.Field(i).Tag.Lookup("env"); ok {
				var err error = nil
				if tag == "" {
					err = parse(f, seperator, prefixes...)
				} else {
					if len(prefixes) > 0 {
						err = parse(f, seperator, append(prefixes, tag)...)
					} else {
						err = parse(f, seperator, tag)
					}
				}
				if err != nil {
					return err
				}
			}
		}
	case reflect.Float32, reflect.Float64:
		env := os.Getenv(strings.Join(prefixes, seperator))
		v, err := strconv.ParseFloat(env, 64)
		if err != nil {
			return fmt.Errorf(
				"unable to parse environment variable: %s, with value: %s, "+
					"to float, with error: %v",
				strings.Join(prefixes, seperator),
				env,
				err,
			)
		}
		val.SetFloat(v)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Int64:
		env := os.Getenv(strings.Join(prefixes, seperator))
		v, err := strconv.ParseInt(env, 10, 64)
		if err != nil {
			return fmt.Errorf(
				"unable to parse environment variable: %s, with value: %s, "+
					"to int, with error: %v",
				strings.Join(prefixes, seperator),
				env,
				err,
			)
		}
		val.SetInt(v)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32,
		reflect.Uint64:
		env := os.Getenv(strings.Join(prefixes, seperator))
		v, err := strconv.ParseUint(env, 10, 64)
		if err != nil {
			return fmt.Errorf(
				"unable to parse environment variable: %s, with value: %s, "+
					"to uint, with error: %v",
				strings.Join(prefixes, seperator),
				env,
				err,
			)
		}
		val.SetUint(v)
	case reflect.String:
		env := os.Getenv(strings.Join(prefixes, seperator))
		val.SetString(env)
	default:
		return fmt.Errorf(
			"configuration struct contains unsupported type: %q",
			typ.Kind().String(),
		)
	}

	return nil
}
