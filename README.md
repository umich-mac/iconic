# iconic

`iconic` is an alternative tool for creating Munki's `_icons_hash.plist` file. It's written
in go to be portable and not require a runtime.

`go build .` to produce the binary.

`./iconic > _icon_hashes.plist` to create the hash.

Probably needs better understanding of file paths someday.
