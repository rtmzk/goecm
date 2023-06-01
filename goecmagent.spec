Name: goecmagent
Version: VERSION
Release: 1%{?dist}
Summary: apm agent do some stuff from server

License: MIT
Url: https://macrowing.com
Packager:  group.devops
Source0: %{name}-%{version}.tar.gz

BuildRoot: %_topdir/BUILDROOT

BuildRequires: gcc,make
%description
goecmagent service
%prep:
%setup:

%build:
go build -o /usr/local/gopj/goecm/goecmagent cmd/goecmagent/goecmagent.go

%install:
mkdir -p ${RPM_BUILD_ROOT}/usr/local/bin/
cp -f /usr/local/gopj/goecm/goecmagent ${RPM_BUILD_ROOT}/usr/local/bin/goecmagent


mkdir -p ${RPM_BUILD_ROOT}/var/log/goecmagent
mkdir -p ${RPM_BUILD_ROOT}/etc/goecmagent
mkdir -p ${RPM_BUILD_ROOT}/etc/systemd/system
cp -f /usr/local/gopj/goecm/agent-config.yaml ${RPM_BUILD_ROOT}/etc/goecmagent/agent-config.yaml

cp -f /usr/local/gopj/goecm/goecmagent.service ${RPM_BUILD_ROOT}/etc/systemd/system/goecmagent.service


%post

%preun

%postun

%clean
rm -fr %{buildroot}

%files
%defattr (-,root,root,-)
/usr/local/bin/goecmagent
/etc/systemd/system/goecmagent.service
/etc/goecmagent/agent-config.yaml

%dir
/etc/goecmagent
/var/log/goecmagent
