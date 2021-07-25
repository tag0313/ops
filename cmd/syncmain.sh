#!/usr/bin/env sh


SERVICES=`ls ../service/`

#for s in ${SERVICES}; do
#  mkdir -p $s;
#  cp ../service/$s/main.go ${s}/main.go
#done

for s in ${SERVICES}; do
  echo $s
done
