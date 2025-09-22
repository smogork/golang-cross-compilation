# golang cross compilation + bazel

To build the project to linux platform just run `bazel build --platforms=@rules_go//go/toolchain:linux_amd64 //:hello`.
Change it to `bazel build --platforms=@rules_go//go/toolchain:linux_amd64_cgo //:hello` to cross compile also cgo code.