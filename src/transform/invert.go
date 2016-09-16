package transform

import (
  "image"
  "image/color"
)

type Invert struct {}

func init() {
  i := new(Invert)
  RegisterTransformer("invert", i)
}
