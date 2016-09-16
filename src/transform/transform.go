package transform

import "image"

type Transformer interface {
    Transform(image.Image) image.Image
}
