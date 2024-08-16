package property

import "time"

type Property struct {
	UUID        string    `json:"uuid" bson:"_id,omitempty"`
	Title       string    `bson:"title,omitempty" json:"title,omitempty"`
	Description string    `bson:"description,omitempty" json:"description,omitempty"`
	Location    string    `bson:"location,omitempty" json:"location,omitempty"`
	Price       float64   `bson:"price,omitempty" json:"price,omitempty"`
	Bedrooms    int       `bson:"bedrooms,omitempty" json:"bedrooms,omitempty"`
	Bathrooms   int       `bson:"bathrooms,omitempty" json:"bathrooms,omitempty"`
	Images      []string  `bson:"images,omitempty" json:"images,omitempty"` // Массив ссылок на изображения
	CreatedAt   time.Time `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt   time.Time `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}

type CreatePropertyDTO struct {
	Title       string   `bson:"title,omitempty" json:"title,omitempty"`
	Description string   `bson:"description,omitempty" json:"description,omitempty"`
	Location    string   `bson:"location,omitempty" json:"location,omitempty"`
	Price       float64  `bson:"price,omitempty" json:"price,omitempty"`
	Bedrooms    int      `bson:"bedrooms,omitempty" json:"bedrooms,omitempty"`
	Bathrooms   int      `bson:"bathrooms,omitempty" json:"bathrooms,omitempty"`
	Images      []string `bson:"images,omitempty" json:"images,omitempty"` // Массив ссылок на изображения
}

func NewProperty(dto CreatePropertyDTO) Property {
	return Property{
		Title:       dto.Title,
		Description: dto.Description,
		Location:    dto.Location,
		Price:       dto.Price,
		Bedrooms:    dto.Bedrooms,
		Bathrooms:   dto.Bathrooms,
		Images:      dto.Images,
	}
}
