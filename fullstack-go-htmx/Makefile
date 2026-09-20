build: 
	@go build -o ./bin/main ./cmd/app

run: build
	./bin/main

test:
	@go test ./...

templ: 
	@templ generate

css: 
	@./bin/tailwindcss -i ./views/css/app.css -o ./public/styles.css

dev:
	air