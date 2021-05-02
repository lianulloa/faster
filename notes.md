# Notes about protobuf encoding size

## Encoding

<https://developers.google.com/protocol-buffers/docs/encoding>

In protobuf, a string field, when encoded uses string.length + 2 bytes

for example:
<https://developers.google.com/protocol-buffers/docs/encoding#strings>

```protobuf
message Test2 {
  optional string b = 2;
}
```

will be encoded as

12 07 [74 65 73 74 69 6e 67]

0x12 defines the wire type <https://developers.google.com/protocol-buffers/docs/encoding#structure>

whereas 0x07 defines the string length

the rest are the UTF8 of "testing"

## Types

The `bytes` type of protobuf translate into a `[]byte` in go. May contain any arbitrary sequence of bytes no longer than 2^32(i.e 4GB).

## Context

Only one context (and one deadline as consequence) can be defined per connection. I read something about ClientInterceptor, but for now I don't know what it is.
