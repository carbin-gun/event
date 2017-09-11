package event

import (
	"errors"
	"fmt"
	"reflect"
)

//New generate a new producer
func New() *Producer {
	return &Producer{
		subscriberContainer: make(map[string][]*subsciber),
	}
}

type subsciber struct {
	event     string
	fn        reflect.Value
	parametes []reflect.Type
}

func newSubscriber(event string, fn reflect.Value, inTypes []reflect.Type) *subsciber {
	return &subsciber{
		event:     event,
		fn:        fn,
		parametes: inTypes,
	}
}

// Producer which produces the events
type Producer struct {
	subscriberContainer map[string][]*subsciber
}

// AddListener add a listener for a name-specified event
func (p *Producer) AddListener(name string, callback interface{}) error {
	funcVal, parametersType, err := p.precheckListenerCallback(callback)
	if err != nil {
		return err
	}
	sub := newSubscriber(name, funcVal, parametersType)
	exists, ok := p.subscriberContainer[name]
	if !ok {
		exists = []*subsciber{sub}
	} else {
		exists = append(exists, sub)
	}
	p.subscriberContainer[name] = exists
	return nil
}

//Fire produce a name-specified event, wait to the subscriber to execute.
func (p *Producer) Fire(name string, params ...interface{}) {
	p.dispatchEvent(name, params...)
}

// AsyncFire make async call for the subscriber to execute.
func (p *Producer) AsyncFire(name string, params ...interface{}) {
	go p.dispatchEvent(name, params...)
}
func (p *Producer) dispatchEvent(name string, params ...interface{}) {
	subscribers, exist := p.subscriberContainer[name]
	if !exist {
		fmt.Println("no subscriber for event:", name)
		return
	}
	types, values := p.checkActualParameters(params...)
	hasSubscriber := false
	for _, sub := range subscribers {
		if p.typesMatchSubscriber(types, sub.parametes) {
			sub.fn.Call(values)
			hasSubscriber = true
		}
	}
	if !hasSubscriber {
		fmt.Printf("it seems that no subscriber for event [%s]\n", name)
	}
}

func (p *Producer) precheckListenerCallback(callback interface{}) (reflect.Value, []reflect.Type, error) {
	var funcVal reflect.Value
	if getType(callback) != reflect.Func {
		return funcVal, nil, errors.New("the callback should be a func")
	}
	funcVal = reflect.ValueOf(callback)
	types := p.getParameterTypes(funcVal)
	return funcVal, types, nil
}

func (p *Producer) getParameterTypes(callback reflect.Value) []reflect.Type {
	callbackType := callback.Type()
	parameterSize := callbackType.NumIn()
	if parameterSize == 0 {
		return nil
	}
	result := make([]reflect.Type, 0, parameterSize)
	for i := 0; i < parameterSize; i++ {
		result = append(result, callbackType.In(i))
	}
	return result
}

func (p *Producer) checkActualParameters(params ...interface{}) ([]reflect.Type, []reflect.Value) {
	if len(params) == 0 {
		return nil, nil
	}
	types := make([]reflect.Type, 0, len(params))
	values := make([]reflect.Value, 0, len(params))
	for _, param := range params {
		val := reflect.ValueOf(param)
		values = append(values, val)
		types = append(types, val.Type())
	}
	return types, values
}

func (p *Producer) typesMatchSubscriber(actual []reflect.Type, def []reflect.Type) bool {

	if actual == nil && def == nil {
		return true
	}

	if len(actual) != len(def) {
		return false
	}
	for i := range actual {
		if actual[i] != def[i] {
			return false
		}
	}
	return true
}

// return reflect the reflect.Kind Value
func getType(i interface{}) reflect.Kind {
	return reflect.TypeOf(i).Kind()
}
