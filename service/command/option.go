package command

import (
	"fmt"
)

type Option struct {
	data map[string]interface{}
}

func (o *Option) set(name string, value interface{}) *Option {
	if o.data == nil {
		o.data = make(map[string]interface{})
	}

	o.data[name] = value

	return o
}

func (o *Option) Get(name string) (interface{}, error) {
	if o.data[name] == nil {
		return nil, fmt.Errorf(`Option name %s does not exists.`, name)
	}

	return o.data[name], nil
}
