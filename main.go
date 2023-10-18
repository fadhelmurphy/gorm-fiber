package main

import (
	"go-pagination-gorm/provider"
)

func main() {
    app := provider.ProvideService()

    app.Listen(":3000")
}