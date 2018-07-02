package carrera

import (
	runewidth "github.com/mattn/go-runewidth"
)

type bufferPresenter struct {
	service BufferService
}

func NewBufferPresenter(service BufferService) BufferPresenter {
	return bufferPresenter{service: service}
}

func (p bufferPresenter) Area(line, width, height int) [][]Rune {
	area := make([][]Rune, 0)
	pos := 0

	for line > 0 {
		if pos == p.service.Length() {
			return area
		}

		if p.service.Rune(pos).Value == '\n' {
			line--
		}

		pos++
	}

	area = append(area, make([]Rune, 0))

	for pos < p.service.Length() && line < height {
		r := p.service.Rune(pos)

		if r.Value == '\n' {
			area = append(area, make([]Rune, 0))
			line++
		} else {
			linewidth := 0

			if r.Value == '\t' {
				linewidth += tabwidth
			} else {
				linewidth += runewidth.RuneWidth(r.Value)
			}

			for _, r := range area[len(area)-1] {
				linewidth += runewidth.RuneWidth(r.Value)
			}

			if linewidth >= width-1 {
				area[len(area)-1] = append(area[len(area)-1], NewRune('\\', Black))
				area = append(area, make([]Rune, 0))
				line++
			}

			if r.Value == '\t' {
				for i := 0; i < tabwidth; i++ {
					area[len(area)-1] = append(area[len(area)-1], NewRune(' ', White))
				}
			} else {
				area[len(area)-1] = append(area[len(area)-1], r)
			}
		}

		pos++
	}

	for len(area) < height {
		area = append(area, []Rune{NewRune('~', Black)})
	}

	return area
}

func (p bufferPresenter) Insert(pos int, r rune) {
	p.service.Insert(pos, r)
}

func (p bufferPresenter) Delete(pos int) rune {
	return p.service.Delete(pos)
}
