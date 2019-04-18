package alexa

import (
  "fmt"
  "errors"
  "net/http"
  "io/ioutil"
  "encoding/json"

  "github.com/webability-go/alexa/request"
)

const (
  APITIMEZONE =         "/v2/devices/{deviceId}/settings/System.timeZone"
  APIDISTANCEUNIT =     "/v2/devices/{deviceId}/settings/System.distanceUnits"
  APITEMPERATUREUNIT =  "/v2/devices/{deviceId}/settings/System.temperatureUnit"

  APINAME =             "/v2/accounts/~current/settings/Profile.name"
  APIGIVENNAME =        "/v2/accounts/~current/settings/Profile.givenName"
  APIEMAIL =            "/v2/accounts/~current/settings/Profile.email"
  APIMOBILENUMBER =     "/v2/accounts/~current/settings/Profile.mobileNumber"
  
  APICOUNTRY =          "/v1/devices/{deviceId}/settings/countryAndPostalCode"
  APIADDRESS =          "/v1/devices/{deviceId}/settings/address"
)

func getAPIAccessToken(request *request.AlexaRequest) string {
  return request.Context.System.APIAccessToken
}

func getAPIEndPoint(request *request.AlexaRequest) string {
  return request.Context.System.APIEndPoint
}

func accessAPI(endPoint string, service string, token string) ([]byte, error) {

  hc := http.Client{}
  req, err := http.NewRequest("GET", endPoint + service, nil )
  if err != nil {
    return nil, err
  }
  req.Header.Set("content-type", "application/json")
  if token != "" {
    req.Header.Set("Authorization", "Bearer " + token)
  }
  resp, err := hc.Do(req)
  if err != nil {
    return nil, err
  }
  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }
  resp.Body.Close()
  
  fmt.Println(string(data))
  return data, nil
  
  aux := struct {
    Code string `json:"code"`
    Message string  `json:"message"`
    Data string  `json:"data"`
  }{}
  if err := json.Unmarshal(data, &aux); err != nil {
    // notify error
    return nil, err
  }
  if aux.Code != "" {
    return nil, errors.New(aux.Code + " " + aux.Message)
  }
  return nil, nil
  
//  return aux.Data, nil
}

func accessAccountAPI(request *request.AlexaRequest, service string) (map[string]interface{}, error) {
  accessToken := getAPIAccessToken(request);
  endPoint := getAPIEndPoint(request);
  if accessToken == "" || endPoint == "" {
    return nil, nil
  }
  accessAPI(endPoint, service, accessToken)
  return nil, nil
}

func accessDeviceAddressAPI(request *request.AlexaRequest, service string) (map[string]interface{}, error) {
  accessToken := getAPIAccessToken(request);
  endPoint := getAPIEndPoint(request);
  if accessToken == "" || endPoint == "" {
    return nil, nil
  }
  accessAPI(endPoint, service, accessToken)
  return nil, nil
}

func accessDeviceAPI(request *request.AlexaRequest, service string) (map[string]interface{}, error) {
  accessToken := getAPIAccessToken(request);
  endPoint := getAPIEndPoint(request);
  if accessToken == "" || endPoint == "" {
    return nil, nil
  }
  accessAPI(endPoint, service, accessToken)
  return nil, nil
}

/* ACCOUNT API */
func GetAccountFullName(request *request.AlexaRequest) (string, error) {
  accessAccountAPI(request, APINAME)
  return "", nil
}

func GetAccountGivenName(request *request.AlexaRequest) (string, error) {
  accessAccountAPI(request, APIGIVENNAME)
  return "", nil
}

func GetAccountEmail(request *request.AlexaRequest) (string, error) {
  accessAccountAPI(request, APIEMAIL)
  return "", nil
}

func GetAccountMobileNumber(request *request.AlexaRequest) (string, error) {
  accessAccountAPI(request, APIMOBILENUMBER)
  return "", nil
}

/* DEVICE ADDRESS AND COUNTRY */
func GetDeviceCountry(request *request.AlexaRequest) (string, error) {
  accessDeviceAddressAPI(request, APICOUNTRY)
  return "", nil
}

func GetDeviceAddress(request *request.AlexaRequest) (string, error) {
  accessDeviceAddressAPI(request, APIADDRESS)
  return "", nil
}

// DEVICE PARAMS (temp, zone, locale, etc)
func GetDeviceTimeZone(request *request.AlexaRequest) (string, error) {
  accessDeviceAPI(request, APIADDRESS)
  return "", nil
}

func GetDeviceDistanceUnit(request *request.AlexaRequest) (string, error) {
  accessDeviceAPI(request, APIADDRESS)
  return "", nil
}

func GetDeviceTemperatureUnit(request *request.AlexaRequest) (string, error) {
  accessDeviceAPI(request, APIADDRESS)
  return "", nil
}

// Still to implement:

// TO DO LISTS

// SHOPPING LISTS

