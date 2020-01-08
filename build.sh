#!/bin/sh


PLATFORMS=("linux/amd64" "darwin/amd64" "freebsd/amd64")
TARGET_DIR="dist"
rm -r $TARGET_DIR

for platform in "${PLATFORMS[@]}"
do
	platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}

	output_name='sparkle-'$GOOS'-'$GOARCH
	echo "building" $GOOS "/" $GOARCH
	target_exec=$TARGET_DIR/$output_name
	GOOS=$GOOS GOARCH=$GOARCH go build -ldflags="-w -s" -o $target_exec

	if command -v upx 1>/dev/null; then
		echo "compressing $target_exec"
		upx $target_exec
	fi
done

