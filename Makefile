.PHONY: build

build: /usr/local/bin/sam build

local: /usr/local/bin/sam local start-api -n env.json

run: build make local
