# emvutil

## Description
emvutil is a utility tool for QR code payloads encoded according to the EMV MPM specification.

## Usage
### Decoding
Decode a payload. Decoded payload is represented as a struct defined in [mercari/go-emv-code](https://godoc.org/go.mercari.io/go-emv-code/mpm#Code).

``` sh
$ emvutil decode <payload>
```

`decode` command also supports JPQR ID (統一店舗識別コード) decoding. [JPQR](https://www.paymentsjapan.or.jp/wordpress/wp-content/uploads/2019/03/MPM_Guideline_1.0.pdf) is a QR code specification based on the EMV MPM specification.

``` sh
$ emvutil decode --jpqr-id <jpqr-id>
```

`decode` command supports JSON formatted output with `--json` flag.
``` sh
$ emvutil decode --json <payload>
```
