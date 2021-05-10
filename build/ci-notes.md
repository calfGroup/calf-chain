You need to run Ubuntu to do test packaging.

Install any version of Go and Debian packaging tools:

    $ sudo apt-get install build-essential golang-go devscripts debhelper python-bzrlib python-paramiko

Create the source packages:

    $ go run build/ci.go debsrc -workdir dist

Then go into the source package directory for your running distribution and build the package:

    $ cd dist/ethereum-unstable-1.9.6+bionic
    $ dpkg-buildpackage

Built packages are placed in the dist/ directory.

    $ cd ..
    $ dpkg-deb -c geth-unstable_1.9.6+bionic_amd64.deb
