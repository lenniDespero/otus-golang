# String unpacker

## Usage
clone or download sthis sources
run `$ go run unpacker.go`

Output will be like: 

    a4bc2d5e => aaaabccddddde (✅  unpacked)
    abcd => abcd (✅  unpacked)
    45 =>  (❌  not unpacked) : Something strange with unpacking string '45' on character - 4, position - 0.
    qwe\4\5 => qwe45 (✅  unpacked)
    qwe\45 => qwe44444 (✅  unpacked)
    qwe\\5 => qwe\\\\\ (✅  unpacked)
    qwe\42 => qwe44 (✅  unpacked)
