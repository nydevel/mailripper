#!/bin/bash

echo 'We start to commit all changes and send them to git repository'

git add --all
git commit -a -m 'new changes'
git push

echo 'Your changes was send to git'
