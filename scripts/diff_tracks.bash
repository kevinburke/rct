#!/bin/bash

go run wip/roundtrip_mine_train.go > rides/mine-train-gold-rush.roundtrip-thru-go.TD6

echo "=== game track:"
hexdump -C rides/mine-train-gold-rush.td6 | head -n 10

echo
echo "=== generated track:"
hexdump -C rides/mine-train-gold-rush.roundtrip-thru-go.TD6 | head -n 10
