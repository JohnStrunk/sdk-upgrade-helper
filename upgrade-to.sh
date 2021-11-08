#! /bin/bash

set -e -o pipefail

scriptdir="$(dirname "$(realpath "$0")")"
cd "$scriptdir"

if [[ $# -ne 1 ]]; then
    echo "Usage: $0 <operator-sdk version>"
    echo "  Example: $0 v1.10.0"
    exit 1
fi

OPERATOR_SDK_VERSION="$1"
OS="$(go env GOOS)"
ARCH="$(go env GOARCH)"


########################################
# Download SDK
echo "Downloading $OPERATOR_SDK_VERSION ..."
curl -sL -o ./osdk https://github.com/operator-framework/operator-sdk/releases/download/${OPERATOR_SDK_VERSION}/operator-sdk_${OS}_${ARCH}
chmod a+x ./osdk
./osdk version


########################################
# Scaffold the project
rm -rf scaffolding && mkdir scaffolding && cd scaffolding

../osdk init \
    --skip-go-version-check \
    --domain backube \
    --owner "The VolSync authors" \
    --project-name volsync \
    --license none \
    --repo github.com/backube/volsync

../osdk create api \
    --controller \
    --resource \
    --group volsync \
    --kind ReplicationSource \
    --version v1alpha1

../osdk create api \
    --controller \
    --resource \
    --group volsync \
    --kind ReplicationDestination \
    --version v1alpha1

cd "$scriptdir"


########################################
# Commit changes
rm -f osdk

echo "Committing changes..."
git add .
git commit -s -m "Upgrade to ${OPERATOR_SDK_VERSION}"
