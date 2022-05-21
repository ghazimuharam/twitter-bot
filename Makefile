init-app:
	mkdir development
	cp ./files/production/twitter-bot.main.production.json ./development/twitter-bot.main.development.json

test:
	go test ./... -cover -count=1

run:
	go run cmd/main.go

build:
	go build -o bin/app.exe cmd/main.go