openssl req -newkey rsa:2048 \
            -x509 \
            -sha256 \
            -days 3650 \
            -nodes \
            -out server.crt \
            -keyout server.key \
            -config sample-admission.conf