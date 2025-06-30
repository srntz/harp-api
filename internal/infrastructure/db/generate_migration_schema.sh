#!/bin/bash

DIR="./**/schema.sql"

CUR_DATE=$(date +"%Y-%m-%d_%H-%M-%S")

OUTPUT_FILE="./migrations/${CUR_DATE}_${1}.sql"
>"$OUTPUT_FILE"

for file in $DIR; do
  cat "$file" >>"$OUTPUT_FILE"
done

echo "generated migration at: $OUTPUT_FILE"
