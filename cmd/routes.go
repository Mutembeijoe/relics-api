package cmd

func registerRoutes() {

	// Product Group
	v1 := r.Group("/products")
	{
		v1.GET("/", app.GetProducts)
		v1.GET("/:id",app.GetProductByID)

		v1.POST("/add", app.AddProduct)
	}
	//
	// Category Group
	v2 := r.Group("/categories")
	{
		v2.GET("/", app.GetCategories)

		v2.POST("/add", app.AddCategory)
	}
}
