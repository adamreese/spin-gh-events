#!/usr/bin/env bash
set -euo pipefail

set -x

# Usage: ./test.sh [EVENT] [PAYLOAD_FILE]

endpoint="${ENDPOINT:?Must set ENDPOINT environment variable}"
secret_key="${SECRET_KEY:?Must set SECRET_KEY environment variable}"
event="${1:-ping}"
payload_file="${2:-testdata/ping.json}"

payload=$(<"${payload_file}")

signature=$(echo -n "${payload}" | openssl dgst -sha256 -hmac "${secret_key}" -binary | xxd -p |tr -d '\n')

curl \
  --header "X-GitHub-Event: ${event}" \
  --header "X-GitHub-Delivery: nil" \
  --header "X-Hub-Signature-256: sha256=${signature}" \
  --header "Content-Type: application/json" \
  --data "${payload}" \
  --location \
  --fail \
  "${endpoint}"
