package main

import (
	"encoding/json"
	"github.com/avesanen/vector"
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
	g.GameArea = [2]float64{640, 480}

	server := websocks.NewServer()
	server.OnConnect(func(c *websocks.Conn) {
		// Add player to the game
		p := &player{}
		p.Type = "player"
		p.SetLocation(vector.Vector{100, 100})
		g.Players = append(g.Players, p)

		// When player disconnects, remove it from game.
		c.On("disconnect", func(m websocks.Msg) {
			for i, k := range g.Players {
				if k == p {
					g.Players = append(g.Players[:i], g.Players[i+1:]...)
				}
			}
		})

		c.On("mousedown", func(m websocks.Msg) {
			// Unmarshal mousedown event
			var e eventMouseDown
			err := json.Unmarshal([]byte(m.Message), &e)
			if err != nil {
				log.Println("Can't unmarshal:", err.Error())
				return
			}

			// Set player velocity from event
			v := vector.Vector{e.X - p.Location[0], e.Y - p.Location[1]}
			v.SetLength(400)

			// Create new bullet
			b := &bullet{}
			b.Type = "bullet"
			b.SetVelocity(v)
			b.Shooter = p
			b.SetLocation(p.Location)

			// Append bullet to game
			g.Bullets = append(g.Bullets, b)
		})

		c.On("mouseover", func(m websocks.Msg) {
			// Unmarshal mousedown event
			var e eventMouseDown
			err := json.Unmarshal([]byte(m.Message), &e)
			if err != nil {
				log.Println("Can't unmarshal:", err.Error())
				return
			}
			p.Aiming = vector.Vector{e.X, e.Y}
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
				p.Velocity[0] += 100
			case 68:
				p.Velocity[0] -= 100
			case 87:
				p.Velocity[1] += 100
			case 83:
				p.Velocity[1] -= 100
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
				p.Velocity[0] -= 100
			case 68:
				p.Velocity[0] += 100
			case 87:
				p.Velocity[1] -= 100
			case 83:
				p.Velocity[1] += 100
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
		time.Sleep(time.Second / 15)
	}
}
