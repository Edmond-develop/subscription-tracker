package main

import (
	internal "github.com/Edmond-develop/subscription-tracker/internal/config"
)

func main() {
	cfg := internal.LoadConfig()
	_ = cfg
}
