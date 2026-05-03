package shaders

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-gl/gl/v4.1-core/gl"
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

	vShaderCode, freeVert := gl.Strs(vertexCode)
	defer freeVert()
	fShaderCode, freeFrag := gl.Strs(fragmentCode)
	defer freeFrag()

	// compile shaders
	var vertexId uint32
	var fragmentId uint32
	var success int32

	infoLog := make([]byte, 512)

	vertexId = gl.CreateShader(gl.VERTEX_SHADER)
	fragmentId = gl.CreateShader(gl.FRAGMENT_SHADER)

	gl.ShaderSource(vertexId, 1, vShaderCode, nil)
	gl.ShaderSource(fragmentId, 1, fShaderCode, nil)

	gl.CompileShader(vertexId)
	gl.CompileShader(fragmentId)

	gl.GetShaderiv(vertexId, gl.COMPILE_STATUS, &success)
	if success == 0 {
		gl.GetShaderInfoLog(vertexId, 512, nil, &infoLog[0])
		return fmt.Errorf("Error compiling vertex shader %v", string(infoLog))
	}

	gl.GetShaderiv(fragmentId, gl.COMPILE_STATUS, &success)
	if success == 0 {
		gl.GetShaderInfoLog(fragmentId, 512, nil, &infoLog[0])
		return fmt.Errorf("Error compiling fragment shader %v", string(infoLog))
	}

	// Shader Program
	shader.programID = gl.CreateProgram()
	gl.AttachShader(shader.programID, vertexId)
	gl.AttachShader(shader.programID, fragmentId)
	gl.LinkProgram(shader.programID)

	gl.GetProgramiv(shader.programID, gl.LINK_STATUS, &success)
	if success == 0 {
		gl.GetProgramInfoLog(shader.programID, 512, nil, &infoLog[0])
	}

	gl.DeleteShader(vertexId)
	gl.DeleteShader(fragmentId)
	return nil
}

func (shader *Shader) Use() {
	gl.UseProgram(shader.programID)
}

func (shader *Shader) SetBoolUniform(name string, value bool) {
	gl.Uniform1i(gl.GetUniformLocation(shader.programID, gl.Str(name)), int32(boolToInt(value)))
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
func (shader *Shader) SetIntUniform(name string, value int) {
	gl.Uniform1i(gl.GetUniformLocation(shader.programID, gl.Str(name)), int32(value))
}
func (shader *Shader) SetFloatUniform(name string, value float32) {
	gl.Uniform1f(gl.GetUniformLocation(shader.programID, gl.Str(name)), value)
}
