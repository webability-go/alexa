package response

/* datatype es:
   text
   break
   audio

   text contains text to say

   data contains the type of data to apply

   Rules:
   normal text               datatype=text, text=[say it]
   whispered effect:         datatype=text, text=[say it], tag=effect, data=whispered
   emphasis effect:          datatype=text, text=[say it], tag=emphasis, data=strong/moderate/reduced
   lang effect:              datatype=text, text=[say it], tag=emphasis, data=strong/moderate/reduced
   audio:                    datatype=audio, data=SRC
   break:                    datatype=break, data="Ns" or "none".to."x-strong" to apply to time or strength
*/

type SSML struct {
	datatype string // text, break, audio
	text     string

	effect      string
	emphasis    string
	lang        string
	paragraph   bool
	sentence    bool
	sayas       string
	sayasformat string
	voice       string
}

type SSMLBuilder struct {
	SSML []SSML
}

func NewSSMLBuilder() *SSMLBuilder {
	return &SSMLBuilder{}
}

// text is the text to say
func (builder *SSMLBuilder) Raw(text string) {
	builder.SSML = append(builder.SSML, SSML{datatype: "ssml", text: text})
}

func (builder *SSMLBuilder) Say(text string) {
	builder.SSML = append(builder.SSML, SSML{datatype: "text", text: text})
}

func (builder *SSMLBuilder) AddEffect(data string) {
	builder.SSML[len(builder.SSML)-1].effect = data
}

func (builder *SSMLBuilder) AddEmphasis(data string) {
	builder.SSML[len(builder.SSML)-1].emphasis = data
}

func (builder *SSMLBuilder) AddLang(data string) {
	builder.SSML[len(builder.SSML)-1].lang = data
}

func (builder *SSMLBuilder) SetParagraph() {
	builder.SSML[len(builder.SSML)-1].paragraph = true
}

func (builder *SSMLBuilder) SetSentence() {
	builder.SSML[len(builder.SSML)-1].sentence = true
}

func (builder *SSMLBuilder) AddSayAs(data string, format string) {
	builder.SSML[len(builder.SSML)-1].sayas = data
	builder.SSML[len(builder.SSML)-1].sayasformat = format
}

func (builder *SSMLBuilder) AddVoice(data string) {
	builder.SSML[len(builder.SSML)-1].voice = data
}

// text is the break time
func (builder *SSMLBuilder) Break(text string) {
	builder.SSML = append(builder.SSML, SSML{datatype: "break", text: text})
}

// text is the audio source
func (builder *SSMLBuilder) Audio(text string) {
	builder.SSML = append(builder.SSML, SSML{datatype: "audio", text: text})
}

func (builder *SSMLBuilder) Build() string {
	var response string
	for _, ssml := range builder.SSML {
		switch ssml.datatype {
		case "ssml":
			response += ssml.text
		case "text":
			if ssml.voice != "" {
				response += "<voice name=\"" + ssml.voice + "\">"
			}
			if ssml.effect != "" {
				response += "<amazon:effect name=\"" + ssml.effect + "\">"
			}

			response += ssml.text + " "

			if ssml.effect != "" {
				response += "</amazon:effect>"
			}
			if ssml.voice != "" {
				response += "</voice>"
			}
		case "break":
			response += "<break time='" + ssml.text + "ms'/> "
		case "audio":
			response += "<audio src='" + ssml.text + "'/> "
		}
	}
	if response[0:5] != "<spea" {
		return "<speak>" + response + "</speak>"
	}
	return response
}
