BIN ?= bin

build/tinynotify:
	mkdir -p build
	go build -o $@ .
	
install: build/tinynotify
	mkdir -p $(BIN)
	cp  $< $(BIN)/$(basename $<)
