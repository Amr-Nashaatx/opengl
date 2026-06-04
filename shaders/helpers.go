package shaders

import (
	"fmt"

	"github.com/go-gl/gl/v4.1-core/gl"
)

func IsShaderCompileSuccess(shaderId uint32, shaderName string) error {
	var success int32
	infoLog := make([]byte, 512)
	gl.GetShaderiv(shaderId, gl.COMPILE_STATUS, &success)
	if success == 0 {
		gl.GetShaderInfoLog(shaderId, 512, nil, &infoLog[0])
		return fmt.Errorf("Error compiling %s: %v", shaderName, string(infoLog))
	}

	return nil
}

func IsShaderLinkSuccess(programId uint32) error {
	var success int32
	infoLog := make([]byte, 512)

	gl.GetProgramiv(programId, gl.LINK_STATUS, &success)
	if success == 0 {
		gl.GetProgramInfoLog(programId, 512, nil, &infoLog[0])
		return fmt.Errorf("Error linking shader program: %v", string(infoLog))
	}
	return nil
}
