FROM alpine
ADD bin/main /home/
COPY conf.yml /home/conf/conf.yml
CMD ["/home/main --config.file=/home/conf/conf.yml"]
