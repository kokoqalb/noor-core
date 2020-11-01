package service

import(
  "encoding/json"
  "fmt"

  "github.com/kokoqalb/noor-core/repository"
  "github.com/kokoqalb/noor-core/auth"

  // IBM/go-sdk-core
  "github.com/IBM/go-sdk-core/core"

  // watson-developer-cloud/go-sdk
  "github.com/watson-developer-cloud/go-sdk/naturallanguageunderstandingv1"
)

// Initialize a default service
func NewDefaultService(
  defaultRepository repository.DefaultRepository
) DefaultService {
  return &defaultRepository {
    defaultRepository: defaultRepository, 
  }
}

type defaultService struct {
  defaultRepository repository.DefaultRepository
}

type request struct {
  Target string `json:"target"`
}

func(service *defaultService) GetResult(bytes []byte) {
  // kokomadakaketenai
  response, responseErr := naturalLanguageUnderstanding.Analyze(
    &naturallanguageunderstandingv1.AnalyzeOptions = options
  )
  if responseErr != nil panic(responseErr)

  result := json.MarshalIndent(response, "", " ")

  return result
}

