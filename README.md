# Package teltonikaparser provides GO parser and validator for Teltonika Codec 8 and Codec 8 Extended

Certain purpose:

Package teltonikaparser was created for parsing data structures from [Teltonika](https://wiki.teltonika.lt/view/Codec#Codec_8) UDP packets. Package can return a raw data and human readable data, see [examples](#example).

Package teltonikaparser is a very fast, low-level implementation, it can decode over **one milion packets** per second per core. See [bechmarks](https://godoc.org/github.com/filipkroca/teltonikaparser#benchmark-Decode)

Performace:
Decode()   788 ns/op 592 B/op 4 allocs/op

## First stage - basic decoding

When a binary packet arrived it is necessary to parse the data out and create a structure which represents a parsed data.

### type Decoded

```go
type Decoded struct {
    IMEI     string    // IMEI number, if len==15 also validated by checksum
    CodecID  byte      // 0x08 (codec 8) or 0x8E (codec 8 extended)
    NoOfData uint8     // Number of Data
    Data     []AvlData // Slice with avl data
}
```

### type AvlData

AvlData represent one block of data.

```go
type AvlData struct {
    UtimeMs    uint64      // Utime in mili seconds
    Utime      uint64      // Utime in seconds
    Priority   uint8       // Priority, [0 Low, 1 High, 2 Panic]
    Lat        int32       // Latitude (between 850000000 and -850000000), fit int32
    Lng        int32       // Longitude (between 1800000000 and -1800000000), fit int32
    Altitude   int16       // Altitude In meters above sea level, 2 bytes
    Angle      uint16      // Angle In degrees, 0 is north, increasing clock-wise, 2 bytes
    VisSat     uint8       // Satellites Number of visible satellites
    Speed      uint16      // Speed in km/h
    EventID    uint16      // Event generated (0 – data generated not on event)
    Elements []Element // Slice containing parsed IO Elements
}
```

### type Element

Element represents one IO element parsed from a binary packet.

```go
type Element struct {
   Length uint16 // Length of element, this should be uint16 because Codec 8 extended has 2Byte of IO len
   IOID   uint16 // IO element ID
   Value  []byte // Value of the element represented by slice of bytes
}
```

### func Decode

Decode is used for basic decoding as see in the example. It takes a pointer to a byte slice and return Decoded struct and error. [FULL DOCUMENTATION](https://godoc.org/github.com/filipkroca/teltonikaparser#Decode) 

Performance per core: 849 ns/op 720 B/op 3 allocs/op

### Example Decode

```go
package main

import (
   "fmt"
   "log"

    "github.com/filipkroca/teltonikaparser"
)

func main() {
    // Example packet Teltonika UDP Codec 8 Extended 0086cafe0101000f3335323039333038353639383230368e0100000167efa919800200000000000000000000000000000000fc0013000800ef0000f00000150500c80000450200010000710000fc00000900b5000000b600000042305600cd432a00ce6064001100090012ff22001303d1000f0000000200f1000059d900100000000000000000010086cafe0191000f3335323039333038353639383230368e0100000167efad92080200000000000000000000000000000000fc0013000800ef0000f00000150500c80000450200010000715800fc01000900b5000000b600000042039d00cd432a00ce60640011015f0012fd930013036f000f0000000200f1000059d900100000000000000000010086cafe01a0000f3335323039333038353639383230368e01000000f9cebaeac80200000000000000000000000000000000fc0013000800ef0000f00000150000c80000450200010000710000fc00000900b5000000b600000042305400cd000000ce0000001103570012fe8900130196000f0000000200f10000000000100000000000000000010083cafe0101000f3335323039333038353639383230368e0100000167f1aeec00000a750e8f1d43443100f800b210000000000012000700ef0000f00000150500c800004501000100007142000900b5000600b6000500422fb300cd432a00ce60640011000700120007001303ec000f0000000200f1000059d90010000000000000000001

    // test with Codec8 Extended packet
    stringData := `0086cafe0101000f3335323039333038353639383230368e0100000167efa919800200000000000000000000000000000000fc0013000800ef0000f00000150500c80000450200010000710000fc00000900b5000000b600000042305600cd432a00ce6064001100090012ff22001303d1000f0000000200f1000059d900100000000000000000010086cafe0191000f3335323039333038353639383230368e0100000167efad92080200000000000000000000000000000000fc0013000800ef0000f00000150500c80000450200010000715800fc01000900b5000000b600000042039d00cd432a00ce60640011015f0012fd930013036f000f0000000200f1000059d900100000000000000000010086cafe01a0000f3335323039333038353639383230368e01000000f9cebaeac80200000000000000000000000000000000fc0013000800ef0000f00000150000c80000450200010000710000fc00000900b5000000b600000042305400cd000000ce0000001103570012fe8900130196000f0000000200f10000000000100000000000000000010083cafe0101000f3335323039333038353639383230368e0100000167f1aeec00000a750e8f1d43443100f800b210000000000012000700ef0000f00000150500c800004501000100007142000900b5000600b6000500422fb300cd432a00ce60640011000700120007001303ec000f0000000200f1000059d90010000000000000000001`

    bs, _ := hex.DecodeString(stringData)

    // decode a raw data byte slice
    parsedData, err := Decode(&bs)
    if err != nil {
        log.Panicf("Error when decoding a bs, %v\n", err)
    }

    // initialize a human decoder
    humanDecoder := HumanDecoder{}

    // loop over raw data
    for _, val := range parsedData.Data {
        // loop over Elements
        for _, ioel := range val.Elements {
            // decode to human readable format
            decoded, err := humanDecoder.Human(&ioel, "FMBXY") // second parameter - device family type ["FMBXY", "FM64"]
            if err != nil {
                log.Panicf("Error when converting human, %v\n", err)
            }

            // get final decoded value to value which is specified in ./teltonikajson/ in paramether FinalConversion
            if val, err := (*decoded).GetFinalValue(); err != nil {
                log.Panicf("Unable to GetFinalValue() %v", err)
            } else if val != nil {
                // print output
                fmt.Printf("Property Name: %v, Value: %v\n", decoded.AvlEncodeKey.PropertyName, val)
            }
        }
    }
}
```

Output:  

```text
Property Name: Ignition, Value: 0  
Property Name: Movement, Value: 0  
Property Name: GSM Signal, Value: 5  
Property Name: Sleep Mode, Value: 0  
Property Name: GNSS Status, Value: 2  
Property Name: Digital Input 1, Value: false  
Property Name: Battery Level, Value: 0  
Property Name: Unplug, Value: 0  
Property Name: GNSS PDOP, Value: 0  
Property Name: GNSS HDOP, Value: 0  
Property Name: External Voltage, Value: 12374  
Property Name: GSM Cell ID, Value: 17194  
Property Name: GSM Area Code, Value: 24676  
Property Name: Axis X, Value: 9  
Property Name: Axis Y, Value: -222  
Property Name: Axis Z, Value: 977  
Property Name: Eco Score, Value: 0  
Property Name: Active GSM Operator, Value: 23001  
Property Name: Total Odometer, Value: 0  
```

## Human readable

This package also provides method (h *HAvlData) GetFinalValue() which can convert values to human-readable form.

Currently are implemented AVLs lists for [FMBXY](https://wiki.teltonika.lt/view/FMB_AVL_ID) and [FMB64](https://wiki.teltonika.lt/view/FMB64_AVL_ID) devices family.

### type HumanDecoder

HumanDecoder is responsible for decoding

```go
type HumanDecoder struct {
    elements map[string]map[uint16]AvlEncodeKey
}
```

### type HAvlData

HAvlData represent human readable set of a pointer to an AvlEncodeKey Decoding key and a pointer to IO element with RAW data

```go
type HAvlData struct {
    AvlEncodeKey *AvlEncodeKey
    Element      *Element
}
```

### type AvlEncodeKey

AvlEncodeKey represent parsed element values from JSON

```go
type AvlEncodeKey struct {
    No              string `json:"No"`
    PropertyName    string `json:"PropertyName"`
    Bytes           string `json:"Bytes"`
    Type            string `json:"Type"`
    Min             string `json:"Min"`
    Max             string `json:"Max"`
    Multiplier      string `json:"Multiplier"`
    Units           string `json:"Units"`
    Description     string `json:"Description"`
    HWSupport       string `json:"HWSupport"`
    ParametrGroup   string `json:"Parametr Group"`
    FinalConversion string `json:"FinalConversion"`
}
```

### Example HumanDecoder

Have a binary packet bs which is Teltonika UDP Codec 8 packet

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
        // loop over Elements
        for _, ioel := range val.Elements {
                // decode to human readable format
                decoded, err := humanDecoder.Human(&ioel, "FMBXY") // second parameter - device family type ["FMBXY", "FM64"]
                if err != nil {
                log.Printf("Hoops, human, %v\n", err)
                return
                }
                // get final decoded value to value which is specified in ./teltonikajson/ in paramether FinalConversion
                if val, err := (*decoded).GetFinalValue(); err != nil {
                log.Panicf("Unable to GetFinalValue() %v", err)
                } else if val != nil {
                // print output
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
