#!/bin/sh

wget -O ./data/enwikinews-20230501-abstract.xml.gz -q --show-progress https://dumps.wikimedia.org/enwikinews/20230501/enwikinews-20230501-abstract.xml.gz
gunzip ./data/enwikinews-20230501-abstract.xml.gz