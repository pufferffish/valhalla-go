{ nixpkgs ? import <nixpkgs> {}
, lib
, stdenv
, abseil-cpp
, cmake
, fetchFromGitHub
, fetchpatch
, gtest
, zlib

# downstream dependencies
, python3

, ...
}:

with nixpkgs;

let
  valhallaCustom = (import ./valhalla) { inherit stdenv fetchFromGitHub cmake; };
in stdenv.mkDerivation rec {
  name = "valhalla-go";
  src = ./.;

  buildInputs = [
    boost179
    valhallaCustom
    zlib.static
    protobuf
  ];

  buildPhase = ''
    c++ \
      valhalla_go.cpp \
      -fPIC \
      -shared \
      -o libvalhalla_go.so \
      -Wl,-Bstatic \
      -lvalhalla \
      -lprotobuf-lite \
      -lz \
      -Wl,-Bdynamic \
      -lpthread
  '';

  installPhase = ''
    mkdir -p $out/lib
    cp libvalhalla_go.so $out/lib
  '';
}
