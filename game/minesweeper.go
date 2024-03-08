package game

import (
	"fmt"
	"math/rand"
)

type Minesweeper struct {
	Gamemap []uint8
	Mask    []bool
	Width   int
	height  int
	mines   int
}

func NewSweeper(Width int, height int, mines int) *Minesweeper {
	gmap := make([]uint8, Width*height)
	gmask := make([]bool, Width*height)
	sweeper := Minesweeper{
		Gamemap: gmap,
		Mask:    gmask,
		Width:   Width,
		height:  height,
		mines:   mines,
	}

	sweeper.set_mines()
	sweeper.set_numbers()
	sweeper.draw()
	return &sweeper
}

func (s *Minesweeper) CheckWin() bool {
	count := 0
	for _, v := range s.Mask {
		if v {
			count++
		}
	}
	return s.Width*s.height == (count + s.mines)
}

func (s *Minesweeper) MapClient() []byte {
	newmap := make([]byte, s.Width*s.height)
	for i := 0; i < len(newmap); i++ {
		if !s.Mask[i] {
			newmap[i] = 10
		} else {
			newmap[i] = s.Gamemap[i]
		}
	}
	return newmap
}

func (s *Minesweeper) Reset(Width int, height int, mines int) {
	if Width != s.Width || height != s.height {
		s.Gamemap = make([]uint8, Width*height)
		s.height = height
		s.Width = Width
	} else {
		clear(s.Gamemap)
	}
	s.set_mines()
	s.set_numbers()
}

func (s *Minesweeper) set_mines() {
	tilec := s.Width * s.height
	for i := 0; i < s.mines; i++ {
		mine_index := rand.Intn(tilec)
		for s.Gamemap[mine_index] != 0 {
			mine_index++
			if mine_index >= tilec {
				mine_index = 0
			}
		}
		s.Gamemap[mine_index] = 9
	}
}

func (s *Minesweeper) set_numbers() {
	for i := range s.Gamemap {
		x := i % s.height
		y := i / s.Width

		if s.Gamemap[i] == 9 {
			continue
		}

		x1 := x - 1
		x2 := x + 1
		y1 := y - 1
		y2 := y + 1

		if x1 < 0 {
			x1 = 0
		}
		if x2 >= s.Width-1 {
			x2 = s.Width - 1
		}
		if y1 < 0 {
			y1 = 0
		}
		if y2 >= s.height-1 {
			y2 = s.height - 1
		}
		area := 0
		for j := x1; j < x2+1; j++ {
			for k := y1; k < y2+1; k++ {
				if s.Gamemap[k*s.Width+j] == 9 {
					area++
				}
			}
		}
		s.Gamemap[i] = uint8(area)
	}
}

func (s *Minesweeper) Reveal(pos int) bool {
	fmt.Println("Reveal, ", pos)
	fmt.Println(s.Gamemap)
	tile := s.Gamemap[pos]

	// mine
	if tile == 9 {
		s.reveal(pos)
		return true
	} else if tile == 0 {
		s.multi_reveal(pos)
	} else {
		s.reveal(pos)
	}
	return false
}

func (s *Minesweeper) reveal(pos int) {
	s.Mask[pos] = true
}

func (s *Minesweeper) multi_reveal(pos int) {
	x := pos % s.height
	y := pos / s.Width
	x1 := x - 1
	x2 := x + 1
	y1 := y - 1
	y2 := y + 1

	if x1 < 0 {
		x1 = 0
	}
	if x2 >= s.Width-1 {
		x2 = s.Width - 1
	}
	if y1 < 0 {
		y1 = 0
	}
	if y2 >= s.height-1 {
		y2 = s.height - 1
	}
	for j := x1; j < x2+1; j++ {
		for k := y1; k < y2+1; k++ {
			pos := k*s.Width + j
			if !s.Mask[pos] {
				s.reveal(pos)
				if s.Gamemap[pos] == 0 {
					s.multi_reveal(pos)
				}
			}
		}
	}
}

// this doesnt work rn
func (s *Minesweeper) draw() {
	for i := 0; i < s.height; i++ {
		fmt.Println(s.Gamemap[i*s.height : i*s.height+s.Width])
	}
}
