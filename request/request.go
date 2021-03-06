package request

import (
//  "fmt"
  "encoding/json"
  "reflect"
)

type AlexaRequest struct {
  Version               string                   `json:"version"`
  Session               Session                  `json:"session"`
  Context               Context                  `json:"context"`
  Request               Request                  `json:"request"`
}

type Session struct {
  New                   bool                     `json:"new"`
  SessionID             string                   `json:"sessionId"`
  Attributes            interface{}              `json:"attributes,omitempty"`
  Application struct {
    ApplicationID       string                   `json:"applicationId,omitempty"`
  }                                              `json:"application"`
  User struct {
    UserID              string                   `json:"userId"`
    AccessToken         string                   `json:"accessToken,omitempty"`
    Permissions struct {
      ConsentToken      string                   `json:"consentToken,omitempty"`
    }                                            `json:"permissions,omitempty"`
  }                                              `json:"user"`
}

type Context struct {
  System struct {
    Application struct {
      ApplicationID     string                   `json:"applicationId,omitempty"`
    }                                            `json:"application,omitempty"`
    User struct {
      UserID            string                   `json:"userId,omitempty"`
      AccessToken       string                   `json:"accessToken,omitempty"`
      Permissions struct {
        ConsentToken    string                   `json:"consentToken,omitempty"`
      }                                          `json:"permissions,omitempty"`
    }                                            `json:"user,omitempty"`
    APIAccessToken      string                   `json:"apiAccessToken"`
    APIEndPoint         string                   `json:"apiEndpoint"`
    Device struct {
      DeviceID          string                   `json:"deviceId,omitempty"`
      SupportedInterfaces struct {
        Display *struct {
          TemplateVersion string                 `json:"templateVersion,omitempty"`
          MarkupVersion string                   `json:"markupVersion,omitempty"`
        }                                        `json:"Display,omitempty"`
        AudioPlayer *struct {
        }                                        `json:"AudioPlayer,omitempty"`
        VideoApp *struct {
        }                                        `json:"VideoApp,omitempty"`
        AlexaPresentationAPL *struct {
        }                                        `json:"Alexa.Presentation.APL,omitempty"`
      }                                          `json:"supportedInterfaces,omitempty"`
    }                                            `json:"device,omitempty"`
  }                                              `json:"System,omitempty"`
  AudioPlayer struct {
    PlayerActivity      string                   `json:"playerActivity,omitempty"`
    Token               string                   `json:"token,omitempty"`
    OffsetInMilliseconds int                     `json:"offsetInMilliseconds,omitempty"`
  } `json:"AudioPlayer,omitempty"`
  
  /* Check documentation for those objects */
  Viewport struct {
    Experiences []struct {
      ArcMinuteWidth    int                      `json:"arcMinuteWidth"`
      ArcMinuteHeight   int                      `json:"arcMinuteHeight"`
      CanRotate         bool                     `json:"canRotate"`
      CanResize         bool                     `json:"canResize"`
    }                                            `json:"experiences"`
    Shape               string                   `json:"shape"`
    PixelWidth          int                      `json:"pixelWidth"`
    PixelHeight         int                      `json:"pixelHeight"`
    DPI                 int                      `json:"dpi"`
    CurrentPixelWidth   int                      `json:"currentPixelWidth"`
    CurrentPixelHeight  int                      `json:"currentPixelHeight"`
    Touch               []string                 `json:"touch"`
  }                                              `json:"Viewport,omitempty"`
  Display struct {
    Token string `json:"token,omitempty"`
  } `json:"Display,omitempty"`
}

/* Partially implemented for now */
type Request struct {
  Type                  string                   `json:"type"`
  RequestID             string                   `json:"requestId"`
  Timestamp             string                   `json:"timestamp"`
  Locale                string                   `json:"locale"`
  // Intent request
  Intent                Intent                   `json:"intent,omitempty"`
  DialogState           string                   `json:"dialogState,omitempty"`

  // Ended session request
  Reason                string                   `json:"reason,omitempty"`
  Error                 map[string]string        `json:"error,omitempty"`
  
  // Audio player request and others
  Token                 string                   `json:"token,omitempty"`
  OffsetInMilliseconds  int                      `json:"offsetInMilliseconds,omitempty"`
  
  // implements other needed type of requests: game request, gadget request, playback request...
}

type Intent struct {
  Name                  string                   `json:"name"`
  ConfirmationStatus    string                   `json:"confirmationStatus"`
  Slots                 map[string]Slot          `json:"slots"`
}

type Slot struct {
  Name                  string                   `json:"name"`
  Value                 string                   `json:"value"`
  Resolutions           Resolutions              `json:"resolutions"`
  ConfirmationStatus    string                   `json:"confirmationStatus"`
  Source                string                   `json:"source"`
}

type Resolutions struct {
  ResolutionsPerAuthority []struct {
    Authority           string                   `json:"authority"`
    Status struct {
      Code              string                   `json:"code"`
    }                                            `json:"status"`
    Values []struct {
      Value struct {
        Name            string                   `json:"name"`
        Id              string                   `json:"id"`
      }                                          `json:"value"`
    }                                            `json:"values"`
  }                                              `json:"resolutionsPerAuthority"`
}

func (request AlexaRequest)GetRequestType() string {
  return request.Request.Type
}

func (request AlexaRequest)GetRequestIntentName() string {
  return request.Request.Intent.Name
}

func (request AlexaRequest)GetSessionId() string {
  return request.Session.SessionID
}

func (request AlexaRequest)GetNewSession() bool {
  return request.Session.New
}

func (request AlexaRequest)GetUserId() string {
  return request.Session.User.UserID
}

func (request AlexaRequest)GetLocale() string {
  return request.Request.Locale
}

func (request AlexaRequest)GetAttributes() interface{} {
  return request.Session.Attributes
}

func (request AlexaRequest)GetSlots() *map[string]Slot {
  return &(request.Request.Intent.Slots)
}

func (request AlexaRequest)GetDisplay() interface{} {
  return request.Context.System.Device.SupportedInterfaces.Display
}

func (request AlexaRequest)HasDisplay() bool {
  display := request.GetDisplay()
  return !(display == nil || reflect.ValueOf(display).IsNil())
}

func (request AlexaRequest)GetVideo() interface{} {
  return request.Context.System.Device.SupportedInterfaces.VideoApp
}

func (request AlexaRequest)HasVideo() bool {
  display := request.GetVideo()
  return !(display == nil || reflect.ValueOf(display).IsNil())
}

func (request AlexaRequest)GetAPL() interface{} {
  return request.Context.System.Device.SupportedInterfaces.AlexaPresentationAPL
}

func (request AlexaRequest)HasAPL() bool {
  display := request.GetAPL()
  return !(display == nil || reflect.ValueOf(display).IsNil())
}

/* =============================================================
   HIJACK THE ATTRIBUTES UNMARSHAL ON THE REQUEST
   =============================================================*/
var SessionUnmarshalerHandler func(data []byte, s *Session) error = nil

func (s *Session)UnmarshalJSON(data []byte) error {
  if SessionUnmarshalerHandler == nil {
    type Alias Session
    aux := &struct {
      Attributes        map[string]interface{}   `json:"attributes,omitempty"`
      *Alias
    }{
      Alias: (*Alias)(s),
    }
    if err := json.Unmarshal(data, &aux); err != nil {
      return err
    }
    s.Attributes = aux.Attributes
    return nil
  }
  return SessionUnmarshalerHandler(data, s)
}

