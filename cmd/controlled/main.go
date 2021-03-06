// Package main sets up the demo world.
package main

import (
	"fmt"
	"image/color"
	"time"

	"github.com/ByteArena/box2d"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"

	"github.com/jemgunay/evolutionary-driving/box"
	"github.com/jemgunay/evolutionary-driving/car"
)

func main() {
	pixelgl.Run(start)
}

func start() {
	// create window config
	cfg := pixelgl.WindowConfig{
		Title:     "Car Demo",
		Bounds:    pixel.R(0, 0, 1080, 720),
		VSync:     false,
		Resizable: true,
	}

	// create window
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		fmt.Printf("failed create new window: %s\n", err)
		return
	}

	// create Box2D world
	world := box2d.MakeB2World(box2d.MakeB2Vec2(0, 0))
	winCentre := win.Bounds().Center()

	// enable contact filter
	world.SetContactFilter(&box2d.B2ContactFilter{})

	// create ground
	box.MainGround = box.NewGround(&world, winCentre, win.Bounds().Size())

	// create car
	mainCar := car.NewCar(&world, winCentre, pixel.V(38, 80))

	// create wall props
	walls := []*box.Wall{
		box.NewWall(&world, pixel.V(winCentre.X, win.Bounds().Min.Y), pixel.V(win.Bounds().W(), 30)), // bottom
		box.NewWall(&world, pixel.V(winCentre.X, win.Bounds().Max.Y), pixel.V(win.Bounds().W(), 30)), // top
		box.NewWall(&world, pixel.V(win.Bounds().Min.X, winCentre.Y), pixel.V(30, win.Bounds().H())), // left
		box.NewWall(&world, pixel.V(win.Bounds().Max.X, winCentre.Y), pixel.V(30, win.Bounds().H())), // right
	}

	// create crates
	crateSize := pixel.V(50, 50)
	crates := []*box.Crate{
		box.NewCrate(&world, pixel.V(winCentre.X, win.Bounds().Min.Y+250), crateSize),
		box.NewCrate(&world, pixel.V(winCentre.X-30, win.Bounds().Min.Y+190), crateSize),
		box.NewCrate(&world, pixel.V(winCentre.X+30, win.Bounds().Min.Y+190), crateSize),
		box.NewCrate(&world, pixel.V(winCentre.X-60, win.Bounds().Min.Y+130), crateSize),
		box.NewCrate(&world, pixel.V(winCentre.X+60, win.Bounds().Min.Y+130), crateSize),
		box.NewCrate(&world, pixel.V(winCentre.X, win.Bounds().Min.Y+130), crateSize),
	}

	// limit update cycles FPS
	frameRateLimiter := time.Tick(time.Second / 120)
	prevTimestamp := time.Now()
	imd := imdraw.New(nil)

	// main game loop
	for !win.Closed() {
		dt := float64(time.Since(prevTimestamp).Nanoseconds()) / 1e6
		prevTimestamp = time.Now()

		// handle keyboard input
		if win.JustPressed(pixelgl.KeyEscape) {
			return
		}
		if win.JustPressed(pixelgl.KeyR) {
			mainCar.Destroy()
			mainCar = car.NewCar(&world, winCentre, pixel.V(38, 80))
		}

		if win.Pressed(pixelgl.KeyA) && !win.Pressed(pixelgl.KeyD) {
			mainCar.SetSteerState(car.SteerLeft)
		} else if win.Pressed(pixelgl.KeyD) && !win.Pressed(pixelgl.KeyA) {
			mainCar.SetSteerState(car.SteerRight)
		} else {
			mainCar.SetSteerState(car.SteerNone)
		}

		if win.JustPressed(pixelgl.KeyQ) {
			newDir := mainCar.ToggleDirection()
			fmt.Println(newDir)
		}

		if win.Pressed(pixelgl.KeyW) {
			mainCar.Accelerating = true
		} else {
			mainCar.Accelerating = false
		}
		if win.Pressed(pixelgl.KeyS) {
			mainCar.Braking = true
		} else {
			mainCar.Braking = false
		}

		mainCar.Update(&world, dt)

		world.Step(dt/1000.0, 8, 3)
		world.ClearForces()

		// draw window
		win.Clear(color.White)
		imd.Clear()
		for _, wall := range walls {
			wall.Draw(imd)
		}
		for _, crate := range crates {
			crate.Draw(imd)
		}
		mainCar.Draw(imd)
		imd.Draw(win)
		win.Update()

		<-frameRateLimiter
	}
}
