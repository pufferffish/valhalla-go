{ nixpkgs ? import <nixpkgs> {}, stdenv, fetchFromGitHub, cmake }:

with nixpkgs;

let
  valhallaCustom = (import ./valhalla) { inherit stdenv fetchFromGitHub cmake; };
in stdenv.mkDerivation rec {
  name = "valhalla-go";
  src = ./.;

  buildInputs = [
    boost172
    valhallaCustom
    zlib.static
    protobuf
    curl
  ];

  buildPhase = ''
    c++ \
      valhalla_go.cpp \
      -lvalhalla \
      -lprotobuf \
      -lcurl \
      -lz \
      -lpthread \
      -shared \
      -fPIC \
      -o libvalhalla_go.so
  '';

  installPhase = ''
    mkdir -p $out/lib
    cp libvalhalla_go.so $out/lib
  '';
}
