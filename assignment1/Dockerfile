FROM ubuntu:16.04

WORKDIR /some_random/dummy_dir
ADD . .

RUN apt-get update && apt-get install sudo -y
CMD ["sh", "./initializer.sh" ]