#!/bin/bash

awk '{print $1}' file.txt | xargs echo | sed 's/\n//g' && awk '{print $2}' file.txt | xargs echo | sed 's/\n//g'
