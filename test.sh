#!/usr/bin/env bash
set -euo pipefail

set -x

endpoint="${ENDPOINT:?Must set ENDPOINT environment variable}"
payload_file="${1:-testdata/ping.json}"
event="${2:-ping}"

payload=$(<"${payload_file}")

webhook_secret="$(<secret.txt)"

webhook_signature=$(echo -n "${payload}" | openssl dgst -sha256 -hmac "${webhook_secret}" -binary | xxd -p |tr -d '\n')

curl \
  --header "X-GitHub-Event: ${event}" \
  --header "X-GitHub-Delivery: nil" \
  --header "X-Hub-Signature-256: sha256=${webhook_signature}" \
  --header "Content-Type: application/json" \
  --data "${payload}" \
  -L \
    "${endpoint}"
