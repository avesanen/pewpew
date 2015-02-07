package main

import (
	"time"
)

type bullet struct {
	Id string `json:"id"`
	Circle
	FiredBy string `json:"firedBy"`
}

func (b *bullet) update(g *game, dt time.Duration) {
	d := b.Position.Add(b.Velocity.Mul(dt.Seconds()))
	line := Line{A: b.Position, B: d}

	for _, ct := range g.Tiles {
		pCol := ct.Intersect(line)
		for _, c := range pCol {
			col := &collision{}
			col.Collider = b.Id
			col.Target = ct.Id
			col.Position = c
			g.Collisions = append(g.Collisions, col)
		}
	}

	// Test every player for collision
	for _, cp := range g.Players {
		if cp.Id != b.FiredBy {
			pCol := cp.Intersect(line)
			for _, c := range pCol {
				col := &collision{}
				col.Collider = b.Id
				col.Target = cp.Id
				col.Position = c
				g.Collisions = append(g.Collisions, col)
			}
			if len(pCol) > 0 {
				b.Position = pCol[0]
				b.Velocity = &Vector{0, 0}
			}
		}
	}
	b.Position.X += b.Velocity.X * dt.Seconds()
	b.Position.Y += b.Velocity.Y * dt.Seconds()
}
