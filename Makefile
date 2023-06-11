all: generate-signs static-site

generate-signs:
	go run main.go

static-site:
	hugo --gc --minify
