package handler

import(
	"net/http"
	"log"
	"github.com/noor-core/service"
)

func NewDefaultHandler(
	defaultService service.DefaultService, 
) DefaultHandler {
	return &DefaultHandler{
		defaultService: defaultService, 
	}
}

type defaultHandler struct {
	defaultService service.DefaultService
}

func(handler *defaultHandler) ServeHTTP(
	writer http.ResponseWriter, 
	request *http.Request) {
	// Don't allow root
	if request.URL.Path != "/" {
		http.StatusBadRequest(writer, request)
		return
	}

	//Attempt to get result by targetUri
	response, err := handler.defaultService.GetResult(request.URL.Parse(targetUri))

	// Return if error
	if err != nil {
		log.Println("Something went wrong with handler!")
		log.Println(err)
		httpError(writer, response)
		return
	}

	// Return JSON otherwise
	responseJSON(writer, response)
	return
}

