package main

import (
	"unsafe"

	"github.com/Amr-Nashaatx/opengl/bos"
	"github.com/Amr-Nashaatx/opengl/shaders"
	textures "github.com/Amr-Nashaatx/opengl/textures"
	"github.com/Amr-Nashaatx/opengl/window"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func main() {

	wndProps := &window.WindowProps{Height: 600, Width: 800, Title: "LearnOpenGL in Go"}
	wnd := window.CreateWindow(wndProps)
	defer glfw.Terminate()

	// shader
	shader := shaders.Shader{}
	if err := shader.New(); err != nil {
		panic(err)
	}

	// Data
	vertices := []float32{
		// positions          // colors           // texture coords
		0.5, 0.5, 0.0, 1.0, 0.0, 0.0, 1.0, 1.0, // top right
		0.5, -0.5, 0.0, 0.0, 1.0, 0.0, 1.0, 0.0, // bottom right
		-0.5, -0.5, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, // bottom let
		-0.5, 0.5, 0.0, 1.0, 1.0, 0.0, 0.0, 1.0, // top left
	}
	indicies := []uint32{
		0, 1, 3,
		1, 2, 3,
	}

	bos.CreateAndBindVAO()
	bos.CreateAndBindVBO(vertices)
	bos.CreateAndBindEBO(indicies)
	// specify the layout of vertex attributes in bound VBO

	// 1 - Position attribute
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 32, nil)
	gl.EnableVertexAttribArray(0)

	// 2 - color attribute
	gl.VertexAttribPointer(1, 3, gl.FLOAT, false, 32, unsafe.Pointer(uintptr(12)))
	gl.EnableVertexAttribArray(1)

	// 3 - texture attribute
	gl.VertexAttribPointer(2, 2, gl.FLOAT, false, 32, unsafe.Pointer(uintptr(24)))
	gl.EnableVertexAttribArray(2)

	// gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE) --> for wireframe mode
	// Textures
	textures.LoadTexture("./textures/container.jpg", 0)
	textures.LoadTexture("./textures/awesomeface.png", 1)
	// 6. The render loop
	for !wnd.ShouldClose() {
		// Check for keyboard/mouse events
		glfw.PollEvents()
		//ear the screen with a dark green-ish color
		gl.ClearColor(0.2, 0.3, 0.3, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		shader.Use()
		shader.SetIntUniform("Texture1", 0)
		shader.SetIntUniform("Texture2", 1)
		// gl.DrawArrays(gl.TRIANGLES, 0, 3)
		gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, nil)

		// Show what we drew
		wnd.SwapBuffers()
	}
}
