package cli

import (
	"github.com/AlexEkdahl/kango/internal/client"
)

type Options struct {
	About  About
	Client client.Client
}

type About struct {
	Name             string
	Version          string
	ShortDescription string
}
