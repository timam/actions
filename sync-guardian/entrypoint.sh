#!/bin/sh -l

echo $GITHUB_RUN_ID
echo "message=$GITHUB_RUN_ID" >> $GITHUB_OUTPUT


