package transform

import (
  "image"
)

type Rotate struct {}

func init() {
  r := new(Rotate)
  RegisterTransformer("rotate", r)
}
