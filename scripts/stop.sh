#!/bin/bash
set -eux

pkill -f cosmappd &> /dev/null || true
