# Local REALITY compatibility patch

Source: https://github.com/XTLS/REALITY/tree/83f621788ca8

Go module version: `v0.0.0-20261014101318-83f621788ca8`.

This directory contains that upstream module, including its original licenses.
The only Go source change is in `tls.go`: remove the server rejection of a
ClientHello without an X25519MLKEM768 key share before the optional X25519 share.

Clients with a classical X25519 key share can authenticate again. The hybrid
key share parser and TLS key exchange implementation are unchanged, so clients
that offer X25519MLKEM768 retain hybrid key exchange support. All other REALITY
authentication checks remain in place.

When updating Xray or REALITY, update this copy and review the compatibility
patch against the new upstream implementation. The root `go.mod` uses this
local module through a `replace` directive.
