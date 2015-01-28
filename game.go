package main

import (
	"encoding/json"
	"time"
)

type game struct {
	Players    []*player  `json:"players"`
	Bullets    []*bullet  `json:"bullets"`
	GameArea   [2]float64 `json:"gameArea"`
	LastUpdate time.Time
}

func (g *game) getState() []byte {
	b, err := json.Marshal(g)
	if err != nil {
		return nil
	}
	return b
}

func (g *game) deleteBullet(b *bullet) {
	for i, k := range g.Bullets {
		if k == b {
			g.Bullets = append(g.Bullets[:i], g.Bullets[i+1:]...)
		}
	}
}

func (g *game) update() {
	for _, p := range g.Players {
		p.Update(time.Since(g.LastUpdate))
		if p.Location[0] < 0.0 {
			p.Location[0] = 0.0
		}
		if p.Location[1] < 0.0 {
			p.Location[1] = 0.0
		}
		if p.Location[0] > g.GameArea[0] {
			p.Location[0] = g.GameArea[0]
		}
		if p.Location[1] > g.GameArea[1] {
			p.Location[1] = g.GameArea[1]
		}
	}

	var newbullets []*bullet
	for _, b := range g.Bullets {
		b.Update(time.Since(g.LastUpdate))
		if b.Location[0] > 0.0 && b.Location[0] < g.GameArea[0] &&
			b.Location[1] > 0.0 && b.Location[1] < g.GameArea[1] {
			newbullets = append(newbullets, b)
		}
	}
	g.Bullets = newbullets
	g.LastUpdate = time.Now()
}

type physics struct {
	Location Vector `json:"location"`
	Velocity Vector `json:"velocity"`
}

func (p *physics) GetLocation() Vector {
	return p.Location
}

func (p *physics) GetVelocity() Vector {
	return p.Velocity
}

func (p *physics) SetVelocity(v Vector) {
	p.Velocity = v
}

func (p *physics) SetLocation(v Vector) {
	p.Location = v
}

func (p *physics) Update(dt time.Duration) {
	p.Location[0] += p.Velocity[0] * dt.Seconds()
	p.Location[1] += p.Velocity[1] * dt.Seconds()
}

type entity struct {
	Type string `json:"type"`
	physics
}
type player struct {
	entity
}

type bullet struct {
	entity
}
