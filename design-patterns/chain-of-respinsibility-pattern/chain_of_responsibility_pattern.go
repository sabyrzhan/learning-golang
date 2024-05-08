package chain_of_respinsibility_pattern

import (
	"fmt"
	"math/rand"
	"strings"
)

/*
Here I created request filters.
Filtering starts from RequestLoggerFilter and is continued down to PrinterFilter.
 */

type Request struct {
	Header map[string]string
	Body string
}

func (r Request) String() string {
	q := func (repeat int) string { return strings.Repeat(" ", repeat) }
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("%sHeader data: \n", q(3)))
	for k, v := range r.Header {
		builder.WriteString(fmt.Sprintf("%s%s: %s\n", q(5), k, v))
	}
	builder.WriteString(fmt.Sprintf("%sBody data: \n", q(3)))
	builder.WriteString(fmt.Sprintf("%s%s", q(5), r.Body))

	return builder.String()
}

type RequestFilter interface {
	Filter(request Request)
}

type RequestLoggerFilter struct {
	next RequestFilter
}
type AuthRequestFilter struct {
	next RequestFilter
}
type TracerSetterFilter struct {
	next RequestFilter
}
type RequestUncompresserFilter struct {
	next RequestFilter
}
type ConverterToYAMLFilter struct {
	next RequestFilter
}

type PrinterFilter struct {
}

func (a *RequestLoggerFilter) Filter(request Request) {
	fmt.Println("RequestLoggerFilter: Received request with data: " + request.Body + " and header " + fmt.Sprintf("%v", request.Header))
	if a.next != nil {
		a.next.Filter(request)
	}
}

func (a *AuthRequestFilter) Filter(request Request) {
	fmt.Println("AuthRequestFilter: Authentication validated")

	if a.next != nil {
		a.next.Filter(request)
	}
}

func (a *TracerSetterFilter) Filter(request Request) {
	traceX := fmt.Sprintf("%d", rand.Intn(1000000000)+1000000000)
	fmt.Println("TracerSetterFilter: Setting X-Trace-Id: " + traceX)
	request.Header["X-Trace-Id"] = traceX

	if a.next != nil {
		a.next.Filter(request)
	}
}

func (a *RequestUncompresserFilter) Filter(request Request) {
	fmt.Println("RequestUncompresserFilter: Uncompressing the request")

	if a.next != nil {
		a.next.Filter(request)
	}
}

func (a *ConverterToYAMLFilter) Filter(request Request) {
	fmt.Println("ConverterToYAMLFilter: Converting request data to YAML file")

	if a.next != nil {
		a.next.Filter(request)
	}
}

func (a *PrinterFilter) Filter(request Request) {
	fmt.Println(fmt.Sprintf("PrinterFilter: Printing the final request data:\n%v", request))
}