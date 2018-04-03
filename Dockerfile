FROM scratch
ADD bin/main /
ADD conf.yml /
CMD ["/main"]
