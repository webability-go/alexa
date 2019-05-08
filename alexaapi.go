package alexa

import (
  "fmt"
  "strings"
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
  
  APICOUNTRY =          "/v1/devices/{deviceId}/settings/address/countryAndPostalCode"
  APIADDRESS =          "/v1/devices/{deviceId}/settings/address"
)

func getAPIAccessToken(request *request.AlexaRequest) string {
  return request.Context.System.APIAccessToken
}

func getAPIEndPoint(request *request.AlexaRequest) string {
  return request.Context.System.APIEndPoint
}

func getAPIDevideId(request *request.AlexaRequest) string {
  return request.Context.System.Device.DeviceID
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
  
  if DEVEL {
    fmt.Println("RAW DATA", string(data))
  }
  return data, nil
}

func accessAccountAPI(request *request.AlexaRequest, service string) (interface{}, error) {
  accessToken := getAPIAccessToken(request)
  endPoint := getAPIEndPoint(request)
  if accessToken == "" || endPoint == "" {
    return nil, errors.New("Error: There is no api endPoint or accessToken in the request.")
  }
  jsonstring, err := accessAPI(endPoint, service, accessToken)
  var data interface{}
  err = json.Unmarshal(jsonstring, &data)
  if err != nil {
    return nil, err
  }
  return data, nil
}

func accessDeviceAddressAPI(request *request.AlexaRequest, service string) (interface{}, error) {
  accessToken := getAPIAccessToken(request)
  endPoint := getAPIEndPoint(request)
  if accessToken == "" || endPoint == "" {
    return nil, errors.New("Error: There is no api endPoint or accessToken in the request.")
  }
  id := getAPIDevideId(request)
  service = strings.Replace(service, "{deviceId}", id, -1)
  jsonstring, err := accessAPI(endPoint, service, accessToken)
  var data interface{}
  err = json.Unmarshal(jsonstring, &data)
  if err != nil {
    return nil, err
  }
  
  if DEVEL {
    fmt.Println("ARRAY DATA", data)
  }
  return data, nil
}

func accessDeviceAPI(request *request.AlexaRequest, service string) (interface{}, error) {
  accessToken := getAPIAccessToken(request);
  endPoint := getAPIEndPoint(request);
  if accessToken == "" || endPoint == "" {
    return nil, nil
  }
  id := getAPIDevideId(request)
  service = strings.Replace(service, "{deviceId}", id, -1)
  jsonstring, err := accessAPI(endPoint, service, accessToken)
  if err != nil {
    return nil, err
  }
  var data interface{}
  err = json.Unmarshal(jsonstring, &data)
  if err != nil {
    return nil, err
  }
  
  if DEVEL {
    fmt.Println("ARRAY DATA", data)
  }
  return data, nil
}

/* ACCOUNT API */
func GetAccountFullName(request *request.AlexaRequest) (string, error) {
  data, err := accessAccountAPI(request, APINAME)
  if err != nil {
    return "", err
  }
  switch data.(type) {
    case string:
      return data.(string), nil
    case map[string]interface{}:
      return "", errors.New(fmt.Sprint(data))
  }
  return "", errors.New("Data type not known: " + fmt.Sprint(data))
}

func GetAccountGivenName(request *request.AlexaRequest) (string, error) {
  data, err := accessAccountAPI(request, APIGIVENNAME)
  if err != nil {
    return "", err
  }
  switch data.(type) {
    case string:
      return data.(string), nil
    case map[string]interface{}:
      return "", errors.New(fmt.Sprint(data))
  }
  return "", errors.New("Data type not known: " + fmt.Sprint(data))
}

func GetAccountEmail(request *request.AlexaRequest) (string, error) {
  data, err := accessAccountAPI(request, APIEMAIL)
  if err != nil {
    return "", err
  }
  switch data.(type) {
    case string:
      return data.(string), nil
    case map[string]interface{}:
      return "", errors.New(fmt.Sprint(data))
  }
  return "", errors.New("Data type not known: " + fmt.Sprint(data))
}

func GetAccountMobileNumber(request *request.AlexaRequest) (map[string]interface{}, error) {
  data, err := accessAccountAPI(request, APIMOBILENUMBER)
  if err != nil {
    return nil, err
  }
  switch data.(type) {
    case map[string]interface{}:
      ndata := data.(map[string]interface{})
      _, ok := ndata["countryCode"]
      if (!ok) {
        return nil, errors.New(fmt.Sprint(data))
      }
      return ndata, nil
  }
  return nil, errors.New("Data type not known: " + fmt.Sprint(data))
}

/* DEVICE ADDRESS AND COUNTRY */
func GetDeviceCountry(request *request.AlexaRequest) (string, error) {
  data, err := accessDeviceAddressAPI(request, APICOUNTRY)
  if err != nil {
    return "", err
  }
  switch data.(type) {
    case string:
      return data.(string), nil
    case map[string]interface{}:
      return "", errors.New(fmt.Sprint(data))
  }
  return "", errors.New("Data type not known: " + fmt.Sprint(data))
}

func GetDeviceAddress(request *request.AlexaRequest) (string, error) {
  data, err := accessDeviceAddressAPI(request, APIADDRESS)
  if err != nil {
    return "", err
  }
  switch data.(type) {
    case string:
      return data.(string), nil
    case map[string]interface{}:
      return "", errors.New(fmt.Sprint(data))
  }
  return "", errors.New("Data type not known: " + fmt.Sprint(data))
}

// DEVICE PARAMS (temp, zone, locale, etc)
func GetDeviceTimeZone(request *request.AlexaRequest) (string, error) {
  data, err := accessDeviceAPI(request, APITIMEZONE)
  if err != nil {
    return "", err
  }
  switch data.(type) {
    case string:
      return data.(string), nil
    case map[string]interface{}:
      return "", errors.New(fmt.Sprint(data))
  }
  return "", errors.New("Data type not known: " + fmt.Sprint(data))
}

func GetDeviceDistanceUnit(request *request.AlexaRequest) (string, error) {
  data, err := accessDeviceAPI(request, APIDISTANCEUNIT)
  if err != nil {
    return "", err
  }
  switch data.(type) {
    case string:
      return data.(string), nil
    case map[string]interface{}:
      return "", errors.New(fmt.Sprint(data))
  }
  return "", errors.New("Data type not known: " + fmt.Sprint(data))
}

func GetDeviceTemperatureUnit(request *request.AlexaRequest) (string, error) {
  data, err := accessDeviceAPI(request, APITEMPERATUREUNIT)
  if err != nil {
    return "", err
  }
  switch data.(type) {
    case string:
      return data.(string), nil
    case map[string]interface{}:
      return "", errors.New(fmt.Sprint(data))
  }
  return "", errors.New("Data type not known: " + fmt.Sprint(data))
}

// Still to implement:

// TO DO LISTS

// SHOPPING LISTS

