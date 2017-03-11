package emojimoji

import (
  "gopkg.in/gographics/imagick.v3/imagick"
  "fmt"
  "unicode/utf8"
)

type Moji struct {
  file string
  linespacing float64
  size float64
  x float64
  y float64
  msg string
  color string
}

func NewMoji(msg string, t int) Moji {
  fontPath := "files/ipaexm.ttf"
  if t == 1 {
    fontPath = "files/ipaexg.ttf"
  }
  // TODO この辺はあとで綺麗にする

  r := []rune(msg)


  moji := Moji {
    file: fontPath,
    linespacing: 0.9,
    msg: string(r[0:2]) + "\n" + string(r[2:4]),
    color: "#0000ff",
  }

  length := utf8.RuneCountInString(msg)

  if length == 1 {
    moji.size = float64(120 * 72 / 96)
    moji.x = float64(6 * 72 / 96)
    moji.y = float64(146 * 72 / 96)
  } else if length == 2 {
    moji.size = float64(62 * 72 / 96)
    moji.x = float64(6 * 4 / 96)
    moji.y = float64(116 * 116 / 96)
  } else {
    moji.size = float64(62 * 72 / 96)
    moji.x = float64(6 * 72 / 96)
    moji.y = float64(78 * 72 / 96)
  }

  return moji
}

func Create(msg string, t int) {
  imagick.Initialize()
  defer imagick.Terminate()

  moji := NewMoji(msg, t)

  mw := imagick.NewMagickWand()
  defer mw.Destroy()
  readImage(mw, "files/template.png")

  dw := imagick.NewDrawingWand()
  dw.SetFont(moji.file)
  dw.SetFontSize(moji.size)
  dw.Annotation(moji.x, moji.y, moji.msg)
  dw.SetTextInterlineSpacing(moji.linespacing)


  err := mw.DrawImage(dw)
  if err == nil {
    mw.WriteImage("./out.png")
  } else {
    fmt.Println("DrawImage Error.")
    fmt.Println(err)
  }
  debug(moji)

}

func readImage(mw *imagick.MagickWand, path string) {
  err := mw.ReadImage(path)
  if err != nil {
    fmt.Println(err)
  }
}

func debug(moji Moji) {
  fmt.Println("x:",moji.x)
  fmt.Println("y:",moji.y)
  fmt.Println("fontSize:", moji.size)
  fmt.Println("msg:",moji.msg)
  fmt.Println("message length:",utf8.RuneCountInString(moji.msg))
}
