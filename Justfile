# Yog-Sothoth Justfile

version := `grep -m 1 '^version = ' pyproject.toml | cut -d '"' -f 2`

# Build the CLI for local dev and output to ~/go/bin/yog
# build:
# 	go build -ldflags="-X 'src/yog_sothoth/cmd.Version={{version}}'" -o ~/go/bin/yog ./src/yog_sothoth
# 	go build -ldflags="-X 'src/yog_sothoth/cmd.Version={{version}}'" -o ./src/yog_sothoth/bin/yog ./src/yog_sothoth

# Cross-compile all platform targets into src/yog_sothoth/bin/ for PyPI packaging
build:
	#!/usr/bin/env bash
	set -e
	mkdir -p src/yog_sothoth/bin

	echo "Cross-compiling yog for all platforms..."

	targets=( \
		"linux   amd64  yog-linux-amd64" \
		"linux   arm64  yog-linux-arm64" \
		"darwin  amd64  yog-darwin-amd64" \
		"darwin  arm64  yog-darwin-arm64" \
		"windows amd64  yog-windows-amd64.exe" \
	)

	for target in "${targets[@]}"; do
		read -r goos goarch outname <<< "$target"
		echo "  → $outname"
		CGO_ENABLED=0 GOOS=$goos GOARCH=$goarch go build \
			-ldflags="-s -w -X src/yog_sothoth/cmd.Version={{version}}" \
			-o src/yog_sothoth/bin/$outname \
			./src/yog_sothoth
		if [[ "$goos" != "windows" ]]; then
			chmod +x src/yog_sothoth/bin/$outname
		fi
	done

	echo ""
	echo "Build complete. Binary sizes:"
	ls -lh src/yog_sothoth/bin/ | awk '{print "  " $5 "  " $9}'

# Remove all compiled binaries from the bin/ dir
clean-bin:
	rm -f src/yog_sothoth/bin/yog-*
	echo "Cleaned bin/"