#!/bin/bash

set -eu

TARGETS="
scc:contracts/OasysStateCommitmentChain.sol
sccverifier:contracts/OasysStateCommitmentChainVerifier.sol
"

OUTPUT_DIR="$(cd $(dirname $0)/../../contracts; pwd)"

for target in $TARGETS; do
    pkg=$(echo $target | cut -d: -f1)
    contract=$(echo $target | cut -d: -f2)
    pkgdir="${OUTPUT_DIR}/${pkg}"
    workdir=$(mktemp -d)

    npx solc \
        --base-path . \
        --include-path ./node_modules \
        --abi "$contract" \
        --bin "$contract" \
        --output-dir "$workdir"

    test -d "$pkgdir" || mkdir -p "$pkgdir"

    contract=$(echo $contract | tr /. _)

    abigen \
        --abi $workdir/${contract}_*.abi \
        --bin $workdir/${contract}_*.bin \
        --pkg "$pkg" \
        --out "${pkgdir}/${pkg}.go"

    ls -l "${pkgdir}/${pkg}.go"
done
