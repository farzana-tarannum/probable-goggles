#!/bin/bash

#cat products.json | jq '.[] | select(.price > 1000) | .name'

#cat products.json | jq 
cat orders.json | jq '.[] | .discount'
