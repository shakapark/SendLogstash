FROM scratch
RUN mkdir /home
ADD bin/main /home
ADD conf.yml /home
CMD ["/home/main"]
