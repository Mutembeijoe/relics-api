package api

type categoryJson struct {
	Name string `json:"name" binding:"required"`
}

//type productJson struct {
//	Name string `json:"name" binding:"required"`
//
//
//}