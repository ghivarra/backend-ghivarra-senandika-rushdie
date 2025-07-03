#!/bin/bash

source ~/.bashrc
cd app
go build -o dist/goshell
./dist/goshell