package entity

import "github.com/google/uuid"

// Image type
type Image struct {
	ID        string  `json:"id"`
	CompanyID string  `json:"companyid"`
	Company   Company `json:"company"`
	URL       string  `json:"url"`
}

func (image *Image) Create() {
	if image.ID == "" {
		image.ID = uuid.New().String()
	}

	db.Create(&image)
}

func (image *Image) Update() {
	db.Save(&image)
}

func (image *Image) Delete() {
	db.Delete(&image)
}

func (image *Image) Read(id string) *Image {
	image = &Image{}

	db.Where("id = ?", id).First(&image)

	return image
}

func (image *Image) ReadByComapny(companyID string) []Image {
	var images []Image
	db.Where("company_id = ?", companyID).Find(&images)

	return images
}
