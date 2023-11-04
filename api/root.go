package api

import (
	"fmt"
	"os"
)

type API struct {
}

func GetUrl() string {
	if os.Getenv("API_URL") == "" {
		return "https://depo.com/api"
	} else {
		return fmt.Sprintf("%s/api", os.Getenv("API_URL"))
	}
}
