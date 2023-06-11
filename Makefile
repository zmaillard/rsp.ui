all: generate-signs javascript static-site

generate-signs:
	go run main.go

static-site:
	hugo --gc --minify

javascript:
	npm install