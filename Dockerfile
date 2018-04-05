FROM alpine
ADD bin/main /home
ADD conf.yml /home
CMD ["/home/main"]
