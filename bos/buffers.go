package bos

import "github.com/go-gl/gl/v4.1-core/gl"

type bufferId = uint32
type unbindBuffer = func()

func CreateAndBindVBO(vertices []float32) (bufferId, unbindBuffer) {
	var VBO uint32
	gl.GenBuffers(1, &VBO)
	gl.BindBuffer(gl.ARRAY_BUFFER, VBO)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	unbind := func() {
		gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	}
	return VBO, unbind
}

type indexArray = []uint32

func CreateAndBindEBO(indicies indexArray) (bufferId, unbindBuffer) {
	var EBO uint32
	gl.GenBuffers(1, &EBO)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, EBO)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indicies)*4, gl.Ptr(indicies), gl.STATIC_DRAW)

	unbind := func() {
		gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0)
	}
	return EBO, unbind
}

func CreateAndBindVAO() (bufferId, unbindBuffer) {
	var VAO uint32
	gl.GenVertexArrays(1, &VAO)
	gl.BindVertexArray(VAO)

	unbind := func() {
		gl.BindVertexArray(0)
	}
	return VAO, unbind
}
