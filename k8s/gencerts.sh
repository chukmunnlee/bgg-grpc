#!/bin/bash
IP_ADDR="$(minikube ip)"
openssl genrsa -out server.key 4096
openssl req -x509 -key server.key -out server.crt \
	-subj '/CN=bgg-grpc' \
	-addext "subjectAltName = DNS:bgg-grpc-${IP_ADDR}.nip.io, DNS:localhost, IP:127.0.0.1, IP:${IP_ADDR}" \
	-addext 'keyUsage = keyEncipherment, digitalSignature' \
	-addext 'extendedKeyUsage = serverAuth' \
	-sha256 -days 365 

kubectl create secret tls bgg-grpc-tls --key=server.key --cert=server.crt -oyaml --dry-run=client | tee bgg-grpc-tls.yaml
