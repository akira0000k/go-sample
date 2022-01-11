#!/bin/bash
id=${1:-4}
set -x

curl -i -X DELETE localhost:8080/todos/$id

