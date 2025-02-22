# REX-Ray 
[![Chat](https://img.shields.io/static/v1?label=Chat&message=Support&color=brightgreen)](https://matrix.to/#/#mesos:matrix.aventer.biz?via=matrix.aventer.biz)
[![GoDoc](https://godoc.org/github.com/AVENTER-UG/rexray?status.svg)](http://godoc.org/github.com/AVENTER-UG/rexray) 

---

![info](https://cdn.rawgit.com/akutz/741a53ec8cd1348753556e8bd4d2836a/raw/399cb9e5b39436d119d77a893dd991db0a7b6f9f/info-circle.svg "info-circle") **Note:* This is a unofficial fork of the unmaintained rexray. 
But we will care issues and PR's so good as possible.

---


The long-term goal of the REX-Ray project is to enable collaboration between
organizations  focused on creating enterprise-grade storage plugins for the
Container Storage Interface (CSI). As a rapidly changing specification, CSI
support within REX-Ray will be planned when CSI reaches version 1.0, currently
projected for a late 2018 release. In the interim, there remains active
engagement with the project to support the community.

REX-Ray provides a vendor agnostic storage orchestration engine.  The primary
design goal is to provide persistent storage for `Docker`, `Kubernetes`, and `Mesos`.

It is additionally available as a Go package, CLI tool, and Linux service which
enables it to be used for additional use cases.

## Documentation [![Docs](https://aventer-ug.github.io/rexray/)](https://aventer-ug.github.io/rexray/)
Documentation is based on the latest stable build.  The `/.docs`
directory in this repo will refer to the latest or specific commit.

## Architecture
REX-Ray is available as a standalone process today and as a distributed
model of client-server.  The `client` performs a level abstraction of local
host processes (request for volume attachment, discovery, format, and mounting
of devices) while the `server` provides the necessary abstraction of the
control plane for multiple storage platforms/

### Storage Provider Support
The following storage providers and platforms are supported by REX-Ray.

| Provider              | Storage Platform  | <center>[Docker](https://docs.docker.com/engine/extend/plugins_volume/)</center> | <center>Containerized</center> |
|-----------------------|----------------------|:---:|:---:|
| Amazon EC2 | [EBS](.docs/user-guide/storage-providers/aws.md#aws-ebs) | ✓ | ✓ |
| | [EFS](.docs/user-guide/storage-providers/aws.md#aws-efs) | ✓ | ✓ |
| | [S3FS](.docs/user-guide/storage-providers/aws.md#aws-s3fs) | ✓ | ✓ |
| Ceph | [RBD](.docs/user-guide/storage-providers/ceph.md#ceph-rbd) | ✓ | ✓ |
| Dell EMC | [Isilon](.docs/user-guide/storage-providers/dellemc.md#dell-emc-isilon) | ✓ | ✓ |
| | [ScaleIO](.docs/user-guide/storage-providers/dellemc.md#dell-emc-scaleio) | ✓ | ✓ |
| DigitalOcean | [Block Storage](.docs/user-guide/storage-providers/digitalocean.md#do-block-storage) | ✓ | ✓ |
| FittedCloud | [EBS Optimizer](.docs/user-guide/storage-providers/fittedcloud.md#ebs-optimizer) | ✓ | |
| Google | [GCE Persistent Disk](.docs/user-guide/storage-providers/google.md#gce-persistent-disk) | ✓ | ✓ |
| Microsoft | [Azure Unmanaged Disk](.docs/user-guide/storage-providers/microsoft.md#azure-ud) | ✓ | ✓ |
| OpenStack | [Cinder](.docs/user-guide/storage-providers/openstack.md#cinder) | ✓ | ✓ |

### Operating System Support
The following operating systems are supported by REX-Ray:

| OS             | <center>Command Line</center> | <center>Service</center> |
|---------------|:---:|:---:|
| Ubuntu 12+     | ✓          | ✓ |
| Debian 6+      | ✓          | ✓ |
| RedHat         | ✓          | ✓ |
| CentOS 6+      | ✓          | ✓ |
| CoreOS         | ✓          | ✓ |
| TinyLinux (boot2docker)| ✓  | ✓ |
| OS X Yosemite+ | ✓          |  |
| Windows        |            |  |

## Installation

We only deliver prebuild version for CentOS 7. You can add our rpm repository: 

```bash
[aventer-rel]
name=AVENTER stable repository $releasever
baseurl=http://rpm.aventer.biz/CentOS/$releasever/$basearch/
enabled=1
gpgkey=https://www.aventer.biz/CentOS/support_aventer.asc
```

## Runtime - CLI
REX-Ray can be run as an interactive CLI to perform volume management
capabilities.

```bash
$ export REXRAY_SERVICE=ebs
$ export EBS_ACCESSKEY=access_key
$ export EBS_SECRETKEY=secret_key
$ rexray volume ls
ID            Name  Status    Size
vol-6ac6c7d6        attached  8
```

## Runtime - Service (Docker)
Additionally, it can be run as a service to support `Docker`, `Mesos`, and other
 platforms that can communicate through `HTTP/JSON`.

```bash
$ export REXRAY_SERVICE=ebs
$ export EBS_ACCESSKEY=access_key
$ export EBS_SECRETKEY=secret_key
$ rexray service start
Starting REX-Ray...SUCCESS!

  The REX-Ray daemon is now running at PID XX. To
  shutdown the daemon execute the following command:

    sudo /usr/bin/rexray stop

$ docker run -ti --volume-driver=rexray -v test:/test busybox
$ df -h /test
```

## Runtime - Docker Plugin
Starting with Docker 1.13, Docker now supports a new plugin architecture in
which plugins can be installed as containers.

```bash
$ docker plugin install rexray/ebs EBS_ACCESSKEY=access_key EBS_SECRETKEY=secret_key
Plugin "rexray/ebs:latest" is requesting the following privileges:
 - network: [host]
 - mount: [/dev]
 - allow-all-devices: [true]
 - capabilities: [CAP_SYS_ADMIN]
Do you grant the above permissions? [y/N] y
latest: Pulling from rexray/ebs
2ef3a0b3d192: Download complete
Digest: sha256:86a3bf7fdab857c955d7ef3fb94c01e350e34ba0f7fd3d0bd485e45f1592e1c2
Status: Downloaded newer image for rexray/ebs:latest
Installed plugin rexray/ebs:latest

$ docker plugin ls
ID                  NAME                   DESCRIPTION              ENABLED
450420731dc3        rexray/ebs:latest      REX-Ray for Amazon EBS   true

$ docker run -ti --volume-driver=rexray/ebs -v test:/test busybox
$ df -h /test
```
