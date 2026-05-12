package textures

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"image"
	_ "image/jpeg"
	"log"
	"os"
)

func loadImage(filePath string) *image.NRGBA {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalln("could not open image file", err)
	}

	img, _, decodeErr := image.Decode(file)
	if decodeErr != nil {
		log.Fatalln("error decoding image", decodeErr)
	}

	rgba := image.NewNRGBA(img.Bounds())
	for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
		for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
			rgba.Set(x, y, img.At(x, y))
		}
	}

	return rgba
}

func LoadTexture(filePath string) {
	var texture uint32
	gl.GenTextures(1, &texture)
	gl.BindTexture(gl.TEXTURE_2D, texture)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR_MIPMAP_LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)

	img := loadImage(filePath)

	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGB, int32(img.Bounds().Dx()), int32(img.Bounds().Dy()), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(img.Pix))
	gl.GenerateMipmap(gl.TEXTURE_2D)
}
