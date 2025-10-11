package config

import (
	"reflect"
)

type Config struct {
	Template      []string
	TemplateFiles []string
	TemplateData  any
	Orientation   string
	CSS           string
	CSSWebLink    string
	WrapWithHTML  NilBool
	DivideArray   NilBool
	TableClass    string
	TrClass       string
	ThClass       string
	TdClass       string
}

func defaultConfig() *Config {
	return &Config{
		WrapWithHTML: CreateNilBool(true),
		Orientation:  "vertical",
		DivideArray:  CreateNilBool(false),
	}
}

func ApplyDefaultConfig(config *Config) {
	defaultCfg := defaultConfig()

	if config == nil {
		config = defaultConfig()
	}

	valueOf := reflect.ValueOf(config).Elem()
	defaultTypeOf := reflect.TypeOf(*defaultCfg)
	defaultValueOf := reflect.ValueOf(*defaultCfg)

	for i := 0; i < defaultTypeOf.NumField(); i++ {
		defaultField := defaultTypeOf.Field(i)
		defaultVal := defaultValueOf.Field(i)
		val := valueOf.FieldByName(defaultField.Name)

		setDefaultByType(val, defaultVal)
	}
}

func setDefaultByType(val reflect.Value, defaultVal reflect.Value) {
	if val.CanSet() {
		switch val.Kind() {
		case reflect.Struct:
			if val.Type() == reflect.TypeOf(NilBool{}) && !val.Interface().(NilBool).IsSet() {
				val.Set(defaultVal)
			}

		case reflect.Int:
			if val.Int() == 0 {
				val.SetInt(defaultVal.Int())
			}

		default:
			if val.String() == "" {
				val.SetString(defaultVal.String())
			}
		}
	}
}
