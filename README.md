# Go Signer

An example Go project utilizing SHA-512 and dumping internal state:

```
$ go run go-signer.go
[115 104 97 7 106 9 230 103 243 188 201 8 187 103 174 133 132 202 167 59 60 110 243 114 254 148 248 43 165 79 245 58 95 29 54 241 81 14 82 127 173 230 130 209 155 5 104 140 43 62 108 31 31 131 217 171 251 65 189 107 91 224 205 25 19 126 33 121 116 104 105 115 32 105 115 32 115 97 109 112 108 101 32 116 101 120 116 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 19]
[20 253 69 133 146 76 11 4 191 13 150 196 9 97 21 35 186 95 254 59 148 60 88 155 127 203 151 216 11 16 228 73 113 23 115 110 198 42 109 92 23 33 70 71 136 219 73 238 135 13 223 117 215 69 243 33 125 109 95 121 213 44 212 166]
```

> Compatibility: Any future changes to hash or crypto packages will endeavor to maintain compatibility with state encoded using previous versions. That is, any released versions of the packages should be able to decode data written with any previously released version, subject to issues such as security fixes. See the Go compatibility document for background: https://golang.org/doc/go1compat 

[Source](https://pkg.go.dev/hash#pkg-types)
