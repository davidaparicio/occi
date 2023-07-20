#!/usr/bin/env bash

# FOR DEV
goreleaser build --single-target --snapshot -f .goreleaser.yml --rm-dist