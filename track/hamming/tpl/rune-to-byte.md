- You are currently converting a `rune` to a `byte`. This conversion is lossy if we have to deal with 
special characters. A rune is `int32` a byte only `uint8`. You could instead loop over bytes by converting 
the `string` to a slice of bytes `[]byte` before the loop.