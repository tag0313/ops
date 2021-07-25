#!/bin/bash

mkdir -p out/
mongodump --forceTableScan -h opsnft-test-1 --port 49161 -o ./out/ 
mongorestore ./out/ 
rm -rf out
