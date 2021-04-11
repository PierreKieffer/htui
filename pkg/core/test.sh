#!/bin/bash
curl -n -v -X POST https://api.heroku.com/apps/pandascore-ms-staging/log-sessions \
-d '{
"dyno": "worker",
"lines": 30,
"source": "app",
"tail": true
}' \
-H "Content-Type: application/json" \
-H "Accept: application/vnd.heroku+json; version=3"

