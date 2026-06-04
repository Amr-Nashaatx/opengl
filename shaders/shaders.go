package shaders

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type Shader struct {
	programID uint32
}

func (shader *Shader) New() error {
	var vertexCode string
	var fragmentCode string

	workDir, wdErr := os.Getwd()
	if wdErr != nil {
		return fmt.Errorf("Cannot get working dir %w\n", wdErr)
	}
	// Read shaders files at glsl directory
	vPath := filepath.Join(workDir, "glsl/vertex.glsl")
	fPath := filepath.Join(workDir, "glsl/fragment.glsl")

	vFile, vFileErr := os.ReadFile(vPath)
	if vFileErr != nil {
		return fmt.Errorf("Error reading vertex shader code %w\n", vFileErr)
	}

	fFile, fFileErr := os.ReadFile(fPath)
	if fFileErr != nil {
		return fmt.Errorf("Error reading fragment shader code %w\n", fFileErr)
	}

	vertexCode = string(vFile)
	fragmentCode = string(fFile)

	/* So far we have the content of shader files as Go strings.
	we need to convert them to C strings so openGl can handle them
	*/
	vShaderCode, freeVert := gl.Strs(vertexCode)
	defer freeVert()
	fShaderCode, freeFrag := gl.Strs(fragmentCode)
	defer freeFrag()

	// SHADER COMPILATION
	var vertexId uint32
	var fragmentId uint32

	// 1- create shader objects and store their id's
	vertexId = gl.CreateShader(gl.VERTEX_SHADER)
	fragmentId = gl.CreateShader(gl.FRAGMENT_SHADER)

	/* 2- Pass the shader source code to their corresponding object.
	we send C-style string contents
	*/
	gl.ShaderSource(vertexId, 1, vShaderCode, nil)
	gl.ShaderSource(fragmentId, 1, fShaderCode, nil)

	// 3- Compile the code inside each object
	gl.CompileShader(vertexId)
	gl.CompileShader(fragmentId)

	/* 4- Check success of shader compilation, returning error in case we find one,
	along with a log message
	*/
	vertexErr := IsShaderCompileSuccess(vertexId, "Vertex Shader")
	if vertexErr != nil {
		return vertexErr
	}

	fragmentErr := IsShaderCompileSuccess(fragmentId, "Fragment Shader")
	if fragmentErr != nil {
		return fragmentErr
	}

	/* 5- Create a new shader program object, to which we attach the compiled shaders from earlier.
	After attaching our shaders we issue a link command to link shaders together into a program.
	*/
	shader.programID = gl.CreateProgram()
	gl.AttachShader(shader.programID, vertexId)
	gl.AttachShader(shader.programID, fragmentId)
	gl.LinkProgram(shader.programID)

	// 6- Check for any linking errors
	linkErr := IsShaderLinkSuccess(shader.programID)
	if linkErr != nil {
		return linkErr
	}

	// We can safely delete shader objects after the program is linked.
	gl.DeleteShader(vertexId)
	gl.DeleteShader(fragmentId)
	return nil
}

func (shader *Shader) Use() {
	gl.UseProgram(shader.programID)
}

func (shader *Shader) SetBoolUniform(name string, value bool) error {
	uniformLoc, err := shader.getUniformLoc(name)
	if err != nil {
		return err
	}
	gl.Uniform1i(uniformLoc, int32(boolToInt(value)))
	return nil
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
func (shader *Shader) SetIntUniform(name string, value int) error {
	uniformLoc, err := shader.getUniformLoc(name)
	if err != nil {
		return err
	}
	gl.Uniform1i(uniformLoc, int32(value))
	return nil
}
func (shader *Shader) SetFloatUniform(name string, value float32) error {
	uniformLoc, err := shader.getUniformLoc(name)
	if err != nil {
		return err
	}
	gl.Uniform1f(uniformLoc, value)
	return nil
}
func (shader *Shader) getUniformLoc(name string) (int32, error) {
	cName := gl.Str(name + "\x00")
	uniformLoc := gl.GetUniformLocation(shader.programID, cName)
	if uniformLoc == -1 {
		return -1, fmt.Errorf("Could not get uniform location")
	}
	return uniformLoc, nil
}

func (shader *Shader) SetMat4Uniform(name string, value mgl32.Mat4) error {
	uniformLoc, err := shader.getUniformLoc(name)
	if err != nil {
		return err
	}

	gl.UniformMatrix4fv(uniformLoc, 1, false, &value[0])
	return nil
}
