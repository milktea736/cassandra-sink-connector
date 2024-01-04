#!/bin/bash

dev_pod=$(kubectl get pod | grep go-dev-client | awk '{print $1}')
kubectl cp . $dev_pod:/go/src