# https://medium.com/@lizrice/non-privileged-containers-based-on-the-scratch-image-a80105d6d341
FROM ubuntu:latest AS userland
RUN useradd --uid 10001 scratchuser
# FROM alpine:latest as certs
# RUN apk --update add ca-certificates
FROM scratch
# ENV PATH=/bin
# COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=userland /etc/passwd /etc/passwd
COPY --chown=scratchuser occi /usr/local/bin/occi
USER scratchuser
ENTRYPOINT [ "/usr/local/bin/occi" ]