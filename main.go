package main

import (
	"github.com/joho/godotenv"
	"github.com/lalizita/calendar-cli-gcp/cmd"
)

func main() {
	godotenv.Load(".env")
	cmd.Execute()
}
