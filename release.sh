#!/bin/bash
# This script updates all the versions to do a release:
# ./release.sh 0.0.13 0.1.6

set -e

echo "previous version: $1";
PREV_VERSION=$1

echo "new version: $2";
VERSION=$2

if [[ -z "${PREV_VERSION}" || -z "${VERSION}" ]]
then
	echo "Must pass the previous and next version as parameters"
	exit 1
fi

make manifests

make bundle IMG="ghcr.io/mondoohq/mondoo-operator:v${VERSION}" VERSION="${VERSION}"
yq -i ".spec.replaces = \"mondoo-operator.v${PREV_VERSION}\"" ./bundle/manifests/mondoo-operator.clusterserviceversion.yaml 

yq -i ".appVersion = \"${VERSION}\"" charts/mondoo-operator/Chart.yaml
yq -i ".version = \"${VERSION}\"" charts/mondoo-operator/Chart.yaml
CHART_NAME=charts/mondoo-operator make helm
