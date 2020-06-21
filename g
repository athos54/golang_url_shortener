#!/bin/bash

message=$(date +%Y/%m/%d\ %H:%M)
[[ -n $@ ]] && message="$@"

git add "."
git commit -a -m "${message}"
git push origin master
