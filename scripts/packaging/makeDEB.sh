#!/bin/bash

### http://linuxconfig.org/easy-way-to-create-a-debian-package-and-local-package-repository

rm -rf ~/debbuild
mkdir -p ~/debbuild/DEBIAN
cp control ~/debbuild/DEBIAN

mkdir -p ~/debbuild/etc/init.d
cp hydra-worker-map-by-limit-init.d.sh ~/debbuild/etc/init.d/hydra-worker-map-by-limit

mkdir -p ~/debbuild/usr/local/hydra
cp ../../bin/hydra-worker-map-by-limit  ~/debbuild/usr/local/hydra

chmod -R 644 ~/debbuild/usr/local/hydra/* ~/debbuild/etc/hydra/*
chmod 755 ~/debbuild/etc/init.d/hydra-worker-map-by-limit
chmod 755 ~/debbuild/usr/local/hydra/hydra-worker-map-by-limit

sudo chown -R root:root ~/debbuild/*

pushd ~
sudo dpkg-deb --build debbuild

popd
sudo mv ~/debbuild.deb hydra-worker-map-by-limit-1-0.x86_64.deb
