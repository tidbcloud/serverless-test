#!/usr/bin/env bash
# create_probe_user_all.sh
set -euo pipefail

USER="probe"
PASSWORD="fake_password"
ROLE="role_admin"

CLUSTERS=(
  10314062803745678922
  10738148427695470632
  10801139632124539502
  10542482610742827921
  10925915939244576672
  10534205190410845784
  10959546687428216535
)

for cid in "${CLUSTERS[@]}"; do
  echo "Creating SQL user '${USER}' on cluster ${cid} ..."
  ticloud s sql-user create \
    -c "${cid}" \
    --user "${USER}" \
    --password "${PASSWORD}" \
    --role "${ROLE}"
done

echo "All users created successfully."