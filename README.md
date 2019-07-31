# Package teltonikaparser provides GO parser and validator for Teltonika Codec 8 and Codec 8 Extended

Certain purpose:

Package teltonikaparser was created for parsing data structures from [Teltonika](https://wiki.teltonika.lt/view/Codec#Codec_8) UDP packets. Package can return a raw data and human readable data, see [examples]().

Package teltonikaparser is a very fast, low-level implementation, it can decode over **one milion packets** per second per core. See [bechmarks](https://godoc.org/github.com/filipkroca/b2n#benchmark-Decode)

Performace:
Decode()   788 ns/op 592 B/op 4 allocs/op

## Human readable

This package also provides method (h *HAvlData) GetFinalValue() which can convert values to human-readable form.

Currently are implemented AVLs lists for [FMBXY](https://wiki.teltonika.lt/view/FMB_AVL_ID) and [FMB64](https://wiki.teltonika.lt/view/FMB64_AVL_ID) devices family.

For example, have binary packet bs which is Teltonika UDP Codec 8 packet

```go
package main

import (
    "encoding/hex"
    "fmt"

    "github.com/filipkroca/teltonikaparser"
)

func main() {
   // Example packet Teltonika UDP Codec 8 007CCAFE0133000F33353230393430383136373231373908020000016C32B488A0000A7A367C1D30018700000000000000F1070301001500EF000342318BCD42DCCE606401F1000059D9000000016C32B48C88000A7A367C1D3001870000000000000015070301001501EF0003423195CD42DCCE606401F1000059D90002

   var bs = []byte{00, 0x7C, 0xCA, 0xFE, 0x01, 0x33, 0x00, 0x0F, 0x33, 0x35, 0x32, 0x30, 0x39, 0x34, 0x30, 0x38, 0x31, 0x36, 0x37, 0x32, 0x31, 0x37, 0x39, 0x08, 0x02, 0x00, 0x00, 0x01, 0x6C, 0x32, 0xB4, 0x88, 0xA0, 0x00, 0x0A, 0x7A, 0x36, 0x7C, 0x1D, 0x30, 0x01, 0x87, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xF1, 0x07, 0x03, 0x01, 0x00, 0x15, 0x00, 0xEF, 0x00, 0x03, 0x42, 0x31, 0x8B, 0xCD, 0x42, 0xDC, 0xCE, 0x60, 0x64, 0x01, 0xF1, 0x00, 0x00, 0x59, 0xD9, 0x00, 0x00, 0x00, 0x01, 0x6C, 0x32, 0xB4, 0x8C, 0x88, 0x00, 0x0A, 0x7A, 0x36, 0x7C, 0x1D, 0x30, 0x01, 0x87, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x15, 0x07, 0x03, 0x01, 0x00, 0x15, 0x01, 0xEF, 0x00, 0x03, 0x42, 0x31, 0x95, 0xCD, 0x42, 0xDC, 0xCE, 0x60, 0x64, 0x01, 0xF1, 0x00, 0x00, 0x59, 0xD9, 0x00, 0x02}

   // initialize human decoder
   humanDecoder := HumanDecoder{}

   // fire go routine
   go func() {
    // decode raw data
    decoded, err := Decode(&bs)
    if err != nil {
        fmt.Println("ExampleDecode error", err)
    }

    // loop over raw data
    for _, val := range decoded.Data {
        // loop over IOElements
        for _, ioel := range val.IOElements {
                // decode to human readable format
                decoded, err := humanDecoder.Human(&ioel, "FMBXY")
                if err != nil {
                log.Printf("Hoops, human, %v\n", err)
                return
                }
                // get final decoded value to value which is specified in ./teltonikajson/ in paramether FinalConversion
                if val, err := (*decoded).GetFinalValue(); err != nil {
                log.Panicf("Unable to GetFinalValue() %v", err)
                } else if val != nil {
                fmt.Printf("%v : %v %v multiplier %v\n", (*decoded).AvlIO.PropertyName, val, (*decoded).AvlIO.Units, (*decoded).AvlIO.Multiplier)
                }
        }
    }
   }()
}
```

Decoded output should be

```text
Output:
GSM Signal : 0 - multiplier -
Ignition : 0 - multiplier -
External Voltage : 12683 mV multiplier -
GSM Cell ID : 17116 - multiplier -
GSM Area Code : 24676 - multiplier -
Active GSM Operator : 23001 - multiplier -
GSM Signal : 1 - multiplier -
Ignition : 0 - multiplier -
External Voltage : 12693 mV multiplier -
GSM Cell ID : 17116 - multiplier -
GSM Area Code : 24676 - multiplier -
Active GSM Operator : 23001 - multiplier -
```

Full documentation [HERE](https://godoc.org/github.com/filipkroca/teltonikaparser)
