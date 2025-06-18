#!/bin/bash

if ! command -v openssl &> /dev/null; then
    echo "openssl command not found. Please install openssl."
    exit 1
fi

DOMAIN=`hostname -f | cut -d. -f2-`
export DOMAIN=$DOMAIN
echo $DOMAIN
HOSTNAME=$(hostname).$DOMAIN
mkdir -p work 
pushd work
openssl genrsa -out ca.key 2048
openssl req -new -x509 -days 3650 -key ca.key -out ca.crt -subj "/C=IN/ST=KAR/L=BLR/O=ARRM/OU=SRE/CN=$HOSTNAME"
openssl genrsa -out server.key 2048
openssl req -new -key server.key -out server.csr -subj "/C=US/ST=CAL/L=CAL/O=CloudBolt/OU=SRE/CN=$HOSTNAME"
cat > server.ext <<EOF
authorityKeyIdentifier=keyid,issuer
basicConstraints=CA:FALSE
keyUsage = digitalSignature, keyEncipherment
extendedKeyUsage = serverAuth, clientAuth
subjectAltName = @alt_names

[alt_names]
DNS.1 = $HOSTNAME
EOF
openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 365 -sha256 -extfile server.ext
cat server.key server.crt > server.pem
openssl verify -CAfile ca.crt server.crt

popd
