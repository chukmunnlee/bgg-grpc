#!/bin/bash
kubectl create secret tls bgg-grpc-tls \
	--key server.key --cert server.crt \
	-oyaml --dry-run=client | tee bgg-grpc-tls.yaml
