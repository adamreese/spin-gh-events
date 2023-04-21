#!/usr/bin/env bash
set -euo pipefail

set -x

endpoint="${ENDPOINT:?Must set ENDPOINT environment variable}"
payload_file="${1:-testdata/ping.json}"
event="${2:-ping}"

payload=$(<"${payload_file}")

webhook_secret="$(<secret.txt)"

WEBHOOK_SIGNATURE=$(echo -n "$payload" | openssl sha1 -hmac "$webhook_secret" -binary | xxd -p)
WEBHOOK_SIGNATURE_256=$(echo -n "$payload" | openssl dgst -sha256 -hmac "$webhook_secret" -binary | xxd -p |tr -d '\n')

curl \
  --header "X-GitHub-Event: ping" \
  --header "X-GitHub-Delivery: nil" \
  --header "X-Hub-Signature: sha1=$WEBHOOK_SIGNATURE" \
  --header "X-Hub-Signature-256: sha256=$WEBHOOK_SIGNATURE_256" \
  --data "$payload" \
  -L \
    http://gh-events-ar.fermyon.app
