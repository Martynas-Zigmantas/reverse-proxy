#!/usr/bin/env bash
set -e

cd /home/ar_admin/reverse-proxy

git fetch origin main

LOCAL="$(git rev-parse HEAD)"
REMOTE="$(git rev-parse origin/main)"

if [ "$LOCAL" = "$REMOTE" ]; then
  echo "$(date) - No changes"
  exit 0
fi

echo "$(date) - Changes detected. Pulling..."
git pull origin main

echo "$(date) - Building..."
go build -o proxy proxy.go
go build -o redirect redirect.go

echo "$(date) - Restarting services..."
sudo systemctl restart proxy
sudo systemctl restart redirect

echo "$(date) - Deploy complete"
