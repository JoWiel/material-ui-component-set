#!/bin/bash 
echo "Running.."

echo $1
echo $2
cd ./$1
pwd
ls
yarn && npx webpack --config webpack.config.js && bb components build
cp -r $1/dist $2
rm rf $1
echo "Done!"