package material

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

// Generic GPU material struct
type Material struct {
	filePath string
	rgb []uint8
	buffer uint32
	width int
	height int
}

// Bind this material to the GPU
func (material *Material) Bind() {
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, material.buffer)
}

// Get the filepath this data was loaded from
func (material *Material) GetFilePath() string {
	return material.filePath
}

//Get all RGB data from this material
func (material *Material) GetColourData() []uint8 {
	return material.rgb
}

// Generate the GPU buffer for this material
func (material *Material) GenerateGPUBuffer() {
	gl.GenTextures(1, &material.buffer)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, material.buffer)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGB,
		int32(material.width),
		int32(material.height),
		0,
		gl.RGB,
		gl.UNSIGNED_BYTE,
		gl.Ptr(material.rgb))

}

func NewMaterial(filepath string, rgb []uint8, width int, height int) *Material {
	return &Material{
		filePath: filepath,
		rgb: rgb,
		width: width,
		height: height,
	}
}