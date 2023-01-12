source "$stdenv"/setup

cp --recursive "$src" ./

chmod --recursive u=rwx ./"$(basename "$src")"

cd ./"$(basename "$src")"

cmake -B build \
  -DENABLE_CCACHE=OFF \
  -DBUILD_SHARED_LIBS=OFF \
  -DENABLE_BENCHMARKS=OFF \
  -DENABLE_PYTHON_BINDINGS=ON \
  -DENABLE_TESTS=OFF \
  -DENABLE_TOOLS=OFF \
  -DENABLE_SERVICES=OFF \
  -DENABLE_HTTP=OFF \
  -DENABLE_CCACHE=OFF \
  -DCMAKE_BUILD_TYPE=Release
cmake --build build -- -j$(nproc)