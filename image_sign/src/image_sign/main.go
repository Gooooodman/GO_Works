package main

import(
	"image/png"
	"os"
	"log"
	"image"
	"image/draw"
	"image/color"
	"code.google.com/p/freetype-go/freetype"
	"io/ioutil"
	"utility/process"
	"strings"
)

var (
	dataPath string = ""
)

func init() {
	rootPath, err := process.RootPath()
	if err != nil{
		log.Fatalln(err)
	}
	dataPath = rootPath + "/dat/"
}

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
	pic1 := getImage(dataPath + `pic1.png`)
	pic1Bound := pic1.Bounds()

	dst := image.NewRGBA(image.Rect(0,0, pic1Bound.Dx(), pic1Bound.Dy()))
	draw.Draw(dst, dst.Bounds(), pic1, image.ZP, draw.Src)

	// 画新的图片
	// 画一个带有字体的图片
	// https://groups.google.com/forum/?fromgroups=#!topic/golang-nuts/p4FcEuxcBbM
	ttfile := dataPath + "luximr.ttf"
	src := drawStringImage("yejianfeng", ttfile)

	mask := image.NewUniform(color.RGBA{0,0,255,50})

	draw.DrawMask(dst, dst.Bounds(), src, image.ZP, mask, image.Point{50,50}, draw.Over)

	// 保存到一个文件中去
	outputPath := dataPath + "out.png"
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

// 画一个带有text的图片
func drawStringImage(text string, fontFile string) image.Image {
	fontBytes, err := ioutil.ReadFile(fontFile)
	if err != nil {
		log.Fatalln(err)
	}

	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Fatalln(err)
	}

	fg, bg := image.Black, image.White 
	rgba := image.NewRGBA(image.Rect(0, 0, 640, 480))
	draw.Draw(rgba, rgba.Bounds(), bg, image.ZP, draw.Src)

	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(font)
	c.SetFontSize(12)
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	c.SetSrc(fg)

	// Draw the text.
	pt := freetype.Pt(10, 10+int(c.PointToFix32(12)>>8))
	for _, s := range strings.Split(text, "\r\n") {
		_, err = c.DrawString(s, pt)
		pt.Y += c.PointToFix32(12 * 1.5)
	}

	return rgba
}