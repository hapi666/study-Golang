#!/bin/bash
cd resume #repository name
git pull
pwd
cd ..
pwd
files=$(ls resume)
for filename in ${files}
do
    cp -r resume/${filename} ${filename}
    echo "now"
    cd ${filename}
    pwd
    git add .
    git commit -m 'test code'
    git push origin master
done

