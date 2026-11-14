# Mach-O ARM64 TLS relocation repro

On macOS arm64 with stock Go 1.27:

```sh
CGO_ENABLED=0 CC=/usr/bin/false go build .
```

Observed:

```text
main(__TEXT/__text): unexpected relocation type 2066 (ARM64_RELOC_TLVP_LOAD_PAGEOFF12)
main(__TEXT/__text): unexpected relocation type 2065 (ARM64_RELOC_TLVP_LOAD_PAGE21)
```

The desired result is a successful build; the program calls `tls_next` twice and checks that the values are 7 and 8. The checked-in `tls_darwin_arm64.syso` is a 936-byte Mach-O object produced from [`producer/tls.c`](producer/tls.c). Its only unresolved symbol is `_tlv_bootstrap`, which [`imports_darwin_arm64.go`](imports_darwin_arm64.go) explicitly imports so the failure isolates TLVP relocation handling. The Go assembly wrapper calls the C ABI function without using cgo.

To regenerate the object on macOS arm64 (producer only):

```sh
clang -target arm64-apple-macos11 -std=c11 -O2 -c producer/tls.c -o tls_darwin_arm64.syso
```

The resulting Go package has no C source files. It builds and runs with our patched linker and `CGO_ENABLED=0`.
