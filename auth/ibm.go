package auth

import(
  "encoding/json"
  "io/ioutil"
  "path/filepath"
  "log"

  // IBM/go-sdk-core
  "github.com/IBM/go-sdk-core/core"
  // watson-developer-cloud/go-sdk
  "github.com/watson-developer-cloud/go-sdk/naturallanguageunderstandingv1"
)

// Credentials: contain folowing:
type Credentials struct {
  ApiKey			string `json:"key"`
  ServiceURL	string `json:"url"`
}

func naturallanguageunderstandinginit() {
  // Read {Root directory of Project}/const/credentials/ibm/naturallanguageunderstanding.json
  bytes, err := ioutil.ReadFile(filepath.Abs("../const/credentials/ibm/naturallanguageunderstanding.json"))
  if err != nil {
    log.Fatal("Couldn't read naturallanguageunderstanding.json; Did you put the file properly?")
  }

  // Decode credentials
  var credentials Credentials
  if err := json.Unmarshal(bytes, &credentials); err != nil {
    log.Fatal("Couldn't decode credentials; Something in JSON could be wrong")
  }

  // Initialize authenticator with an apikey
  authenticator = &core.IamAuthenticator {
    ApiKey: credentials.ApiKey, 
  }

  // Set options to authenticator
  options :=
  &naturallanguageunderstandingv1.NaturalLanguageUnderstandingV1Options {
    Authenticator = authenticator
  }

  // Get naturalLanguageUnderstanding
  naturalLanguageUnderstanding, naturalLanguageUnderstandingErr
  := naturallanguageunderstandingv1.NewNaturalLanguageUnderstandingV1(options)

  if naturalLanguageUnderstandingErr != nil {
    panic(naturalLanguageUnderstandingErr)
  }

  // Set naturalLanguageUnderstanding a service uri
  naturalLanguageUnderstanding.SetServiceURL(credentials.ServiceURL)
}

