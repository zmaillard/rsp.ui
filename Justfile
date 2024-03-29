set dotenv-load

clean-content:
    rm -Rf content/{country,county,feature,highwayType,highway,place,sign,state}

local:
    hugo serve

build:
    go run .

all: build static-site

static-site:
	hugo --gc --minify
