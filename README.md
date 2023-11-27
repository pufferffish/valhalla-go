# Valhalla for Go

[![Test and build bindings](https://github.com/vandreltd/valhalla-go/actions/workflows/build.yml/badge.svg)](https://github.com/vandreltd/valhalla-go/actions/workflows/build.yml)

This repo simply offers Go bindings to the [Valhalla project](https://github.com/valhalla/valhalla).

## Usage

The library offer functions that directly take JSON string request and return JSON string response.
Example code on how to use the library can be found in the [test units](/valhalla_test.go).

Note that the library depends on C++ bindings. If you have the [Nix package manager](https://nixos.org/) you can simply build the bindings as such:
```
git clone https://github.com/vandreltd/valhalla-go
cd valhalla-go
nix-build # the shared library will be in result/lib/
LD_LIBRARY_PATH=./result/lib go test -v # build and run the test units
```

If you do not wish the build the library yourself, you can grab a pre-built binary in the [CI Artifacts](https://github.com/vandreltd/valhalla-go/actions).

A preprocessed tiles file of the entire world (2023 Jan) can be found [here](https://archive.org/download/valhalla-planet-221219).

## License

`valhalla-go` is licensed with ISC, see [LICENSE](./LICENSE).
