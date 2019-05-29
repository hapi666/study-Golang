#!/bin/bash
cd resume #repository name
git pull
pwd
cd ..
pwd
files=$(ls resume)
echo ${files}
for filename in ${files}
do
    ls
    cp -r ./resume/${filename} ./${filename}
    cd ${filename}
    pwd
    ls
    git add .
    git commit -m 'test code'
    git push origin master
done

