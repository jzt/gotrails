package main

import (
	"runtime"

	GL "github.com/go-gl/gl/v4.1-core/gl"

	// TODO(jzt): rename this package, it's too clashy
	"github.com/jzt/gotrails/gl"
	"github.com/jzt/gotrails/mesh"
)

const (
	// TODO(jzt): take width & height as args
	width  = 800
	height = 800
)

var (
	meshes []mesh.Mesh
)

func main() {
	runtime.LockOSThread()
	window := gl.InitGlfw(width, height)
	defer gl.GLFWTerminate()

	program := gl.InitOpenGL()

	meshes = append(meshes, mesh.Triangle(program))

	for !window.ShouldClose() {
		GL.Clear(GL.COLOR_BUFFER_BIT | GL.DEPTH_BUFFER_BIT)
		GL.UseProgram(program)

		for _, m := range meshes {
			m.Render()
		}

		gl.PollEvents()
		window.SwapBuffers()
	}
}
