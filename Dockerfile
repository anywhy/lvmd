FROM alpine:latest

RUN apk update && apk add udev blkid file util-linux e2fsprogs lvm2 udev sgdisk device-mapper

ADD ./build/bin/lvmd-server /usr/local/bin/lvmd-server