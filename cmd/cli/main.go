package main

import (
	"fmt"

	"github.com/AlexEkdahl/kango/config"
	"github.com/AlexEkdahl/kango/internal/cli"
	"github.com/AlexEkdahl/kango/internal/client"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	conf := config.New()
	c, err := client.New(*conf)
	if err != nil {
		fmt.Println("err", err)
	}
	opt := cli.Options{
		About:  cli.About{Name: "Kango", Version: "v0.0.1", ShortDescription: "Simple and efficient Kanban manager"},
		Client: c,
	}

	model := cli.New(opt)

	p := tea.NewProgram(model)
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
