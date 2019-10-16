#!/bin/bash

# GitHub Stuff
API_ENDPOINT=https://api.github.com
REPO_OWNER=ragecryx
REPO_NAME=bob
TOKEN=$GITHUB_STATUS_TOKEN

UpdateStatus () {
    curl --silent --output /dev/null -d "{\"state\": \"$2\", \"context\": \"continuous-integration/bob\"}" -H "Authorization: token $TOKEN" -H "Content-Type: application/json" -X POST $API_ENDPOINT/repos/$REPO_OWNER/$REPO_NAME/statuses/$1
}

# Resolve current & script path
CURRENT_PATH=`pwd`

SCRIPT_PATH="`dirname \"$0\"`"
SCRIPT_PATH="`( cd \"$SCRIPT_PATH\" && pwd )`"  # absolutized and normalized
if [ -z "$SCRIPT_PATH" ] ; then
  echo "Cannot access the directory where '$0' resides. Exiting..."
  exit 1
fi

# Get Version Info
export BUILD_BRANCH=$(git branch --no-color 2> /dev/null | sed -e '/^[^*]/d' -e "s/* \(.*\)/\1/")
export BUILD_COMMIT=$(git rev-parse HEAD 2> /dev/null | sed "s/\(.*\)/\1/")
export BUILD_DATE=$(date -u +"%Y-%m-%dT%H:%M:%SZ")

UpdateStatus $BUILD_COMMIT "pending"

make build-linux
test_result=$?

if [ $test_result -ne 0 ]; then
    UpdateStatus $BUILD_COMMIT failure
    exit 1
else
    UpdateStatus $BUILD_COMMIT success
fi
