#!/bin/bash 
echo "Running.."

echo $2
echo $2
cd $2
pwd
ls
yarn --ignore-optional && npx webpack --config webpack.config.js && bb components build
cd $1
pwd
cp -r $2/dist/. $3
rm -r $2
echo "Done!"