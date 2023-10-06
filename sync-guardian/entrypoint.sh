#!/bin/sh -l

ls -larth /bin/sync-guardian
./bin/sync-guardian 1 1
echo "message=$message" >> $GITHUB_OUTPUT


