// Copyright 2019 Filip Kroča. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package teltonikaparser

import (
	"encoding/hex"
	"fmt"
	"log"
	"testing"
)

func ExampleDecode() {
	var bs = []byte{00, 0x7C, 0xCA, 0xFE, 0x01, 0x33, 0x00, 0x0F, 0x33, 0x35, 0x32, 0x30, 0x39, 0x34, 0x30, 0x38, 0x31, 0x36, 0x37, 0x32, 0x31, 0x37, 0x39, 0x08, 0x02, 0x00, 0x00, 0x01, 0x6C, 0x32, 0xB4, 0x88, 0xA0, 0x00, 0x0A, 0x7A, 0x36, 0x7C, 0x1D, 0x30, 0x01, 0x87, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xF1, 0x07, 0x03, 0x01, 0x00, 0x15, 0x00, 0xEF, 0x00, 0x03, 0x42, 0x31, 0x8B, 0xCD, 0x42, 0xDC, 0xCE, 0x60, 0x64, 0x01, 0xF1, 0x00, 0x00, 0x59, 0xD9, 0x00, 0x00, 0x00, 0x01, 0x6C, 0x32, 0xB4, 0x8C, 0x88, 0x00, 0x0A, 0x7A, 0x36, 0x7C, 0x1D, 0x30, 0x01, 0x87, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x15, 0x07, 0x03, 0x01, 0x00, 0x15, 0x01, 0xEF, 0x00, 0x03, 0x42, 0x31, 0x95, 0xCD, 0x42, 0xDC, 0xCE, 0x60, 0x64, 0x01, 0xF1, 0x00, 0x00, 0x59, 0xD9, 0x00, 0x02}

	// decode a raw data byte slice
	parsedData, err := Decode(&bs)
	if err != nil {
		log.Panicf("Error when decoding a bs, %v\n", err)
	}
	fmt.Printf("Decoded packet codec 8:\n%+v\n", parsedData)

	// Output:
	// {IMEI:352094081672179 CodecID:8 NoOfData:2 Data:[{UtimeMs:1564218788000 Utime:1564218788 Priority:0 Lat:175781500 Lng:489685383 Altitude:0 Angle:0 VisSat:0 Speed:0 EventID:241 Elements:[{Length:1 IOID:1 Value:[0]} {Length:1 IOID:21 Value:[0]} {Length:1 IOID:239 Value:[0]} {Length:2 IOID:66 Value:[49 139]} {Length:2 IOID:205 Value:[66 220]} {Length:2 IOID:206 Value:[96 100]} {Length:4 IOID:241 Value:[0 0 89 217]}]} {UtimeMs:1564218789000 Utime:1564218789 Priority:0 Lat:175781500 Lng:489685383 Altitude:0 Angle:0 VisSat:0 Speed:0 EventID:21 Elements:[{Length:1 IOID:1 Value:[0]} {Length:1 IOID:21 Value:[1]} {Length:1 IOID:239 Value:[0]} {Length:2 IOID:66 Value:[49 149]} {Length:2 IOID:205 Value:[66 220]} {Length:2 IOID:206 Value:[96 100]} {Length:4 IOID:241 Value:[0 0 89 217]}]}]}

	// test with Codec8 Extended packet
	stringData := `0086cafe0101000f3335323039333038353639383230368e0100000167efa919800200000000000000000000000000000000fc0013000800ef0000f00000150500c80000450200010000710000fc00000900b5000000b600000042305600cd432a00ce6064001100090012ff22001303d1000f0000000200f1000059d900100000000000000000010086cafe0191000f3335323039333038353639383230368e0100000167efad92080200000000000000000000000000000000fc0013000800ef0000f00000150500c80000450200010000715800fc01000900b5000000b600000042039d00cd432a00ce60640011015f0012fd930013036f000f0000000200f1000059d900100000000000000000010086cafe01a0000f3335323039333038353639383230368e01000000f9cebaeac80200000000000000000000000000000000fc0013000800ef0000f00000150000c80000450200010000710000fc00000900b5000000b600000042305400cd000000ce0000001103570012fe8900130196000f0000000200f10000000000100000000000000000010083cafe0101000f3335323039333038353639383230368e0100000167f1aeec00000a750e8f1d43443100f800b210000000000012000700ef0000f00000150500c800004501000100007142000900b5000600b6000500422fb300cd432a00ce60640011000700120007001303ec000f0000000200f1000059d90010000000000000000001`

	bs, _ = hex.DecodeString(stringData)

	// decode a raw data byte slice
	parsedData, err = Decode(&bs)
	if err != nil {
		log.Panicf("Error when decoding a bs, %v\n", err)
	}
	fmt.Printf("Decoded packet codec 8 extended:\n%+v\n", parsedData)

	// Output:
	// Decoded packet codec 8:
	//{IMEI:352094081672179 CodecID:8 NoOfData:2 Data:[{UtimeMs:1564218788000 Utime:1564218788 Priority:0 Lat:175781500 Lng:489685383 Altitude:0 Angle:0 VisSat:0 Speed:0 EventID:241 Elements:[{Length:1 IOID:1 Value:[0]} {Length:1 IOID:21 Value:[0]} {Length:1 IOID:239 Value:[0]} {Length:2 IOID:66 Value:[49 139]} {Length:2 IOID:205 Value:[66 220]} {Length:2 IOID:206 Value:[96 100]} {Length:4 IOID:241 Value:[0 0 89 217]}]} {UtimeMs:1564218789000 Utime:1564218789 Priority:0 Lat:175781500 Lng:489685383 Altitude:0 Angle:0 VisSat:0 Speed:0 EventID:21 Elements:[{Length:1 IOID:1 Value:[0]} {Length:1 IOID:21 Value:[1]} {Length:1 IOID:239 Value:[0]} {Length:2 IOID:66 Value:[49 149]} {Length:2 IOID:205 Value:[66 220]} {Length:2 IOID:206 Value:[96 100]} {Length:4 IOID:241 Value:[0 0 89 217]}]}]}
	//Decoded packet codec 8 extended:
	//{IMEI:352093085698206 CodecID:142 NoOfData:1 Data:[{UtimeMs:1545914096000 Utime:1545914096 Priority:2 Lat:0 Lng:0 Altitude:0 Angle:0 VisSat:0 Speed:0 EventID:252 Elements:[{Length:1 IOID:239 Value:[0]} {Length:1 IOID:240 Value:[0]} {Length:1 IOID:21 Value:[5]} {Length:1 IOID:200 Value:[0]} {Length:1 IOID:69 Value:[2]} {Length:1 IOID:1 Value:[0]} {Length:1 IOID:113 Value:[0]} {Length:1 IOID:252 Value:[0]} {Length:2 IOID:181 Value:[0 0]} {Length:2 IOID:182 Value:[0 0]} {Length:2 IOID:66 Value:[48 86]} {Length:2 IOID:205 Value:[67 42]} {Length:2 IOID:206 Value:[96 100]} {Length:2 IOID:17 Value:[0 9]} {Length:2 IOID:18 Value:[255 34]} {Length:2 IOID:19 Value:[3 209]} {Length:2 IOID:15 Value:[0 0]} {Length:4 IOID:241 Value:[0 0 89 217]} {Length:4 IOID:16 Value:[0 0 0 0]}]}]}
}

func ExampleHuman() {

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

	// Output:
	// Property Name: Ignition, Value: 0
	// Property Name: Movement, Value: 0
	// Property Name: GSM Signal, Value: 5
	// Property Name: Sleep Mode, Value: 0
	// Property Name: GNSS Status, Value: 2
	// Property Name: Digital Input 1, Value: false
	// Property Name: Battery Level, Value: 0
	// Property Name: Unplug, Value: 0
	// Property Name: GNSS PDOP, Value: 0
	// Property Name: GNSS HDOP, Value: 0
	// Property Name: External Voltage, Value: 12374
	// Property Name: GSM Cell ID, Value: 17194
	// Property Name: GSM Area Code, Value: 24676
	// Property Name: Axis X, Value: 9
	// Property Name: Axis Y, Value: -222
	// Property Name: Axis Z, Value: 977
	// Property Name: Eco Score, Value: 0
	// Property Name: Active GSM Operator, Value: 23001
	// Property Name: Total Odometer, Value: 0
}

func BenchmarkDecode(b *testing.B) {
	stringData := `0086cafe0101000f3335323039333038353639383230368e0100000167efa919800200000000000000000000000000000000fc0013000800ef0000f00000150500c80000450200010000710000fc00000900b5000000b600000042305600cd432a00ce6064001100090012ff22001303d1000f0000000200f1000059d900100000000000000000010086cafe0191000f3335323039333038353639383230368e0100000167efad92080200000000000000000000000000000000fc0013000800ef0000f00000150500c80000450200010000715800fc01000900b5000000b600000042039d00cd432a00ce60640011015f0012fd930013036f000f0000000200f1000059d900100000000000000000010086cafe01a0000f3335323039333038353639383230368e01000000f9cebaeac80200000000000000000000000000000000fc0013000800ef0000f00000150000c80000450200010000710000fc00000900b5000000b600000042305400cd000000ce0000001103570012fe8900130196000f0000000200f10000000000100000000000000000010083cafe0101000f3335323039333038353639383230368e0100000167f1aeec00000a750e8f1d43443100f800b210000000000012000700ef0000f00000150500c800004501000100007142000900b5000600b6000500422fb300cd432a00ce60640011000700120007001303ec000f0000000200f1000059d90010000000000000000001`

	bs, _ := hex.DecodeString(stringData)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := Decode(&bs)
		if err != nil {
			log.Panicf("Error when decoding a bs, %v\n", err)
		}
	}
}

func BenchmarkHuman(b *testing.B) {
	stringData := `0086cafe0101000f3335323039333038353639383230368e0100000167efa919800200000000000000000000000000000000fc0013000800ef0000f00000150500c80000450200010000710000fc00000900b5000000b600000042305600cd432a00ce6064001100090012ff22001303d1000f0000000200f1000059d900100000000000000000010086cafe0191000f3335323039333038353639383230368e0100000167efad92080200000000000000000000000000000000fc0013000800ef0000f00000150500c80000450200010000715800fc01000900b5000000b600000042039d00cd432a00ce60640011015f0012fd930013036f000f0000000200f1000059d900100000000000000000010086cafe01a0000f3335323039333038353639383230368e01000000f9cebaeac80200000000000000000000000000000000fc0013000800ef0000f00000150000c80000450200010000710000fc00000900b5000000b600000042305400cd000000ce0000001103570012fe8900130196000f0000000200f10000000000100000000000000000010083cafe0101000f3335323039333038353639383230368e0100000167f1aeec00000a750e8f1d43443100f800b210000000000012000700ef0000f00000150500c800004501000100007142000900b5000600b6000500422fb300cd432a00ce60640011000700120007001303ec000f0000000200f1000059d90010000000000000000001`

	bs, _ := hex.DecodeString(stringData)
	// initialize a human decoder
	humanDecoder := HumanDecoder{}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		// decode a raw data byte slice
		parsedData, err := Decode(&bs)
		if err != nil {
			log.Panicf("Error when decoding a bs, %v\n", err)
		}

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
				if _, err := (*decoded).GetFinalValue(); err != nil {
					log.Panicf("Unable to GetFinalValue() %v", err)
				}
			}
		}
	}
}
