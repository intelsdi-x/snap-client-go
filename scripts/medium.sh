#!/usr/bin/env bash

#http://www.apache.org/licenses/LICENSE-2.0.txt
#
#
#Copyright 2017 Intel Corporation
#
#Licensed under the Apache License, Version 2.0 (the "License");
#you may not use this file except in compliance with the License.
#You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
#Unless required by applicable law or agreed to in writing, software
#distributed under the License is distributed on an "AS IS" BASIS,
#WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#See the License for the specific language governing permissions and
#limitations under the License.

set -e
set -u
set -o pipefail


set -e
set -u
set -o pipefail

EXIT_ON_ERROR=0
command -v docker >/dev/null 2>&1 || { echo >&2 "Error: docker needs to be installed."; EXIT_ON_ERROR=1; }
command -v docker-compose >/dev/null 2>&1 || { echo >&2 "Error: docker-compose needs to be installed."; EXIT_ON_ERROR=1; }
docker version >/dev/null 2>&1 || { echo >&2 "Error: docker needs to be configured."; EXIT_ON_ERROR=1; }
if [[ $EXIT_ON_ERROR > 0 ]]; then
    exit 1
fi

docker-compose -f test/docker-compose.yml up --build -d

DOCKER_HOST=${DOCKER_HOST-}
if [[ -z "${DOCKER_HOST}" ]]; then
  SNAP_CLIENT_GO_HOST="127.0.0.1"
else
  SNAP_CLIENT_GO_HOST=`echo $DOCKER_HOST | grep -o '[0-9]\+[.][0-9]\+[.][0-9]\+[.][0-9]\+'`
fi

export SNAP_CLIENT_GO_HOST
echo "snap-client-go Host: ${SNAP_CLIENT_GO_HOST}"

echo "Waiting for snap client go test docker container"
while ! curl -sG "http://${SNAP_CLIENT_GO_HOST}:8181/v2/plugins" > /dev/null 2>&1; do
  sleep 1
  echo -n "."
done
echo

# for plugin load
(cd /tmp && curl -sfLSO https://s3-us-west-2.amazonaws.com/snap.ci.snap-telemetry.io/snap/1.1.0/linux/x86_64/snap-plugin-collector-mock1 && chmod 755 snap-plugin-collector-mock1)

UNIT_TEST="go_test"
test_unit

echo "Cleanup container"
docker-compose -f test/docker-compose.yml down
