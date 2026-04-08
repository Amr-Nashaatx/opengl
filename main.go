package main

import (
	"fmt"
	"log"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
	windowWidth  = 800
	windowHeight = 600
	windowTitle  = "LearnOpenGL in Go"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	// 1. Initialize GLFW
	if err := glfw.Init(); err != nil {
		log.Fatalln("Failed to initialize GLFW:", err)
	}
	defer glfw.Terminate()

	// 2. Configure the window
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	// 3. Create the window
	window, err := glfw.CreateWindow(windowWidth, windowHeight, windowTitle, nil, nil)
	if err != nil {
		log.Fatalln("Failed to create window:", err)
	}
	window.MakeContextCurrent()

	// 4. Initialize OpenGL
	if err := gl.Init(); err != nil {
		log.Fatalln("Failed to initialize OpenGL:", err)
	}

	// 5. Set the viewport
	gl.Viewport(0, 0, windowWidth, windowHeight)

	vertices := []float32{
		-0.5, -0.5, 0.0,
		0.5, -0.5, 0.0,
		0.0, 0.5, 0.0,
	}

	// create the vertex array object that stores all vertex attribute config
	var VAO uint32
	gl.GenVertexArrays(1, &VAO)

	// bind the vertex array
	gl.BindVertexArray(VAO)
	// 1 - generate a buffer object, hold its identifier in VBO var
	// 2 - Tell opengl to take the buffer created and bind it to ARRAY_BUFFER which is the vertex buffer object
	// 3 - Send the data to the created buffer which resides on GPU memory
	var VBO uint32
	gl.GenBuffers(1, &VBO)
	gl.BindBuffer(gl.ARRAY_BUFFER, VBO)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

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

	// 6. The render loop
	for !window.ShouldClose() {
		// Check for keyboard/mouse events
		glfw.PollEvents()

		// Clear the screen with a dark green-ish color
		gl.ClearColor(0.2, 0.3, 0.3, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		gl.UseProgram(shaderProgram)
		gl.BindVertexArray(VAO)
		gl.DrawArrays(gl.TRIANGLES, 0, 3)

		// Show what we drew
		window.SwapBuffers()
	}
}
