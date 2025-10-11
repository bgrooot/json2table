package config

type NilBool struct {
	Value *bool
}

func CreateNilBool(val bool) NilBool {
	return NilBool{Value: &val}

}
func (n NilBool) IsSet() bool {
	return n.Value != nil
}

func (n NilBool) IsTrue() bool {
	return *n.Value
}
