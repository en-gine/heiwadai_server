FROM envoyproxy/envoy:v1.20.0

ADD ./docker/envoy/envoy.yaml /etc/envoy/

CMD /usr/local/bin/envoy -c /etc/envoy/envoy.yaml