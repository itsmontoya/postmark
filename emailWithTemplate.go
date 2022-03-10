package postmark

type EmailWithTemplate struct {
	Email

	TemplateID    int    `json:"TemplateId"`
	TemplateAlias string `json:"TemplateAlias"`

	TemplateModel map[string]interface{} `json:"TemplateModel"`

	InlineCSS bool `json:"InlineCss"`
}
