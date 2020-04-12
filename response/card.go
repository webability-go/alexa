package response

const (
	PERMISSION_FULLNAME                = "alexa::profile:name:read"
	PERMISSION_FIRSTNAME               = "alexa::profile:given_name:read"
	PERMISSION_EMAIL                   = "alexa::profile:email:read"
	PERMISSION_MOBILE                  = "alexa::profile:mobile_number:read"
	PERMISSION_COUNTRY_AND_POSTAL_CODE = "read::alexa:device:all:address:country_and_postal_code"
	PERMISSION_ADDRESS                 = "read::alexa:device:all:address"
)

type CardBuilder struct {
	Type        string
	Title       string
	Message     string
	Image       string
	Icon        string
	Permissions []string
}

// main builder
func NewCardBuilder(title string, message string, image string, icon string) *CardBuilder {
	return &CardBuilder{
		Type:    "Simple",
		Title:   title,
		Message: message,
		Image:   image,
		Icon:    icon,
	}
}

// Simple card: no image
func NewSimpleCardBuilder(title string, message string, image string, icon string) *CardBuilder {
	return &CardBuilder{
		Type:    "Simple",
		Title:   title,
		Message: message,
		Image:   image,
		Icon:    icon,
	}
}

// Standard card: with image
func NewStandardCardBuilder(title string, message string, image string, icon string) *CardBuilder {
	return &CardBuilder{
		Type:    "Standard",
		Title:   title,
		Message: message,
		Image:   image,
		Icon:    icon,
	}
}

// Link account: login with oauth or basic
func NewLinkAccountCardBuilder(title string, message string, image string, icon string) *CardBuilder {
	return &CardBuilder{
		Type:    "LinkAccount",
		Title:   title,
		Message: message,
		Image:   image,
		Icon:    icon,
	}
}

// Ask for permission card
func NewPermissionCardBuilder(permissions []string) *CardBuilder {
	return &CardBuilder{
		Type:        "AskForPermissionsConsent",
		Permissions: permissions,
	}
}

// Build the code for the specified card
func (builder *CardBuilder) Build() *Card {

	var card *Card

	switch builder.Type {
	case "Simple":
		card = &Card{
			Type:    "Simple",
			Title:   builder.Title,
			Content: builder.Message,
		}
	case "Standard":
		card = &Card{
			Type:    "Simple",
			Title:   builder.Title,
			Content: builder.Message,
			Image: &CardImage{
				SmallImageURL: builder.Icon,
				LargeImageURL: builder.Image,
			},
		}
	case "LinkAccount":
		card = &Card{
			Type: "LinkAccount",
		}
	case "AskForPermissionsConsent":
		card = &Card{
			Type:        "AskForPermissionsConsent",
			Permissions: builder.Permissions,
		}
	}

	return card
}
