#!/usr/bin/env bash
set -euo pipefail

spo=--namespace=security-profiles-operator

make image
make deployments
podman push localhost/security-profiles-operator:latest ghcr.io/mhils/security-profiles-operator:latest

IMAGE=ghcr.io/mhils/security-profiles-operator:latest make deploy

k $spo rollout restart daemonset/spod
k $spo rollout restart deployment/security-profiles-operator
sleep 3
k $spo get pods
k $spo get pods