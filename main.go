package main

import (
	"fmt"

	"github.com/Amr-Nashaatx/opengl/bos"
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

	vertexShaderSource := `#version 330 core
	layout (location = 0) in vec3 aPos;
	void main() {
	    gl_Position = vec4(aPos.x, aPos.y, aPos.z, 1.0);
	}` + "\x00"
	// 1 - creaate a vertex shader object and get the identifier into vertexShader variable
	// 2 - hand the source code of shader to vertex shader object
	// 3 - compile the shader
	var vertexShader uint32
	vertexShader = gl.CreateShader(gl.VERTEX_SHADER)

	vertexShaderCsource, freeVert := gl.Strs(vertexShaderSource)
	gl.ShaderSource(vertexShader, 1, vertexShaderCsource, nil)
	gl.CompileShader(vertexShader)
	defer freeVert()

	// Catch Vertex shader compilation errors
	var success int32
	gl.GetShaderiv(vertexShader, gl.COMPILE_STATUS, &success)

	vsInfoLog := make([]byte, 512)
	if success == 0 {
		gl.GetShaderInfoLog(vertexShader, 512, nil, &vsInfoLog[0])
		fmt.Println("ERROR: Vertex shader compilation failed")
		fmt.Println(string(vsInfoLog))
	}

	// ------- FRAGMENT SHADER ---------

	fragmentShaderSource := `
		#version 330 core
		out vec4 FragColor;

		void main()
		{
 		   FragColor = vec4(1.0f, 0.5f, 0.2f, 1.0f);
		} 
	` + "\x00"
	var fragmentShader uint32
	fragmentShader = gl.CreateShader(gl.FRAGMENT_SHADER)

	fragmentShaderCsource, freeFrag := gl.Strs(fragmentShaderSource)
	gl.ShaderSource(fragmentShader, 1, fragmentShaderCsource, nil)
	gl.CompileShader(fragmentShader)
	defer freeFrag()
	// Catch fragment shader compilation errors
	var fgShaderSuccess int32
	gl.GetShaderiv(fragmentShader, gl.COMPILE_STATUS, &fgShaderSuccess)

	fgShaderInfoLog := make([]byte, 512)
	if success == 0 {
		gl.GetShaderInfoLog(fragmentShader, 512, nil, &fgShaderInfoLog[0])
		fmt.Println("ERROR: Vertex shader compilation failed")
		fmt.Println(string(fgShaderInfoLog))
	}

	var shaderProgram uint32
	shaderProgram = gl.CreateProgram()

	gl.AttachShader(shaderProgram, vertexShader)
	gl.AttachShader(shaderProgram, fragmentShader)
	gl.LinkProgram(shaderProgram)

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	// gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE) --> for wireframe mode
	// 6. The render loop
	for !wnd.ShouldClose() {
		// Check for keyboard/mouse events
		glfw.PollEvents()
		//ear the screen with a dark green-ish color
		gl.ClearColor(0.2, 0.3, 0.3, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		gl.UseProgram(shaderProgram)
		// gl.DrawArrays(gl.TRIANGLES, 0, 3)
		gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, nil)

		// Show what we drew
		wnd.SwapBuffers()
	}
}
