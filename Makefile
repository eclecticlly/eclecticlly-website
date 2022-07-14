.DEFAULT_GOAL := build

dev:
# Windows
# start air -c .air-win.toml && start tailwindcss -i ./tailwind.css -o ./static/css/styles.css --watch && start goconvey
# Unix
	tailwindcss -i ./tailwind.css -o ./static/css/styles.css --watch &! goconvey &! air -c .air-unix.toml
.PHONY: dev

fmt:
	go fmt ./...
.PHONY:fmt

lint: fmt
	golint ./...
.PHONY:lint

vet: lint
	go vet ./...
.PHONY:vet

build: vet
	tailwindcss -i ./tailwind.css -o ./static/css/styles.css
	go build main.go
.PHONY:build