#!/usr/bin/bash

cd ../
zip -r cloudreveplus-source.zip ./PlusBackend/ -x './PlusBackend/assets/node_modules/*' --exclude '**/.git/**'
