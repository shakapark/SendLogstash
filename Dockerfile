FROM alpine
RUN mkdir /home/conf
ADD bin/main /home/
COPY conf.yml /home/conf/
CMD ["/home/main","--config.file=/home/conf/conf.yml"]
