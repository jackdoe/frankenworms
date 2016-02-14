package main

import (
	tl "github.com/JoelOtter/termloop"
	"time"
)

func world() {
	game := tl.NewGame()
	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorWhite,
		Fg: tl.ColorWhite,
		Ch: '_',
	})
	for i := -1000; i < 1000; i = i + 40 {
		if i == 0 {
			continue
		}
		for j := -1000; j < 1000; j = j + 40 {
			level.AddEntity(tl.NewRectangle(i, j, 20, 10, tl.ColorBlue))
		}
	}
	player := Player{
		entity: tl.NewEntity(1, 1, 1, 1),
		level:  level,
	}

	player.entity.SetCell(0, 0, &tl.Cell{Fg: tl.ColorBlack, Ch: '옷'})
	level.AddEntity(&player)
	game.Screen().SetLevel(level)
	go func() {
		for {
			player.Tick(tl.Event{})
			time.Sleep(200 * time.Millisecond)
		}
	}()
	game.Start()
}

type Player struct {
	entity *tl.Entity
	level  *tl.BaseLevel
}

func (player *Player) Draw(screen *tl.Screen) {
	screenWidth, screenHeight := screen.Size()
	x, y := player.entity.Position()
	player.level.SetOffset(screenWidth/2-x, screenHeight/2-y)
	player.entity.Draw(screen)
}

func (player *Player) Tick(event tl.Event) {

	x, y := player.entity.Position()

	delta := BODY.left.delta()
	if delta > 0 {
		x--
		y++
	} else if delta < 0 {
		y--
		x--
	}
	delta = BODY.right.delta()
	if delta > 0 {
		x++
		y++
	} else if delta < 0 {
		x++
		y--
	}
	player.entity.SetPosition(x, y)
}

func (player *Player) Size() (int, int)              { return player.entity.Size() }
func (player *Player) Position() (int, int)          { return player.entity.Position() }
func (player *Player) Collide(collision tl.Physical) {}
