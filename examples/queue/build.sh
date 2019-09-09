#!/bin/bash
cat ./queue_generic.go | ../../penny gen "Generic=string,int" > queue_generic_gen.go
