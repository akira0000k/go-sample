#!/bin/bash
ttl=${1:-test}
txt=${2:-テストです。}
set -x

curl -i -X POST -H "Content-Type: application/json" -d '{"title":"'$ttl'", "content":"'$txt'"}' localhost:8080/todos/

