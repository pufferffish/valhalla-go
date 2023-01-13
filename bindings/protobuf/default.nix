{ lib
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

let
  version = "3.21.8";
  self = stdenv.mkDerivation {
    name = "protobuf";
  
    src = fetchFromGitHub {
      owner = "protocolbuffers";
      repo = "protobuf";
      rev = "v${version}";
      sha256 = "sha256-cSNHX18CvMmydpYWqfe6WWk9rGxIlFfY/85rfSyznU4=";
    };
  
    # re-create submodule logic
    postPatch = ''
      rm -rf gmock
      cp -r ${gtest.src}/googlemock third_party/gmock
      cp -r ${gtest.src}/googletest third_party/
      chmod -R a+w third_party/
      ln -s ../googletest third_party/gmock/gtest
      ln -s ../gmock third_party/googletest/googlemock
      ln -s $(pwd)/third_party/googletest third_party/googletest/googletest
    '' + lib.optionalString stdenv.isDarwin ''
      substituteInPlace src/google/protobuf/testing/googletest.cc \
        --replace 'tmpnam(b)' '"'$TMPDIR'/foo"'
    '';
  
    patches = [
      # fix protobuf-targets.cmake installation paths, and allow for CMAKE_INSTALL_LIBDIR to be absolute
      # https://github.com/protocolbuffers/protobuf/pull/10090
      (fetchpatch {
        url = "https://github.com/protocolbuffers/protobuf/commit/a7324f88e92bc16b57f3683403b6c993bf68070b.patch";
        sha256 = "sha256-SmwaUjOjjZulg/wgNmR/F5b8rhYA2wkKAjHIOxjcQdQ=";
      })
      ./static-executables-have-no-rpath.patch
    ];
  
    nativeBuildInputs = let
      protobufVersion = "${version}";
    in [
      cmake
    ];
  
    buildInputs = [
      abseil-cpp
      zlib
    ];
  
    # After 3.20, CMakeLists.txt can now be found at the top-level, however
    # a stub cmake/CMakeLists.txt still exists for compatibility with previous build assumptions
    cmakeDir = "../cmake";
    cmakeFlags = [
      "-Dprotobuf_ABSL_PROVIDER=package"
    ];
  
    # unfortunately the shared libraries have yet to been patched by nix, thus tests will fail
    doCheck = false;
  
    passthru = {
      tests = {
        pythonProtobuf = python3.pkgs.protobuf.override(_: {
          protobuf = self;
        });
      };
    };
  
    meta = {
      description = "Google's data interchange format";
      longDescription = ''
        Protocol Buffers are a way of encoding structured data in an efficient
        yet extensible format. Google uses Protocol Buffers for almost all of
        its internal RPC protocols and file formats.
      '';
      license = lib.licenses.bsd3;
      platforms = lib.platforms.unix;
      homepage = "https://developers.google.com/protocol-buffers/";
      maintainers = with lib.maintainers; [ jonringer ];
      mainProgram = "protoc";
    };
  };
in
  self
