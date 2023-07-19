#!/bin/sh
#
# Create an ed25519 key with default parameters
# like recommended in GitHub docs:
#   https://docs.github.com/en/authentication/connecting-to-github-with-ssh/generating-a-new-ssh-key-and-adding-it-to-the-ssh-agent#generating-a-new-ssh-key

KEYFILE=$(basename "$0")
KEYFILE="${KEYFILE%.sh}"
ssh-keygen -t ed25519 -C "INSECURE TEST KEY" -f "$KEYFILE"
