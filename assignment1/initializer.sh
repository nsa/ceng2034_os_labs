#!/bin/bash

WORKDIR="some_random/working_dir"
if [ -d "some_random" ]; then
    rm -rf some_random/ 
fi
mkdir -p $WORKDIR

cp -r template $WORKDIR
cp assignment1 $WORKDIR

chmod +x assignment1
chmod +x $WORKDIR/assignment1
sudo chroot $WORKDIR /assignment1