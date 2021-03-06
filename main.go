package main

import (
	"fmt"
	_ "image/jpeg"
	"runtime"
	"time"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.1/glfw"

	"sync/atomic"

	"github.com/brandonnelson3/GoPlay/camera"
	"github.com/brandonnelson3/GoPlay/input"
	"github.com/brandonnelson3/GoPlay/voxelterrain"
	"github.com/brandonnelson3/GoPlay/window"
)

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}

var fps uint32

func printFPS() {
	//log.Printf("FPS is currently %d/second", atomic.SwapUint32(&fps, 0))
	time.AfterFunc(1*time.Second, printFPS)
}

func main() {
	defer glfw.Terminate()
	go printFPS()

	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version", version)

	// Configure global settings
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.ClearColor(1.0, 1.0, 1.0, 1.0)

	//cube, err := gameobjects.NewCube()
	//if err != nil {
	//	panic(err)
	//}

	terrain, err := voxelterrain.NewTerrain()
	if err != nil {
		panic(err)
	}

	previousTime := glfw.GetTime()
	gl.ClearColor(0, 0, 0, 0)
	for !window.M.W.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		// Update
		t := glfw.GetTime()
		elapsed := t - previousTime
		previousTime = t

		input.M.RunKeys(float32(elapsed))

		camera.C.Update(elapsed)

		//cube.Render()
		terrain.Render()

		atomic.AddUint32(&fps, 1)

		// Maintenance
		window.M.W.SwapBuffers()
		glfw.PollEvents()

		window.M.W.SetCursorPos(float64(window.M.Width)/2, float64(window.M.Height)/2)
	}
}
