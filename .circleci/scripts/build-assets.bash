#!/usr/bin/env bash

OSARCH='darwin/386 darwin/amd64 linux/386 linux/amd64 linux/arm linux/arm64'

mkdir pkg
cd pkg

gox -osarch "$OSARCH" \
  -ldflags="-X github.com/aizu-go-kapro/whroom/vendor/github.com/ktr0731/go-updater/github.isGitHubReleasedBinary=true" \
  ..

for f in *; do
  mv "$f" whroom
  tar cvf "$f.tar.gz" whroom
  rm -f whroom
done
