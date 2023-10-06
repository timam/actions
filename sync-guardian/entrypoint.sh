#!/bin/sh

ls -larth /
./sync-guardian $1 $2
echo "message=$1" >> $GITHUB_OUTPUT


