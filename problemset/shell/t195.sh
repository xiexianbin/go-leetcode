#!/bin/bash

awk "NR==10" file.txt
awk 'NR==10{print}' file.txt
sed -n "10p" file.txt
tail -n +10 file.txt | head -1
