package chain_of_respinsibility_pattern

import "testing"

func TestCoR(t *testing.T) {
	request := Request{
		Header: map[string]string {
			"Content-Type": "application/json",
		},
		Body: "My request data",
	}

	printerFilter 	:= &PrinterFilter{}
	unzipFilter 	:= &RequestUncompresserFilter{printerFilter}
	tracerFilter 	:= &TracerSetterFilter{unzipFilter}
	authFilter 		:= &AuthRequestFilter{tracerFilter}
	loggerFilter 	:= &RequestLoggerFilter{authFilter}


	loggerFilter.Filter(request)
}
