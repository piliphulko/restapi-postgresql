#!/bin/sh
set -e

cat <<EOF > /vault/config/custom_config.hcl
storage "postgresql" {
  connection_url = "${VAULT_POSTGRES_URL}"
}

listener "tcp" {
  address         = "0.0.0.0:8200"
  tls_cert_file   = "/usr/local/bin/certificate.crt"
  tls_key_file    = "/usr/local/bin/private.key"
}

api_addr     = "https://host.docker.internal:8200"
cluster_addr = "https://host.docker.internal:8201"
ui = true
EOF

vault server -config=/vault/config/custom_config.hcl &


sleep 3
rm -f /vault/config/custom_config.hcl

wait