#!/bin/bash 
echo "Running.."

echo $1
cp -r ./static/. ./$1
cd ./$1
pwd
ls
yarn && npx webpack --config webpack.config.js && bb components build
echo "Done!"