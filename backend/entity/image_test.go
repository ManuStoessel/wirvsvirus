package entity

import (
	"testing"
)

func init() {
	Initialize("")
}

func TestImage(t *testing.T) {
	company := &Company{}
	company.Create()

	companyID := company.ID

	image := &Image{URL: "testUrl", CompanyID: companyID}
	image2 := &Image{URL: "testUrl", CompanyID: companyID}

	image.Create()
	image2.Create()

	imageID := image.ID

	readImage := &Image{}
	readImage = readImage.Read(imageID)

	if readImage.ID != image.ID {
		t.Errorf("Read image should be equal created image")
	}

	image.URL = "testURL2"

	image.Update()

	readImage = readImage.Read(imageID)

	if readImage.URL != "testURL2" {
		t.Errorf("Image should be changed")
	}

	images := readImage.ReadByComapny(companyID)

	if len(images) != 2 {
		t.Errorf("There should be 2 images")
	}

	image.Delete()

	readImage = readImage.Read(imageID)

	if readImage.ID != "" {
		t.Errorf("Image should not exist anymore")
	}
}
