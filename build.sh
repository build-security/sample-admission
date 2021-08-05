#!/bin/sh
# Generate cert & CA bundle
function cert_gen () {
  openssl req \
            -new \
            -x509 \
            -days 3650 \
            -nodes \
            -subj '/CN=my-ca' \
            -out cert/ca.crt \
            -keyout cert/ca.key \
  openssl genrsa \
            -out cert/server.key 2048
  openssl req \
            -new \
            -key cert/server.key \
            -subj '/CN=sample-admission.kube-system.svc' \
            -out cert/server.csr
  openssl x509 \
  -req \
  -in cert/server.csr \
  -CA cert/ca.crt \
  -CAkey cert/ca.key \
  -CAcreateserial \
  -days 365 \
  -out cert/server.crt
}

#base64 encode and sed to k8s.yaml file on validationwebhookconfiguration(maybe have tpl & mv replace on build)
function ca_encode () {
  unset CA_BUNDLE_ENV
  export CA_BUNDLE_ENV=$(cat cert/ca.crt | base64) 
  sed -i '' 's/CAplaceholder/'"$CA_BUNDLE_ENV"'/g' validating-webhook.yaml
}

#Run dockerbuild cmd with the proper tags
function dockerbuild_eval () {
  docker build -t buildsecurity/sample-admission .
  kubectl apply -f sample-admission.yaml --dry-run=server
  kubectl apply -f sample-admission.yaml
  kubectl apply -f validating-webhook.yaml
}


main() {
cert_gen
ca_encode
dockerbuild_eval
}

main "$@"; exit


