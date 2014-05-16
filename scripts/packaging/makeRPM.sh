#!/bin/bash

### http://tecadmin.net/create-rpm-of-your-own-script-in-centosredhat/#

sudo yum install rpm-build rpmdevtools
rm -rf ~/rpmbuild
rpmdev-setuptree

mkdir ~/rpmbuild/SOURCES/hydra-worker-map-by-limit-1
cp hydra-worker-map-by-limit-init.d.sh ~/rpmbuild/SOURCES/hydra-worker-map-by-limit-1
cp ../../bin/hydra-worker-map-by-limit ~/rpmbuild/SOURCES/hydra-worker-map-by-limit-1

cp hydra-worker-map-by-limit.spec ~/rpmbuild/SPECS

pushd ~/rpmbuild/SOURCES/
tar czf hydra-worker-map-by-limit-1.0.tar.gz hydra-worker-map-by-limit-1/
cd ~/rpmbuild 
rpmbuild -ba SPECS/hydra-worker-map-by-limit.spec

popd
cp ~/rpmbuild/RPMS/x86_64/hydra-worker-map-by-limit-1-0.x86_64.rpm .
