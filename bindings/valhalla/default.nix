{ nixpkgs ? import <nixpkgs> {}, stdenv, fetchgit, cmake }:

with nixpkgs;

stdenv.mkDerivation rec {

  name = "valhalla";

  src = fetchgit {
    url = "https://github.com/valhalla/valhalla.git";
    rev = "3.3.0";
    sha256 = "honnvgmT1u26vv2AdtLfHou7B640PXaV3s0XXNkd/QE=";
  };

  buildInputs = [
    cmake
    zlib
    boost172
    curl
    protobuf
    sqlite
    libspatialite
    luajit
    python310
  ];

  builder = ./builder.sh;
}