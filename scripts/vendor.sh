#!/usr/bin/env bash
# Install runtime dependencies into workflow/vendor/
# Run this after adding packages to vendor-requirements.txt.
set -euo pipefail

REPO_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
VENDOR_DIR="$REPO_ROOT/workflow/vendor"

echo "→ Installing dependencies into $VENDOR_DIR"
mkdir -p "$VENDOR_DIR"

# Clear existing vendor dir to avoid stale packages
rm -rf "${VENDOR_DIR:?}"/*

if [[ ! -f "$REPO_ROOT/vendor-requirements.txt" ]]; then
  echo "  No vendor-requirements.txt found - skipping vendor install"
  exit 0
fi

uv pip install \
  --quiet \
  --requirement "$REPO_ROOT/vendor-requirements.txt" \
  --target "$VENDOR_DIR"

echo "✓ Vendor install complete"
