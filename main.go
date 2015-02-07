package main

import (
	"encoding/json"
	"github.com/avesanen/websocks"
	"github.com/zenazn/goji"
	"log"
	"net/http"
	"time"
)

const (
	fps = 2
)

type eventMouseDown struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}
type eventKeyboard struct {
	KeyCode int `json:"key"`
}

func main() {
	g := &game{}
	g.GameArea = [2]float64{80, 60}
	g.newTile(&Vector{40, 30}, 80, 60)
	g.newTile(&Vector{40, 30}, 5, 5)

	//g.newTile(&Vector{g.GameArea[0] / 2, g.GameArea[1] / 2}, g.GameArea[0], g.GameArea[1])
	server := websocks.NewServer()

	// When a player connects.
	server.OnConnect(func(c *websocks.Conn) {
		// Add player to the game
		startingPosition := &Vector{40, 32}
		name := "greg"
		p := g.newPlayer(startingPosition, name)
		log.Println(p.Position, p.Velocity)

		// When player disconnects, remove it from game.
		c.On("disconnect", func(m websocks.Msg) {
			g.rmPlayer(p)
		})

		// When player clicks, fire a bullet.
		c.On("mousedown", func(m websocks.Msg) {
			// Unmarshal mousedown event
			var e eventMouseDown
			log.Println(string(m.Message))
			err := json.Unmarshal([]byte(m.Message), &e)
			if err != nil {
				log.Println("Can't unmarshal:", err.Error())
				return
			}
			pos := &Vector{p.Position.X, p.Position.Y}
			vel := &Vector{p.LookingAt.X, p.LookingAt.Y}
			vel = vel.Sub(pos).Normalize().Mul(50)
			g.newBullet(pos, vel, p.Id)
		})

		// When player moves mouse, aim at the location.
		c.On("mouseover", func(m websocks.Msg) {
			// Unmarshal mousedown event
			var e eventMouseDown
			err := json.Unmarshal([]byte(m.Message), &e)
			if err != nil {
				log.Println("Can't unmarshal:", err.Error())
				return
			}
			p.LookingAt = &Vector{e.X, e.Y}
		})

		c.On("keyup", func(m websocks.Msg) {
			var e eventKeyboard
			err := json.Unmarshal([]byte(m.Message), &e)
			if err != nil {
				log.Println("Can't unmarshal:", err.Error())
				return
			}
			switch e.KeyCode {
			case 65:
				p.Velocity.X += 10
			case 68:
				p.Velocity.X -= 10
			case 87:
				p.Velocity.Y += 10
			case 83:
				p.Velocity.Y -= 10
			}
		})
		c.On("keydown", func(m websocks.Msg) {
			var e eventKeyboard
			err := json.Unmarshal([]byte(m.Message), &e)
			if err != nil {
				log.Println("Can't unmarshal:", err.Error())
				return
			}
			switch e.KeyCode {
			case 65:
				p.Velocity.X -= 10
			case 68:
				p.Velocity.X += 10
			case 87:
				p.Velocity.Y -= 10
			case 83:
				p.Velocity.Y += 10
			}
		})

	})
	goji.Get("/ws/", server.WebsocketHandler)
	goji.Get("/*", http.FileServer(http.Dir("./web")))
	go goji.Serve()

	for {
		g.update()
		for _, c := range server.Conns {
			c.Send(websocks.Msg{Type: "gamestate", Message: string(g.getState())})
		}
		time.Sleep(time.Second / 10)
	}
}
