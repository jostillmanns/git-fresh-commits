#!/bin/bash

ORIG_PWD=$PWD

mkdir -p /tmp/git-test/orig
cd /tmp/git-test/orig

git init
echo foo > foo
git add foo
git commit -am "initial commit"

cd /tmp/git-test
mkdir -p /tmp/git-test/remote
cd /tmp/git-test/remote
git init --bare

cd /tmp/git-test/orig
git remote add origin /tmp/git-test/remote
git push origin master

cd /tmp/git-test
git clone /tmp/git-test/remote developer
cd developer

git checkout -b feature
echo bar > bar
git add bar
git commit -am "remote commit"

git push origin feature:feature

cd $ORIG_PWD
