#!/usr/bin/env bash
set -euo pipefail

PROJECT_OWNER_ID="1372813089454594739"
SPEND_LIMIT="100"

regions=(
  "aws-ap-northeast-1"
  "aws-us-west-2"
  "aws-us-east-1"
  "aws-ap-southeast-1"
  "aws-eu-central-1"
  "alicloud-ap-southeast-1"
)

for r in "${regions[@]}"; do
  echo "Creating project: $r in regions/$r ..."
  ticloud s create \
    -p "$PROJECT_OWNER_ID" \
    --display-name "$r" \
    --spending-limit-monthly "$SPEND_LIMIT" \
    --region "regions/$r"
done

echo "All projects created successfully."