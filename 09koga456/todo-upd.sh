#!/bin/bash
id=${1:-4}
ttl=${2:-test}
txt=${3:-変更テスト}
set -x

curl -i -X PUT -H "Content-Type: application/json" -d '{"title":"'$ttl'", "content":"'$txt'"}' localhost:8080/todos/$id
