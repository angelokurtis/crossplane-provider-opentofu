#!/usr/bin/env bash

set -e

# Fail if there are uncommitted changes
if ! git diff-index --quiet HEAD --; then
  echo "Error: Uncommitted changes detected. Please commit or stash them before proceeding."
  exit 1
fi

git submodule update --init --recursive
make build

SOURCE_IMAGE="build-94c9a9c2/provider-opentofu-amd64:latest"

# Use current commit SHA as tag
COMMIT_SHA="$(git rev-parse --short HEAD)"
DEST_IMAGE="kurtis/crossplane-opentofu-provider:${COMMIT_SHA}"

docker tag "$SOURCE_IMAGE" "$DEST_IMAGE"

crossplane xpkg build \
  --embed-runtime-image="$SOURCE_IMAGE" \
  --package-root ./package \
  -o provider-opentofu.xpkg

crossplane xpkg push "$DEST_IMAGE" -f provider-opentofu.xpkg
