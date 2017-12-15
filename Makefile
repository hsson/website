DISPLAY_BOLD		:= "\033[01m"
DISPLAY_RESET		:= "\033[0;0m"
BINARY_NAME 		:= go-website
ARTIFACT_DIR 		:= build
CONFIG_FILE 		:= config.yaml

.DEFAULT_GOAL 		:= run

test:
	@echo $(DISPLAY_BOLD)"==> Testing code..."$(DISPLAY_RESET)
	go test ./... -cover

build:
	@echo $(DISPLAY_BOLD)"==> Compiling program..."$(DISPLAY_RESET)
	go build

run: build
	@echo $(DISPLAY_BOLD)"==> Generating website..."$(DISPLAY_RESET)
	./$(BINARY_NAME) -out $(ARTIFACT_DIR) -config $(CONFIG_FILE)

new: build
	@echo $(DISPLAY_BOLD)"==> Creating new post..."$(DISPLAY_RESET)
	./$(BINARY_NAME) -new

clean:
	@echo $(DISPLAY_BOLD)"==> Cleaning up artifacts..."$(DISPLAY_RESET)
	rm $(BINARY_NAME)
	rm -r $(ARTIFACT_DIR)