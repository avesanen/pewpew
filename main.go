package main

import (
	//"encoding/base64"
	"encoding/json"
	"github.com/googollee/go-socket.io"
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
	Key string `json:"key"`
}

func main() {
	g := &game{}
	g.GameArea = [2]float64{640, 480}

	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	// Set handler for player connection
	server.On("connection", func(so socketio.Socket) {
		so.Join("game")

		// Add player to the game
		p := &player{}
		p.Type = "player"
		p.SetLocation(Vector{10, 10})
		g.Players = append(g.Players, p)

		// When player disconnects, remove it from game.
		so.On("disconnection", func() {
			for i, k := range g.Players {
				if k == p {
					g.Players = append(g.Players[:i], g.Players[i+1:]...)
				}
			}
		})

		// Handler on mousedown from frontend
		so.On("mousedown", func(msg string) {
			// Unmarshal mousedown event
			var e eventMouseDown
			err := json.Unmarshal([]byte(msg), &e)
			if err != nil {
				log.Println("Can't unmarshal:", err.Error())
				return
			}

			// Set player velocity from event
			v := Vector{e.X - p.Location[0], e.Y - p.Location[1]}
			v.SetLength(200)

			// Create new bullet
			b := &bullet{}
			b.Type = "bullet"
			b.SetVelocity(v)
			b.SetLocation(p.Location)

			// Append bullet to game
			g.Bullets = append(g.Bullets, b)
			log.Println(g.Bullets)
		})

		so.On("keyup", func(msg string) {
			var e eventKeyboard
			err := json.Unmarshal([]byte(msg), &e)
			if err != nil {
				log.Println("Can't unmarshal:", err.Error())
				return
			}
			switch e.Key {
			case "a":
				p.Velocity[0] += 100
				break
			case "d":
				p.Velocity[0] -= 100
				break
			case "w":
				p.Velocity[1] += 100
				break
			case "s":
				p.Velocity[1] -= 100
				break
			}
		})

		so.On("keydown", func(msg string) {
			var e eventKeyboard
			err := json.Unmarshal([]byte(msg), &e)
			if err != nil {
				log.Println("Can't unmarshal:", err.Error())
				return
			}
			switch e.Key {
			case "a":
				p.Velocity[0] -= 100
				break
			case "d":
				p.Velocity[0] += 100
				break
			case "w":
				p.Velocity[1] -= 100
				break
			case "s":
				p.Velocity[1] += 100
				break
			}
		})
	})

	goji.Get("/socket.io/", server)
	goji.Get("/*", http.FileServer(http.Dir("./web")))
	go goji.Serve()

	for {
		g.update()
		server.BroadcastTo("game", "gamestate", string(g.getState()))
		time.Sleep(time.Second / 15)
	}
}
