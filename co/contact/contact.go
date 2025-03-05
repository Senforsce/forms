package contact

import (
	"githum/senforsce/forms/co/input"
	"githum/senforsce/forms/co/textarea"

	t1 "github.com/senforsce/tndr"
	"github.com/senforsce/tndr0cean/router"
)

type FormOptions struct {
	NameOptions    input.ComponentInputConfig
	EmailOptions   input.ComponentInputConfig
	MessageOptions textarea.ComponentTextAreaConfig
	Action         t1.SafeURL
	Method         string
}

func RenderForm(c *router.Context) {
	errs := map[string]string{}
	options := FormOptions{
		NameOptions: input.ComponentInputConfig{
			Name:        "username",
			Type:        "text",
			Placeholder: "enter your username",
			Id:          "ContactFormUsername",
		},
		EmailOptions: input.ComponentInputConfig{
			Name:        "email",
			Type:        "email",
			Placeholder: "enter your email",
			Id:          "ContactFormEmail",
		},

		MessageOptions: textarea.ComponentTextAreaConfig{
			Name:        "message",
			Placeholder: "enter your message",
			Id:          "ContactFormMessage",
		},
		Action: t1.SafeURL("/contact/process"),
		Method: "POST",
	}

	c.Render(Form(c, options, errs))
}
