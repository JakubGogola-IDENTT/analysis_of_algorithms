#!/bin/bash

time go run main.go -testMode
time go run main.go -testMode -withRepetitions
python3 plots.py