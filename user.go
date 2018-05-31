package publicbio

import (
	"github.com/microcosm-cc/bluemonday"
	"github.com/writeas/saturday"
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
	Links     []Link                   `json:"links"`
}

func (p *Profile) RenderedBio() template.HTML {
	return template.HTML(applyMarkdown(p.Bio.String))
}

func applyMarkdown(data string) string {
	mdExtensions := 0 |
		blackfriday.EXTENSION_TABLES |
		blackfriday.EXTENSION_FENCED_CODE |
		blackfriday.EXTENSION_AUTOLINK |
		blackfriday.EXTENSION_STRIKETHROUGH |
		blackfriday.EXTENSION_SPACE_HEADERS |
		blackfriday.EXTENSION_HEADER_IDS
	htmlFlags := 0 |
		blackfriday.HTML_USE_SMARTYPANTS |
		blackfriday.HTML_SMARTYPANTS_DASHES

	// Generate Markdown
	md := blackfriday.Markdown([]byte(data), blackfriday.HtmlRenderer(htmlFlags, "", ""), mdExtensions)
	// Strip out bad HTML
	policy := bluemonday.UGCPolicy()
	policy.AllowAttrs("target").OnElements("a")
	policy.AllowAttrs("style", "class", "id").Globally()
	return string(policy.SanitizeBytes(md))
}

type Link struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}
