OUTPUT_DIR := output

.PHONY: render
render:
	rm -rf $(OUTPUT_DIR) && mkdir $(OUTPUT_DIR)
	go run hello.go > $(OUTPUT_DIR)/hello.html
