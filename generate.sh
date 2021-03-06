#!/bin/bash

set -euo pipefail
cd "$(dirname "$0")"

VPP_REPO_URL=${VPP_REPO_URL-'https://github.com/FDio/vpp'}

# Update vpp
[[ -d vpp ]] || git clone ${VPP_REPO_URL} vpp
cd vpp
CUR_BRANCH=$(git branch | grep \* | cut -d ' ' -f2 2>/dev/null)
VPP_VERSION=$(git describe --always --tags)
[[ ${CUR_BRANCH} == "master" ]] && git pull
cd ../
rm -rf ./binapi/*

echo "Generating for VPP ${VPP_VERSION}"

# Generate .api.json files
find vpp -name \*.api -printf "echo %p \
 && vpp/src/tools/vppapigen/vppapigen --includedir vpp/src \
 --input %p --output binapi/%f.json JSON\n" | xargs -0 sh -c

# Generate Go code
go get -v github.com/FDio/govpp/cmd/binapi-generator
binapi-generator --input-dir=binapi --output-dir=binapi

# Store VPP version to file
echo ${VPP_VERSION} > binapi/VPP_VERSION
