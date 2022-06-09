package image_test

import (
	"testing"

	"swimdata.de/nuvoled/image"
)

func TestCreateImage(t *testing.T) {
	image.CreateImage()
	t.Fail()
}

func TestFreetype(t *testing.T) {
	image.CreateImageFreeType()
	t.Fail()
}
