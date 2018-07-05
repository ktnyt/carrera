package buffer

import (
	. "github.com/ktnyt/carrera"
	runewidth "github.com/mattn/go-runewidth"
)

type bufferPresenter struct {
	service BufferService
}

var tabwidth = DefaultConfig.TabWidth

func NewBufferPresenter(service BufferService) BufferPresenter {
	return bufferPresenter{service: service}
}

func (p bufferPresenter) Area(line, width, height int) [][]Cell {
	area := make([][]Cell, 0)
	pos := 0

	for line > 0 {
		if pos == p.service.Length() {
			return area
		}

		if p.service.Cell(pos).Ch == '\n' {
			line--
		}

		pos++
	}

	area = append(area, make([]Cell, 0))

	for pos < p.service.Length() && line < height {
		c := p.service.Cell(pos)

		if c.Ch == '\n' {
			area = append(area, make([]Cell, 0))
			line++
		} else {
			linewidth := 0

			if c.Ch == '\t' {
				linewidth += tabwidth
			} else {
				linewidth += runewidth.RuneWidth(c.Ch)
			}

			for _, c := range area[len(area)-1] {
				linewidth += runewidth.RuneWidth(c.Ch)
			}

			if linewidth >= width-1 {
				area[len(area)-1] = append(area[len(area)-1], Cell{
					Ch: '\\',
					Fg: Black,
					Bg: Default,
				})
				area = append(area, make([]Cell, 0))
				line++
			}

			if c.Ch == '\t' {
				for i := 0; i < tabwidth; i++ {
					area[len(area)-1] = append(area[len(area)-1], Cell{
						Ch: ' ',
						Fg: White,
						Bg: Default,
					})
				}
			} else {
				area[len(area)-1] = append(area[len(area)-1], c)
			}
		}

		pos++
	}

	for len(area) < height {
		area = append(area, []Cell{Cell{
			Ch: '~',
			Fg: Black,
			Bg: Default,
		}})
	}

	return area
}

func (p bufferPresenter) Insert(pos int, r rune) {
	p.service.Insert(pos, r)
}

func (p bufferPresenter) Delete(pos int) rune {
	return p.service.Delete(pos)
}
