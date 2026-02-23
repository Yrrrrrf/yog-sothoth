# Yog-Sothoth Justfile

# Build the CLI and output to ~/go/bin/yog
build:
	go build -o ~/go/bin/yog ./src/yog_sothoth
