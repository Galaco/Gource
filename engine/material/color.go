package material

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

type Color struct {
	Material
}

func (error *Color) Format() uint32 {
	return gl.RGB
}

func (error *Color) PixelDataForFrame(frame int) []byte {
	return error.rawColourData
}

func (error *Color) Finish() {
	gl.GenTextures(1, &error.Buffer)

	error.bindInternal(gl.TEXTURE0)
}

func (error *Color) bindInternal(textureSlot uint32) {
	gl.GenTextures(1, &error.Buffer)
	gl.ActiveTexture(textureSlot)
	gl.BindTexture(gl.TEXTURE_2D, error.Buffer)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)

	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(error.width),
		int32(error.height),
		0,
		error.Format(),
		gl.UNSIGNED_BYTE,
		gl.Ptr(error.rawColourData))
}

func NewError(name string) *Color {
	mat := Color{}

	mat.width = 8
	mat.height = 8
	mat.filePath = name
	mat.rawColourData = []uint8{
		255, 0, 255,
		255, 0, 255,
		255, 0, 255,
		255, 0, 255,
		0, 0, 0,
		0, 0, 0,
		0, 0, 0,
		0, 0, 0,

		255, 0, 255,
		255, 0, 255,
		255, 0, 255,
		255, 0, 255,
		0, 0, 0,
		0, 0, 0,
		0, 0, 0,
		0, 0, 0,

		255, 0, 255,
		255, 0, 255,
		255, 0, 255,
		255, 0, 255,
		0, 0, 0,
		0, 0, 0,
		0, 0, 0,
		0, 0, 0,

		255, 0, 255,
		255, 0, 255,
		255, 0, 255,
		255, 0, 255,
		0, 0, 0,
		0, 0, 0,
		0, 0, 0,
		0, 0, 0,

		0, 0, 0,
		0, 0, 0,
		0, 0, 0,
		0, 0, 0,
		255, 0, 255,
		255, 0, 255,
		255, 0, 255,
		255, 0, 255,

		0, 0, 0,
		0, 0, 0,
		0, 0, 0,
		0, 0, 0,
		255, 0, 255,
		255, 0, 255,
		255, 0, 255,
		255, 0, 255,

		0, 0, 0,
		0, 0, 0,
		0, 0, 0,
		0, 0, 0,
		255, 0, 255,
		255, 0, 255,
		255, 0, 255,
		255, 0, 255,

		0, 0, 0,
		0, 0, 0,
		0, 0, 0,
		0, 0, 0,
		255, 0, 255,
		255, 0, 255,
		255, 0, 255,
		255, 0, 255,
	}

	mat.Finish()

	return &mat
}