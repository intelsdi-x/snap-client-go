FROM intelsdi/snap:alpine_test

ENV FILE_PUB_PLUGIN_VER=2

ADD http://snap.ci.snap-telemetry.io/plugins/snap-plugin-publisher-file/$FILE_PUB_PLUGIN_VER/linux/x86_64/snap-plugin-publisher-file /opt/snap/plugins/snap-plugin-publisher-file
ADD https://s3-us-west-2.amazonaws.com/snap.ci.snap-telemetry.io/snap/$SNAP_VERSION/linux/x86_64/snap-plugin-collector-mock1 /opt/snap/plugins/snap-plugin-collector-mock1
ADD https://s3-us-west-2.amazonaws.com/snap.ci.snap-telemetry.io/snap/$SNAP_VERSION/linux/x86_64/snap-plugin-collector-mock2 /opt/snap/plugins/snap-plugin-collector-mock2
