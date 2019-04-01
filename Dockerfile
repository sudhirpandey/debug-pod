FROM scratch
MAINTAINER Sudhir Pandey <sudpande@redhat.com>
COPY debug-pod /debug-pod
EXPOSE 8080 8888
USER 1001
ENTRYPOINT ["/debug-pod"]
