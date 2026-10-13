#!/bin/bash
mkdir -p resultados/saidas

for v in v0 v1 v2 v3; do
	for i in 1 2 3 4 5; do
		echo "$v execucao $i"
		go run -trimpath ./$v > resultados/saidas/${v}_exec$i.txt 2>&1
	done
done
