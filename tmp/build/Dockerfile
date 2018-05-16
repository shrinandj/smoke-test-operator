FROM alpine:3.6

ADD tmp/_output/bin/smoke-test-operator /usr/local/bin/smoke-test-operator

RUN adduser -D smoke-test-operator
USER smoke-test-operator
