#!/bin/bash

set -e
##Test creation of cluster and container
CONTAINER_NAME=`cat /dev/urandom | env LC_CTYPE=C tr -cd 'a-f' | head -c 5`
CLUSTER_NAME=`cat /dev/urandom | env LC_CTYPE=C tr -cd 'a-f' | head -c 5`

./workdir/amazonian --VPC=vpc-c7aa77be --HostedZoneName=vssdevelopment.com \
--Image=willejs/go-hello-world --ServiceName=${CONTAINER_NAME} --ContainerName=${CONTAINER_NAME} \
--ClusterName=${CLUSTER_NAME} --ClusterExists=false --Subnets=subnet-b61d81fe,subnet-0202dc58 --KeyName=dummy_key1 \
ClusterSize=1 mazSizePrt=1 instanceTypePrt=t2.medium

curl --fail https://${CONTAINER_NAME}.vssdevelopment.com/

aws cloudformation delete-stack --stack-name "${CONTAINER_NAME}"

##Now test and ensure it can reuse the same cluster
CONTAINER_NAME=`cat /dev/urandom | env LC_CTYPE=C tr -cd 'a-f' | head -c 5`

./workdir/amazonian --VPC=vpc-c7aa77be --HostedZoneName=vssdevelopment.com \
--Image=willejs/go-hello-world --ServiceName=${CONTAINER_NAME} --ContainerName=${CONTAINER_NAME} \
--ClusterName=${CLUSTER_NAME} --ClusterExists

curl --fail https://${CONTAINER_NAME}.vssdevelopment.com/

aws cloudformation delete-stack --stack-name "${CONTAINER_NAME}"
aws cloudformation delete-stack --stack-name "${CLUSTER_NAME}"
