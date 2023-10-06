#!/bin/sh -l

echo $GITHUB_RUN_ID
echo "message=$message" >> $GITHUB_OUTPUT


