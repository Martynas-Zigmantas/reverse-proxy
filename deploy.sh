#!/usr/bin/env bash
set -euo pipefail

cd /home/ar_admin/reverse-proxy

git fetch origin main

LOCAL="$(git rev-parse HEAD)"
REMOTE="$(git rev-parse origin/main)"

# 1. If nothing changed at all → exit immediately
if [ "$LOCAL" = "$REMOTE" ]; then
  echo "$(date) - No changes"
  exit 0
fi

echo "$(date) - Changes detected"

# 2. Check what changed BEFORE updating
CHANGED_FILES=$(git diff --name-only "$LOCAL" "$REMOTE")

# 3. Update repo
git reset --hard origin/main

# 4. Only rebuild + restart if Go files changed
if echo "$CHANGED_FILES" | grep -qE '\.go$'; then
  echo "$(date) - Go files changed → building + restarting"

  go build -o proxy proxy.go
  go build -o redirect redirect.go

  sudo systemctl restart proxy
  sudo systemctl restart redirect
else
  echo "$(date) - No Go changes → skipping build and restart"
fi

echo "$(date) - Deploy complete"
