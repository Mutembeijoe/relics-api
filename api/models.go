package api

import "encoding/json"

type categoryJson struct {
	Name string `json:"name" binding:"required"`
	Options json.RawMessage `json:"options"`
}

type productJson struct {
	Name string `json:"name" binding:"required"`
	Price float32 `json:"price" binding:"required"`
	Description string `json:"description" binding:"required"`
	ImageUrl string `json:"image_url" binding:"required"`
	CategoryID uint `json:"category_id" binding:"required"`
}