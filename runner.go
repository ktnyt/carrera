package carrera

import (
	termbox "github.com/nsf/termbox-go"
)

func Run() error {
	termbox.Init()

	var err error
	run := true

	for run {
		switch e := termbox.PollEvent(); e.Type {
		case termbox.EventKey:
			switch e.Key {
			case termbox.KeyEsc:
				run = false
			}
		case termbox.EventError:
			run = true
		}
	}

	termbox.Close()
	return err
}
