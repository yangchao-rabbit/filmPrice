package main

import "filmPrice/cmd"

// @title FilmPrice
// @version 1.0
// @description 胶片比价系统

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// @BasePath /api
func main() {
	cmd.Execute()
}
