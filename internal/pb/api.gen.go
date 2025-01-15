//go:build go1.22

// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/oapi-codegen/runtime"
)

// AddCurrencyRequest defines model for AddCurrencyRequest.
type AddCurrencyRequest struct {
	// Coin The cryptocurrency name (e.g., BTC, ETH)
	Coin string `json:"coin"`

	// Interval The interval (in seconds) for periodic price collection
	Interval int `json:"interval"`
}

// AddCurrencyResponse defines model for AddCurrencyResponse.
type AddCurrencyResponse struct {
	// Coin The cryptocurrency name added
	Coin *string `json:"coin,omitempty"`

	// Interval The interval at which the cryptocurrency price will be fetched
	Interval *int    `json:"interval,omitempty"`
	Message  *string `json:"message,omitempty"`
}

// PriceResponse defines model for PriceResponse.
type PriceResponse struct {
	// Coin The cryptocurrency name
	Coin string `json:"coin"`

	// Price The price of the cryptocurrency at the specified timestamp
	Price float32 `json:"price"`

	// Timestamp The timestamp of the price data
	Timestamp int `json:"timestamp"`
}

// RemoveCurrencyResponse defines model for RemoveCurrencyResponse.
type RemoveCurrencyResponse struct {
	Message *string `json:"message,omitempty"`
}

// GetCurrencyPriceParams defines parameters for GetCurrencyPrice.
type GetCurrencyPriceParams struct {
	Coin      string `form:"coin" json:"coin"`
	Timestamp int    `form:"timestamp" json:"timestamp"`
}

// DeleteCurrencyRemoveParams defines parameters for DeleteCurrencyRemove.
type DeleteCurrencyRemoveParams struct {
	Coin string `form:"coin" json:"coin"`
}

// PostCurrencyAddJSONRequestBody defines body for PostCurrencyAdd for application/json ContentType.
type PostCurrencyAddJSONRequestBody = AddCurrencyRequest

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Add a cryptocurrency to the watch list
	// (POST /currency/add)
	PostCurrencyAdd(w http.ResponseWriter, r *http.Request)
	// Get the price of a cryptocurrency at a specific timestamp
	// (GET /currency/price)
	GetCurrencyPrice(w http.ResponseWriter, r *http.Request, params GetCurrencyPriceParams)
	// Remove a cryptocurrency from the watch list
	// (DELETE /currency/remove)
	DeleteCurrencyRemove(w http.ResponseWriter, r *http.Request, params DeleteCurrencyRemoveParams)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// PostCurrencyAdd operation middleware
func (siw *ServerInterfaceWrapper) PostCurrencyAdd(w http.ResponseWriter, r *http.Request) {

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostCurrencyAdd(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetCurrencyPrice operation middleware
func (siw *ServerInterfaceWrapper) GetCurrencyPrice(w http.ResponseWriter, r *http.Request) {

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetCurrencyPriceParams

	// ------------- Required query parameter "coin" -------------

	if paramValue := r.URL.Query().Get("coin"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "coin"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "coin", r.URL.Query(), &params.Coin)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "coin", Err: err})
		return
	}

	// ------------- Required query parameter "timestamp" -------------

	if paramValue := r.URL.Query().Get("timestamp"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "timestamp"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "timestamp", r.URL.Query(), &params.Timestamp)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "timestamp", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetCurrencyPrice(w, r, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// DeleteCurrencyRemove operation middleware
func (siw *ServerInterfaceWrapper) DeleteCurrencyRemove(w http.ResponseWriter, r *http.Request) {

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params DeleteCurrencyRemoveParams

	// ------------- Required query parameter "coin" -------------

	if paramValue := r.URL.Query().Get("coin"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "coin"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "coin", r.URL.Query(), &params.Coin)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "coin", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteCurrencyRemove(w, r, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, StdHTTPServerOptions{})
}

// ServeMux is an abstraction of http.ServeMux.
type ServeMux interface {
	HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type StdHTTPServerOptions struct {
	BaseURL          string
	BaseRouter       ServeMux
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, m ServeMux) http.Handler {
	return HandlerWithOptions(si, StdHTTPServerOptions{
		BaseRouter: m,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, m ServeMux, baseURL string) http.Handler {
	return HandlerWithOptions(si, StdHTTPServerOptions{
		BaseURL:    baseURL,
		BaseRouter: m,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options StdHTTPServerOptions) http.Handler {
	m := options.BaseRouter

	if m == nil {
		m = http.NewServeMux()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	m.HandleFunc("POST "+options.BaseURL+"/currency/add", wrapper.PostCurrencyAdd)
	m.HandleFunc("GET "+options.BaseURL+"/currency/price", wrapper.GetCurrencyPrice)
	m.HandleFunc("DELETE "+options.BaseURL+"/currency/remove", wrapper.DeleteCurrencyRemove)

	return m
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/7RUTW/bOBD9K8TsHhJAsL27N92SbJH21CDILciBJkc2A4lkhiOnRqD/XpD0h1SpdYrU",
	"Jwkk5+u9N+8NlGu8s2g5QPkGQa2xken3Suublgit2t7jS4uB46kn55HYYHqjnLHxqzEoMp6Ns1DCwxqF",
	"oq1np3YJhJUNigucrWaFuH64KcSnh8+XUABvPUIJgcnYFXQFGMtIG1lPZ93figtjRUDlrA6XonIkPJJx",
	"2ijhySgUytU1qhR5KBKDV0jQdQUQvrSGUEP5mIfoVX46RLjlMyqObQ3ACN7ZgB9FQ2qN+gMQSBava6PW",
	"gscFMgivpq7FEkWFrNb9WgckCmgwBLlKw+A32fg63t8Ms6VORWiVwhCqtq6347a7CdDuYhd/CK4poNKU",
	"0xkyAK6aAkdyOg0elakMasGmwcCy8VBA5aiRDCVUtZN8rGrbZpkRO76erHy43lfPrWjJ8t1SzIP1S01p",
	"8h4bt8HTsnwPxZRy/T7JXdJr5WJyNtxLLb4uA9IGCQrYIIWM0D+zxWwRm3cerfQGSvgvHRXgJa9Tv/N9",
	"V3OpdRrHZe+JQ8kI9RcNJdy5wPvhr3RUN2WbunZ6m/VlGW2KlN7XRqXY+XNw9mh18e9vwgpK+Gt+9ML5",
	"zgjnEy7YDVljajEdZPzTCP8uFufpYMdxamGovdM7G2NC2zSStlBGRxPyx9VglyT7KlmtRW3SsEWPkMPG",
	"rXCCkVs8EHK3U7CXJBtkpADl4xvEhYeXFimKK611udf8ENCiB85IhNN5+lt8MtlxAZ/OSN3QACdISw8E",
	"IZPBza/5ukXu2YmrxuRJFnLvaqpnakMK86Zn76qRcUzj/+n8qLn0/jxUnhP8n9jj6dWZ9MIhGTn3mIKK",
	"XDPeoK77HgAA//8XBee+ZwkAAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
