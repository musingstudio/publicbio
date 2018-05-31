package publicbio

import (
	"github.com/writeas/web-core/converter"
	"html/template"
)

// Profile is a publicly-viewable user, containing only the data necessary
// to display a profile.
type Profile struct {
	AvatarURL string                   `json:"avatar_url"`
	Username  string                   `json:"username"`
	Name      converter.NullJSONString `json:"name"`
	Header    converter.NullJSONString `json:"header"`
	Bio       converter.NullJSONString `json:"bio"`
}

func (p *Profile) RenderedBio() template.HTML {
	return template.HTML(p.Bio.String)
}
