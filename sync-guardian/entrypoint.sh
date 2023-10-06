#!/bin/sh -l

ls -larth /bin/sync-guardian

echo "We got $1 $2"
message=$(./bin/sync-guardian $1 $2 )
echo "message=$message" >> $GITHUB_OUTPUT


