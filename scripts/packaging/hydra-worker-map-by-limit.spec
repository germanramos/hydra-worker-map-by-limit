Name: hydra-worker-map-by-limit
Version: 1
Release: 0
Summary: hydra-worker-map-by-limit
Source0: hydra-worker-map-by-limit-1.0.tar.gz
License: MIT
Group: custom
URL: https://github.com/innotech/hydra-worker-map-by-limit
BuildArch: x86_64
BuildRoot: %{_tmppath}/%{name}-buildroot
Requires: libzmq3
%description
Map instances by limit number and sort them.
%prep
%setup -q
%build
%install
install -m 0755 -d $RPM_BUILD_ROOT/usr/local/hydra
install -m 0755 hydra-worker-map-by-limit $RPM_BUILD_ROOT/usr/local/hydra/hydra-worker-map-by-limit

install -m 0755 -d $RPM_BUILD_ROOT/etc/init.d
install -m 0755 hydra-worker-map-by-limit-init.d.sh $RPM_BUILD_ROOT/etc/init.d/hydra-worker-map-by-limit

install -m 0755 -d $RPM_BUILD_ROOT/etc/hydra
install -m 0644 hydra.conf $RPM_BUILD_ROOT/etc/hydra/hydra-worker-map-by-limit.conf
%clean
rm -rf $RPM_BUILD_ROOT
%post
echo   You should edit config file /etc/hydra/hydra-worker-map-by-limit.conf
echo   When finished, you may want to run \"update-rc.d hydra-worker-map-by-limit defaults\"
%files
/usr/local/hydra/hydra-worker-map-by-limit
/etc/init.d/hydra-worker-map-by-limit
