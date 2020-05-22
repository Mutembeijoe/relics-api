package cmd

func registerRoutes() {

	// Product Group
	v1 := r.Group("/products")
	{
		v1.GET("/", app.GetProducts)
	}
	//
	// Category Group
	v2 := r.Group("/categories")
	{
		v2.GET("/", app.GetCategories)
	}
}
