# http://www.apache.org/licenses/LICENSE-2.0.txt
#
#
# Copyright 2016 Intel Corporation
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

FROM intelsdi/snap:alpine_test

ENV FILE_PUB_PLUGIN_VER=2

ADD http://snap.ci.snap-telemetry.io/plugins/snap-plugin-publisher-file/$FILE_PUB_PLUGIN_VER/linux/x86_64/snap-plugin-publisher-file /opt/snap/plugins/snap-plugin-publisher-file
ADD https://s3-us-west-2.amazonaws.com/snap.ci.snap-telemetry.io/snap/$SNAP_VERSION/linux/x86_64/snap-plugin-collector-mock1 /opt/snap/plugins/snap-plugin-collector-mock1
ADD https://s3-us-west-2.amazonaws.com/snap.ci.snap-telemetry.io/snap/$SNAP_VERSION/linux/x86_64/snap-plugin-collector-mock2 /opt/snap/plugins/snap-plugin-collector-mock2
