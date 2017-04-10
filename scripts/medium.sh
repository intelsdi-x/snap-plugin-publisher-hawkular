#!/bin/bash

set -e
set -u
set -o pipefail

cassId=$(docker run -d --name snapcass -e CASSANDRA_START_RPC=true cassandra:3.0.9)
hawkId=$(docker run -d --name snaphawk --link=snapcass -e CASSANDRA_NODES=snapcass -p 8080:8080 -e ADMIN_TOKEN=topsecret hawkular/hawkular-services:latest)
DOCKER_HOST=${DOCKER_HOST-}
if [[ -z "${DOCKER_HOST}" ]]; then
  host="127.0.0.1"
else
  host=`echo $DOCKER_HOST | grep -o '[0-9]\+[.][0-9]\+[.][0-9]\+[.][0-9]\+'`
fi

export SNAP_HAWKULAR_HOST="${host}"
_info "Hawkular Host: ${SNAP_HAWKULAR_HOST}"


_info "Waiting for hawkular docker container"
while ! curl -u jdoe:password -X GET http://${SNAP_HAWKULAR_HOST}:8080/hawkular/metrics/metrics -H "Content-Type: application/json" -H "Hawkular-Tenant: myTenant" > /dev/null 2>&1; do
  sleep 1
  echo -n "."
done
echo

secs=$((2 * 60))
while [ $secs -gt 0 ]; do
   echo -ne "hawkular is preparing cassandra; second left: $secs\033[0K\r"
   sleep 1
   : $((secs--))
done
echo

UNIT_TEST="go_test"
test_unit

_debug "Cleanup container ${cassId} and ${hawkId}"
docker rm -f "${hawkId}"
docker rm -f "${cassId}"
