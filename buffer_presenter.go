package carrera

import (
	"log"

	runewidth "github.com/mattn/go-runewidth"
)

type bufferPresenter struct {
	service BufferService
}

func NewBufferPresenter(service BufferService) BufferPresenter {
	return bufferPresenter{service: service}
}

func (p bufferPresenter) Area(line, width, height int) [][]rune {
	area := make([][]rune, 0)
	pos := 0

	for line > 0 {
		if pos == p.service.Length() {
			return area
		}

		r, err := p.service.Rune(pos)
		if err != nil {
			log.Fatal(err)
		}

		if r == '\n' {
			line--
		}

		pos++
	}

	area = append(area, make([]rune, 0))

	for pos < p.service.Length() && line < height {
		r, err := p.service.Rune(pos)
		if err != nil {
			log.Fatal(err)
		}

		if r == '\n' {
			area = append(area, make([]rune, 0))
			line++
		} else {
			linewidth := 0

			if r == '\t' {
				linewidth += tabwidth
			} else {
				linewidth += runewidth.RuneWidth(r)
			}

			for _, r := range area[len(area)-1] {
				linewidth += runewidth.RuneWidth(r)
			}

			if linewidth >= width-1 {
				area[len(area)-1] = append(area[len(area)-1], '\\')
				area = append(area, make([]rune, 0))
				line++
			}

			if r == '\t' {
				for i := 0; i < tabwidth; i++ {
					area[len(area)-1] = append(area[len(area)-1], ' ')
				}
			} else {
				area[len(area)-1] = append(area[len(area)-1], r)
			}
		}

		pos++
	}

	for len(area) < height {
		area = append(area, []rune{'~'})
	}

	return area
}

func (p bufferPresenter) Insert(pos int, r rune) {
	if err := p.service.Insert(pos, r); err != nil {
		log.Fatal(err)
	}
}

func (p bufferPresenter) Delete(pos int) rune {
	r, err := p.service.Delete(pos)
	if err != nil {
		log.Fatal(err)
	}

	return r
}
