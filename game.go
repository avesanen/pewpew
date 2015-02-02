package main

import (
	"encoding/json"
	"github.com/avesanen/vector"
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
		p.update(time.Since(g.LastUpdate))
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
		p.Targets = []vector.Vector{}
		for _, p2 := range g.Players {
			if p2 != p {
				cols := vector.CircleCollision(p.Location, p.Aiming, p2.Location, 10)
				if len(cols) > 0 {
					p.Targets = append(p.Targets, cols...)
				}
			}
		}
	}

	var newbullets []*bullet
	for _, b := range g.Bullets {
		b.Collisions = []vector.Vector{}
		for _, p := range g.Players {
			var path vector.Vector
			path[0] = b.Location[0] + b.Velocity[0]*time.Since(g.LastUpdate).Seconds()
			path[1] = b.Location[1] + b.Velocity[1]*time.Since(g.LastUpdate).Seconds()
			cols := vector.CircleCollision(b.Location, path, p.Location, 10)
			if len(cols) > 0 && b.Shooter != p {
				b.Collisions = append(b.Collisions, cols...)
			}
		}
		b.update(time.Since(g.LastUpdate))
		if b.Location[0] > 0.0 && b.Location[0] < g.GameArea[0] &&
			b.Location[1] > 0.0 && b.Location[1] < g.GameArea[1] {
			newbullets = append(newbullets, b)
		}
	}
	g.Bullets = newbullets
	g.LastUpdate = time.Now()
}

// Physics sub-entity
type physics struct {
	Location vector.Vector `json:"location"`
	Velocity vector.Vector `json:"velocity"`
}

func (p *physics) GetLocation() vector.Vector {
	return p.Location
}

func (p *physics) GetVelocity() vector.Vector {
	return p.Velocity
}

func (p *physics) SetVelocity(v vector.Vector) {
	p.Velocity = v
}

func (p *physics) SetLocation(v vector.Vector) {
	p.Location = v
}

func (p *physics) update(dt time.Duration) {
	p.Location[0] += p.Velocity[0] * dt.Seconds()
	p.Location[1] += p.Velocity[1] * dt.Seconds()
}

// Entity type
type entity struct {
	Type string `json:"type"`
	physics
}

// Player entity
type player struct {
	entity
	Aiming  vector.Vector   `json:"aiming"`
	Targets []vector.Vector `json:"targets"`
}

func (p *player) update(dt time.Duration) {
	p.entity.update(dt)
}

// Bullet entity
type bullet struct {
	entity
	Collisions []vector.Vector `json:"collisions"`
	Shooter    *player         `json:"-"`
}

func (b *bullet) update(dt time.Duration) {
	b.entity.update(dt)
}
