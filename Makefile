week:
	@go run main.go events week

today:
	@go run main.go events today

build:
	@go build -o gcal main.go