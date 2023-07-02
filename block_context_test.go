package slack

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewContextBlock(t *testing.T) {

	locationPinImage := NewImageBlockElement("https://api.slack.com/img/blocks/bkb_template_images/tripAgentLocationMarker.png", "Location Pin Icon")
	textExample := NewTextBlockObject("plain_text", "Location: Central Business District", true, false)

	elements := []MixedElement{locationPinImage, textExample}

	contextBlock := NewContextBlock("test", elements...)
	assert.Equal(t, string(contextBlock.Type), "context")
	assert.Equal(t, contextBlock.BlockID, "test")
	assert.Equal(t, len(contextBlock.ContextElements.Elements), 2)

}

func TestContextElements_MarshalJSON(t *testing.T) {
	elements := ContextElements{
		Elements: []MixedElement{
			ImageBlockElement{
				Type:     "image",
				ImageURL: "https://example.com/1.jpg",
				AltText:  "image",
			},
			TextBlockObject{
				Type: "mrkdwn",
				Text: "sample",
			},
		},
	}

	assert := require.New(t)
	marshaled, err := json.Marshal(elements)
	assert.NoError(err)

	const want = `[{"type":"image","image_url":"https://example.com/1.jpg","alt_text":"image"},{"type":"mrkdwn","text":"sample"}]`
	assert.JSONEq(want, string(marshaled))
}
