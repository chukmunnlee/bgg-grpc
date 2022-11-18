#!/bin/bash
IP_ADDR="$(minikube ip)"

openssl genrsa -out server.key 4096

openssl req -new -key server.key -out server.csr \
	-subj "/CN=system:node:bgg-grpc/O=system:nodes" \
	-addext "subjectAltName = DNS:bgg-grpc-${IP_ADDR}.nip.io, DNS:localhost, IP:127.0.0.1,IP:${IP_ADDR}" \
	-addext 'keyUsage = keyEncipherment, digitalSignature' \
	-addext 'extendedKeyUsage = serverAuth' 

