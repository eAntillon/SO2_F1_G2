# Instrucciones de ejecución:
- `apt-get update`
- `apt-get install build-essential`
- `apt install gcc`

- `apt-get install libseccomp-dev pkg-config`


_____________________________________________________________________________________________
//// se debe hacer esto si da problema la ultima instalacion  `apt-get install libseccomp-dev pkg-config`
                  sudo dpkg --configure -a
                  sudo apt-get clean && sudo apt-get autoclean
                  sudo apt-get update --fix-missing
                  sudo apt-get install -f
                  sudo apt-get dist-upgrade

      SE BORRA LO QUE TENGA ESTE ARCHIVO
      nano /etc/apt/sources.list


      SE PEGAN TODAS ESTAS URL
deb http://archive.ubuntu.com/ubuntu/ focal main restricted universe multiverse
deb-src http://archive.ubuntu.com/ubuntu/ focal main restricted universe multiverse

deb http://archive.ubuntu.com/ubuntu/ focal-updates main restricted universe multiverse
deb-src http://archive.ubuntu.com/ubuntu/ focal-updates main restricted universe multiverse

deb http://archive.ubuntu.com/ubuntu/ focal-security main restricted universe multiverse
deb-src http://archive.ubuntu.com/ubuntu/ focal-security main restricted universe multiverse

deb http://archive.ubuntu.com/ubuntu/ focal-backports main restricted universe multiverse
deb-src http://archive.ubuntu.com/ubuntu/ focal-backports main restricted universe multiverse

deb http://archive.canonical.com/ubuntu focal partner
deb-src http://archive.canonical.com/ubuntu focal partner



POR ULTIMO
apt-get update
apt-get upgrade     (este es el mas tardado)
_____________________________________________________________________________________________




- `go get github.com/seccomp/libseccomp-golang`
- `go mod init [nombre del modulo golang]`
  - Ejemplo `go mod init strace`
- `go mod tidy`
- `go build .`
- `./main`


"EJEMPLO PARA EJECUTAR "

  echo hello
  ls -l