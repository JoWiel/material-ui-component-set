#!/bin/bash 
echo "Running.."

cd $2
npm install && npm i fsevents@latest -f --save-optional && npx webpack --config webpack.config.js && bb components build

cd $1
cp -r $2/dist/. $3
rm -r $2
echo "Done!"