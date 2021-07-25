#!/usr/bin/env sh

SERVICES=`ls ../service/`

for s in ${SERVICES}; do
  echo "$s";
  cp ../service/$s/conf/conf.yaml ${s}.yaml
  cp ../service/$s/conf/conf_online.yaml ${s}_online.yaml
done