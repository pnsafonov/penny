#!/bin/bash
cat ./generic_max.go | ../../penny gen "NumberType=NUMBERS" > numbers_max_get.go
cat ./func_thing.go | ../../penny gen "ThisNumberType=NUMBERS" > numbers_func_thing.go
