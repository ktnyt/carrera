package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"unicode/utf8"

	"github.com/ktnyt/carrera"
	termbox "github.com/nsf/termbox-go"
)

func newBufferServiceFromFile(file io.Reader) carrera.BufferService {
	p, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	buffer := carrera.NewBufferWithCapacity(len(p))
	service := carrera.NewBufferService(buffer)

	i := 0
	pos := 0
	for i < len(p) {
		r, size := utf8.DecodeRune(p[i:])
		service.Insert(pos, r)
		i += size
		pos++
	}

	return service
}

func run(filename string) error {
	var err error
	run := true

	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	service := newBufferServiceFromFile(file)
	file.Close()

	termbox.Init()

	presenter := carrera.NewBufferPresenter(service)
	width, height := termbox.Size()
	view := carrera.NewBufferView(presenter, width, height)

	for run {
		view.Draw()
		termbox.Flush()
		switch e := termbox.PollEvent(); e.Type {
		case termbox.EventKey:
			switch e.Key {
			case termbox.KeyEsc:
				run = false
			default:
				switch e.Ch {
				case 'q':
					run = false
				}
			}
		case termbox.EventError:
			run = true
		}
	}

	termbox.Close()
	return err
}

func main() {
	if err := run(os.Args[1]); err != nil {
		log.Fatal(err)
	}
}
