{ nixpkgs ? import <nixpkgs> {}, stdenv, fetchFromGitHub, cmake }:

with nixpkgs;

stdenv.mkDerivation rec {
  name = "valhalla";

  src = fetchFromGitHub {
    owner = "valhalla";
    repo = "valhalla";
    rev = "3.4.0";
    sha256 = "honnvgmT1u26vv2AdtLfHou7B640PXaV3s0XXNkd/QE=";
    fetchSubmodules = true;
  };

  cmakeFlags = [
    "-DENABLE_CCACHE=OFF"
    "-DBUILD_SHARED_LIBS=OFF"
    "-DENABLE_BENCHMARKS=OFF"
    "-DENABLE_PYTHON_BINDINGS=OFF"
    "-DENABLE_TESTS=OFF"
    "-DENABLE_TOOLS=OFF"
    "-DENABLE_SERVICES=OFF"
    "-DENABLE_HTTP=OFF"
    "-DENABLE_CCACHE=OFF"
    "-DENABLE_DATA_TOOLS=OFF"
    "-DCMAKE_BUILD_TYPE=Release"
  ];

  buildInputs = [
    cmake
    zlib
    boost179
    protobuf
    sqlite
    libspatialite
    luajit
    geos
  ];

  # install necessary headers
  postInstall = ''
    cp -r $src/third_party/rapidjson/include/* $out/include
    cp -r $src/third_party/date/include/* $out/include
  '';

}
