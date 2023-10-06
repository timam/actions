#!/bin/sh -l

echo $GITHUB_RUN_ID
echo "message=$1" >> $GITHUB_OUTPUT


