package main

import (
	"errors"
	"golang-test-id-token/logic"
	"io"
	"log"
	"net/url"
)

func main() {}

func init() {
	log.Printf("[%v] [%v] loaded", ModifierRegisterer, ModifierRegisterer)
}

// ModifierRegisterer is the symbol the plugin loader will be looking for. It must
// implement the plugin.Registerer interface
// https://github.com/luraproject/lura/blob/master/proxy/plugin/modifier.go#L71
var ModifierRegisterer = registerer("cloud-run-service-account")

type registerer string

// RegisterModifiers is the function the plugin loader will call to register the
// modifier(s) contained in the plugin using the function passed as argument.
// f will register the factoryFunc under the name and mark it as a request
// and/or response modifier.
func (r registerer) RegisterModifiers(f func(
	name string,
	factoryFunc func(map[string]interface{}) func(interface{}) (interface{}, error),
	appliesToRequest bool,
	appliesToResponse bool,
)) {
	f(string(r), r.requestDump, true, false)
	log.Printf("[%v] Registered cloud-run-service-account", ModifierRegisterer)
}

var unknownTypeError = errors.New("unknown request type")

func (r registerer) requestDump(
	cfg map[string]interface{},
) func(interface{}) (interface{}, error) {
	// return the modifier
	log.Printf("[%v] Request dumper injected", ModifierRegisterer)
	return func(input interface{}) (interface{}, error) {
		req, ok := input.(RequestWrapper)
		if !ok {
			return nil, unknownTypeError
		}

		r := modifier(req)

		return r, nil
	}
}

func modifier(req RequestWrapper) requestWrapper {
	headers := req.Headers()

	log.Printf("[%v] Adding Google Service Account", ModifierRegisterer)
	headers["Authorization"] = append(headers["Authorization"], logic.GoogleId(req.URL().Scheme+"://"+req.URL().Host))
	log.Printf("[%v] Added Google Service Account", ModifierRegisterer)

	return requestWrapper{
		params:  req.Params(),
		headers: headers,
		body:    req.Body(),
		method:  req.Method(),
		url:     req.URL(),
		query:   req.Query(),
		path:    req.Path(),
	}
}

type RequestWrapper interface {
	Params() map[string]string
	Headers() map[string][]string
	Body() io.ReadCloser
	Method() string
	URL() *url.URL
	Query() url.Values
	Path() string
}

type requestWrapper struct {
	method  string
	url     *url.URL
	query   url.Values
	path    string
	body    io.ReadCloser
	params  map[string]string
	headers map[string][]string
}
