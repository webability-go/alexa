package response

type CardBuilder struct {
  Title string
  Message string
  Image string
  Icon string
}

func NewCardBuilder(title string, message string, image string, icon string) *CardBuilder {
  return &CardBuilder{
    Title: title,
    Message: message,
    Image: image,
    Icon: icon,
  }
}

// Make 2 types of cards: simple and rich

func (builder *CardBuilder) Build() *Card {
  card := &Card{
        Type: "Simple",
        Title: builder.Title,
        Content: builder.Message,
        Image: &CardImage{
          SmallImageURL: builder.Icon,
          LargeImageURL: builder.Image,
        },
      }
  return card
}

