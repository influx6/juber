# Juber
A github project template for creating quick startup for projects built 
and deployed to Kubernetes.

## Platform Requirements

- Docker Compose
- Docker (Install [Docker for Windows](https://docs.docker.com/docker-for-windows/) or [Docker Toolbox](https://docs.docker.com/toolbox/toolbox_install_windows/), [Vagrant](https://www.vagrantup.com/) and [Virutalbox](https://www.virtualbox.org/) if you have no hyper-v support)
- Make (Gnu Make for windows, Install via [Chocolatey](https://chocolatey.org/packages/make) or [SourceForge](https://sourceforge.net/projects/gnuwin32/))
- Homebrew (Mac-Only, install [Scoop for Windows](https://scoop.sh/) or [GoFish for Linux/Windows](https://gofi.sh/) or [LinuxBrew for Linux](https://docs.brew.sh/Homebrew-on-Linux))
- Kubernetes Options:
    - MicroK8S (Preferably, Installing [Microk8s](https://microk8s.io/docs/install-alternatives))
    - Kubernetes in Docker ([Kubernetes in Docker](https://www.docker.com/products/kubernetes))
- Tilt (Used [Tilt](https://tilt.dev/) for development lifecycle automation)

## Spin up dev environment for project on K8S

```bash
make devup
```