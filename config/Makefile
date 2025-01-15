generate_crt_key:

	openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout private.key -out certificate.crt -config openssl.conf -extensions v3_req

describe_certificate_into_cmd:

	openssl x509 -in certificate.crt -text -noout