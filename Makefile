BINARY_NAME = urlshortener

build:
	@echo "🔨 Building the binary..."
	go build -o $(BINARY_NAME)
	@echo "✅ Build complete!"

run: build
	@echo "🚀 Running the CLI..."
	./$(BINARY_NAME) --url https://equitas.cc

clean:
	@echo "🧹 Cleaning up..."
	rm -f $(BINARY_NAME)
	@echo "✅ Clean complete!"

watch:
	@echo "👀 Watching for file changes..."
	@while true; do \
		inotifywait -e modify *.go; \
		clear; \
		make run; \
	done
