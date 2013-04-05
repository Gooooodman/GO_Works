package main

import(
	"image/png"
	"os"
	"log"
	"image"
	"image/draw"
	"image/color"
)

func main() {
	// read pic1 and pic 2 
	getImage := func(pic string) (image.Image) {
		file, err := os.Open(pic)
		if err != nil {
			log.Fatalln(err)
		}
		defer file.Close()

		img, err := png.Decode(file)
		if err != nil {
			log.Fatalln(err)
		}
		return img
	}

	// 拷贝png图片
	pic1 := getImage(`C:\Users\yejianfeng\Documents\GitHub\GO_Works\image_sign\pic1.png`)
	pic1Bound := pic1.Bounds()

	dst := image.NewRGBA(image.Rect(0,0, pic1Bound.Dx(), pic1Bound.Dy()))
	draw.Draw(dst, dst.Bounds(), pic1, image.ZP, draw.Src)

	// 画新的图片
	src := getImage(`C:\Users\yejianfeng\Documents\GitHub\GO_Works\image_sign\pic2.png`)

	mask := image.NewUniform(color.Alpha{uint8(10)})

	draw.DrawMask(dst, dst.Bounds(), src, image.ZP, mask, image.Point{50,50}, draw.Over)

	// 保存到一个文件中去
	outputPath := `C:\Users\yejianfeng\Documents\GitHub\GO_Works\image_sign\out.png`
	output, err := os.OpenFile(outputPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	defer output.Close()

	err = png.Encode(output, dst)
	if err != nil {
		log.Fatalln(err)
	}
}