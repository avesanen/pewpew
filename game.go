package main

import (
	"encoding/json"
	"log"
	"time"
)

type collision struct {
	Collider string  `json:"collider"`
	Target   string  `json:"target"`
	Position *Vector `json:"position"`
}

type game struct {
	LastUpdate time.Time
	GameArea   [2]float64   `json:"gameArea"`
	Tiles      []*tile      `json:"tiles"`
	Players    []*player    `json:"players"`
	Bullets    []*bullet    `json:"bullets"`
	Collisions []*collision `json:"collisions"`
}

func (g *game) getState() []byte {
	b, _ := json.Marshal(g)
	return b
}

func (g *game) newPlayer(position *Vector, name string) *player {
	p := &player{}
	p.Id = Uuid()
	p.Radius = 1
	p.Position = position
	p.Velocity = &Vector{0, 0}
	p.Name = name
	g.Players = append(g.Players, p)
	log.Println("newplayer", p.Position, p.Velocity)
	return p
}

func (g *game) newBullet(position *Vector, velocity *Vector, firedBy string) *bullet {
	b := &bullet{}
	b.Id = Uuid()
	b.Radius = 0.1
	b.Position = position
	b.Velocity = velocity
	b.FiredBy = firedBy
	g.Bullets = append(g.Bullets, b)
	return b
}

func (g *game) newTile(position *Vector, width, height float64) *tile {
	t := &tile{}
	t.Id = Uuid()
	t.Width = width
	t.Height = height
	t.Position = position
	t.Velocity = &Vector{0, 0}
	g.Tiles = append(g.Tiles, t)
	return t
}

func (g *game) rmPlayer(p *player) {
	for i, k := range g.Players {
		if k == p {
			g.Players = append(g.Players[:i], g.Players[i+1:]...)
		}
	}
}

func (g *game) rmBullet(b *bullet) {
	for i, k := range g.Bullets {
		if k == b {
			g.Bullets = append(g.Bullets[:i], g.Bullets[i+1:]...)
		}
	}
}

func (g *game) rmTile(t *tile) {
	for i, k := range g.Tiles {
		if k == t {
			g.Tiles = append(g.Tiles[:i], g.Tiles[i+1:]...)
		}
	}
}

func (g *game) update() {
	g.Collisions = nil
	dt := time.Since(g.LastUpdate)
	for _, b := range g.Bullets {
		b.update(g, dt)
		// Remove all bullets that are dead.
		var newBullets []*bullet
		for _, k := range g.Bullets {
			if !k.Dead {
				newBullets = append(newBullets, k)
			}
		}
		g.Bullets = newBullets
	}
	for _, p := range g.Players {
		p.update(g, dt)
	}
	g.LastUpdate = time.Now()
}
