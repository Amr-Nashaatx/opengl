package window

import (
	"log"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type WindowProps struct {
	Width  int
	Height int
	Title  string
}

func init() {
	runtime.LockOSThread()
}

func CreateWindow(props *WindowProps) *glfw.Window {
	if err := glfw.Init(); err != nil {
		log.Fatalln("Failed to initialize GLFW:", err)
	}

	// 2. Configure the window
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	// 3. Create the window
	window, err := glfw.CreateWindow(props.Width, props.Height, props.Title, nil, nil)
	if err != nil {
		log.Fatalln("Failed to create window:", err)
	}
	window.MakeContextCurrent()

	// 4. Initialize OpenGL
	if err := gl.Init(); err != nil {
		log.Fatalln("Failed to initialize OpenGL:", err)
	}

	// 5. Set the viewport
	gl.Viewport(0, 0, int32(props.Width), int32(props.Height))

	return window
}
