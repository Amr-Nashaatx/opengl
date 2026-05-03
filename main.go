package main

import (
	"github.com/Amr-Nashaatx/opengl/bos"
	"github.com/Amr-Nashaatx/opengl/shaders"
	"github.com/Amr-Nashaatx/opengl/window"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func main() {

	wndProps := &window.WindowProps{Height: 600, Width: 800, Title: "LearnOpenGL in Go"}
	wnd := window.CreateWindow(wndProps)
	defer glfw.Terminate()
	vertices := []float32{
		0.5, 0.5, 0.0,
		0.5, -0.5, 0.0,
		-0.5, -0.5, 0.0,
		-0.5, 0.5, 0.0,
	}
	indicies := []uint32{
		0, 1, 3,
		1, 2, 3,
	}

	bos.CreateAndBindVAO()

	bos.CreateAndBindVBO(vertices)
	bos.CreateAndBindEBO(indicies)
	// specify the layout of vertex attributes in bound VBO
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 12, nil)
	gl.EnableVertexAttribArray(0)

	shader := shaders.Shader{}
	if err := shader.New(); err != nil {
		panic(err)
	}

	// gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE) --> for wireframe mode
	// 6. The render loop
	for !wnd.ShouldClose() {
		// Check for keyboard/mouse events
		glfw.PollEvents()
		//ear the screen with a dark green-ish color
		gl.ClearColor(0.2, 0.3, 0.3, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		shader.Use()
		// gl.DrawArrays(gl.TRIANGLES, 0, 3)
		gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, nil)

		// Show what we drew
		wnd.SwapBuffers()
	}
}
