#!/usr/bin/bash

cd ../
# zip -r - assets/build >assets_full.zip
find assets/build/ -name "*.map" -type f -delete
zip -r - assets/build >assets.zip
