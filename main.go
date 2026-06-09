package main

import (
	"unsafe"

	"github.com/Amr-Nashaatx/opengl/glbuffers"
	"github.com/Amr-Nashaatx/opengl/shaders"
	"github.com/Amr-Nashaatx/opengl/textures"
	"github.com/Amr-Nashaatx/opengl/window"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
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

	glbuffers.CreateAndBindVAO()
	glbuffers.CreateAndBindVBO(vertices)
	glbuffers.CreateAndBindEBO(indicies)
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

	// Transformation --> Model -> View -> Projection
	model := mgl32.Ident4()
	projection := mgl32.Perspective(45, float32(wndProps.Width)/float32(wndProps.Height), 0.1, 100.0)
	cameraPos := mgl32.Vec3{0, 0, 3}
	cameraFront := mgl32.Vec3{0, 0, -1}
	cameraUp := mgl32.Vec3{0, 1, 0}

	// Enable Depth Testing
	gl.Enable(gl.DEPTH_TEST)
	// gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE) --> for wireframe mode
	// Textures
	textures.LoadTexture("./textures/container.jpg", 0)
	textures.LoadTexture("./textures/awesomeface.png", 1)

	// plane Positions:
	planePositions := []mgl32.Vec3{
		{0.0, 0.0, 0.0},
		{2.0, 5.0, -15.0},
		{-1.5, -2.2, -2.5},
		{-3.8, -2.0, -12.3},
		{2.4, -0.4, -3.5},
		{-1.7, 3.0, -7.5},
		{1.3, -2.0, -2.5},
		{1.5, 2.0, -2.5},
		{1.5, 0.2, -1.5},
		{-1.3, 1.0, -1.5},
	}
	// Frame delta-time
	lastFrame := float32(0)
	cameraSpeed := float32(2.5)
	// 6. The render loop
	for !wnd.ShouldClose() {
		currentFrame := float32(glfw.GetTime())
		deltaTime := currentFrame - lastFrame
		lastFrame = currentFrame

		speed := cameraSpeed * deltaTime
		// Check for keyboard/mouse events
		glfw.PollEvents()

		if wnd.GetKey(glfw.KeyW) == glfw.Press {
			cameraPos = cameraPos.Add(cameraFront.Mul(speed))
		}
		if wnd.GetKey(glfw.KeyS) == glfw.Press {
			cameraPos = cameraPos.Sub(cameraFront.Mul(speed))
		}
		if wnd.GetKey(glfw.KeyA) == glfw.Press {
			right := cameraFront.Cross(cameraUp).Normalize()
			cameraPos = cameraPos.Sub(right.Mul(speed))
		}
		if wnd.GetKey(glfw.KeyD) == glfw.Press {
			right := cameraFront.Cross(cameraUp).Normalize()
			cameraPos = cameraPos.Add(right.Mul(speed))
		}
		//ear the screen with a dark green-ish color
		gl.ClearColor(0.2, 0.3, 0.3, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		shader.Use()
		shader.SetIntUniform("Texture1", 0)
		shader.SetIntUniform("Texture2", 1)

		shader.SetMat4Uniform("projection", projection)
		view := mgl32.LookAtV(cameraPos, cameraPos.Add(cameraFront), cameraUp)
		shader.SetMat4Uniform("view", view)
		for i := range 10 {
			model = mgl32.Translate3D(planePositions[i].X(), planePositions[i].Y(), planePositions[i].Z())
			angle := float32(20) * float32(i)
			rotation := mgl32.HomogRotate3D(angle, mgl32.Vec3{1, 0.3, 0.5})
			model = model.Mul4(rotation)
			shader.SetMat4Uniform("model", model)
			gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, nil)
		}
		// gl.DrawArrays(gl.TRIANGLES, 0, 36)

		// Show what we drew
		wnd.SwapBuffers()
	}
}
