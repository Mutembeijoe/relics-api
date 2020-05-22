package cmd

func registerRoutes(){
	r.GET("/products", app.GetProducts)
}