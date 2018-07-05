package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/alecthomas/chroma/lexers"
	"github.com/ktnyt/carrera"
	carreraBuffer "github.com/ktnyt/carrera/buffer"
	termbox "github.com/ktnyt/termbox-go"
)

func openFile(filename string) carrera.BufferService {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	p, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	lexer := lexers.Match(filename)
	syntax := carrera.NewSyntax(lexer)
	buffer := carrera.Buffer([]rune(string(p)))
	service := carreraBuffer.NewBufferService(buffer, syntax.Highlight)

	return service
}

func run(filename string) error {
	var err error
	run := true

	service := openFile(filename)

	termbox.Init()
	termbox.SetOutputMode(termbox.OutputXterm)

	presenter := carreraBuffer.NewBufferPresenter(service)
	width, height := termbox.Size()
	view := carreraBuffer.NewBufferView(presenter, width, height)

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
