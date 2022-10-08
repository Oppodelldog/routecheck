# routecheck

This tool checks route entries for intersections.

It is for example helpful when working with windows + vpn + wsl + docker since with such a setup networks
are created dynamically and are partially not possible to configure properly.

In such a case the tool would print a warning and exit with a non-zero exit code.

## Usage

```bash
# linux
routecheck

#windows
routecheck.exe
```

## Build

If go is installed, you can download, build and install the binary with the following command:

```bash
go install github.com/Oppodelldog/routecheck
```

Alternatively there are scripts for linux and windows to build the binary:

```bash
./build.sh
```

```commandline
./build.bat
```