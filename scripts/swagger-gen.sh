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

__dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
__file="${__dir}/../vendor/github.com/intelsdi-x/snap/swagger.json"

# shellcheck source=scripts/common.sh
. "${__dir}/common.sh"

if [ -f "${__file}" ]
then
    rm -rf ${__dir}/../client
    rm -rf ${__dir}/../models
    swagger generate client -f ${__file} -A snap -a snap
else
    msg="${__file} not found.
        please run the following command to get Snap in your GOPATH.
        go get github.com/intelsdi-x/snap"

   _error $msg
fi

