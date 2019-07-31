// Copyright 2019 Filip Kroča. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package teltonikaparser

import (
	"fmt"
	"log"
	"testing"
)

func ExampleDecode() {
	// Example packet Teltonika UDP Codec 8 007CCAFE0133000F33353230393430383136373231373908020000016C32B488A0000A7A367C1D30018700000000000000F1070301001500EF000342318BCD42DCCE606401F1000059D9000000016C32B48C88000A7A367C1D3001870000000000000015070301001501EF0003423195CD42DCCE606401F1000059D90002

	var bs = []byte{00, 0x7C, 0xCA, 0xFE, 0x01, 0x33, 0x00, 0x0F, 0x33, 0x35, 0x32, 0x30, 0x39, 0x34, 0x30, 0x38, 0x31, 0x36, 0x37, 0x32, 0x31, 0x37, 0x39, 0x08, 0x02, 0x00, 0x00, 0x01, 0x6C, 0x32, 0xB4, 0x88, 0xA0, 0x00, 0x0A, 0x7A, 0x36, 0x7C, 0x1D, 0x30, 0x01, 0x87, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xF1, 0x07, 0x03, 0x01, 0x00, 0x15, 0x00, 0xEF, 0x00, 0x03, 0x42, 0x31, 0x8B, 0xCD, 0x42, 0xDC, 0xCE, 0x60, 0x64, 0x01, 0xF1, 0x00, 0x00, 0x59, 0xD9, 0x00, 0x00, 0x00, 0x01, 0x6C, 0x32, 0xB4, 0x8C, 0x88, 0x00, 0x0A, 0x7A, 0x36, 0x7C, 0x1D, 0x30, 0x01, 0x87, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x15, 0x07, 0x03, 0x01, 0x00, 0x15, 0x01, 0xEF, 0x00, 0x03, 0x42, 0x31, 0x95, 0xCD, 0x42, 0xDC, 0xCE, 0x60, 0x64, 0x01, 0xF1, 0x00, 0x00, 0x59, 0xD9, 0x00, 0x02}

	decoded, err := Decode(&bs)
	if err != nil {
		fmt.Println("ExampleDecode error", err)
	}

	fmt.Printf("%+v", decoded)

	// Output:
	// {IMEI:352094081672179 CodecID:8 NoOfData:2 Data:[{UtimeMs:1564218788000 Utime:1564218788 Priority:0 Lat:175781500 Lng:489685383 Altitude:0 Angle:0 VisSat:0 Speed:0 EventID:241 IOElements:[{Length:1 IOID:1 Value:[0]} {Length:1 IOID:21 Value:[0]} {Length:1 IOID:239 Value:[0]} {Length:2 IOID:66 Value:[49 139]} {Length:2 IOID:205 Value:[66 220]} {Length:2 IOID:206 Value:[96 100]} {Length:4 IOID:241 Value:[0 0 89 217]}]} {UtimeMs:1564218789000 Utime:1564218789 Priority:0 Lat:175781500 Lng:489685383 Altitude:0 Angle:0 VisSat:0 Speed:0 EventID:21 IOElements:[{Length:1 IOID:1 Value:[0]} {Length:1 IOID:21 Value:[1]} {Length:1 IOID:239 Value:[0]} {Length:2 IOID:66 Value:[49 149]} {Length:2 IOID:205 Value:[66 220]} {Length:2 IOID:206 Value:[96 100]} {Length:4 IOID:241 Value:[0 0 89 217]}]}]}
}

func ExampleHumanDecoder() {
	// Example packet Teltonika UDP Codec 8 007CCAFE0133000F33353230393430383136373231373908020000016C32B488A0000A7A367C1D30018700000000000000F1070301001500EF000342318BCD42DCCE606401F1000059D9000000016C32B48C88000A7A367C1D3001870000000000000015070301001501EF0003423195CD42DCCE606401F1000059D90002

	var bs = []byte{00, 0x7C, 0xCA, 0xFE, 0x01, 0x33, 0x00, 0x0F, 0x33, 0x35, 0x32, 0x30, 0x39, 0x34, 0x30, 0x38, 0x31, 0x36, 0x37, 0x32, 0x31, 0x37, 0x39, 0x08, 0x02, 0x00, 0x00, 0x01, 0x6C, 0x32, 0xB4, 0x88, 0xA0, 0x00, 0x0A, 0x7A, 0x36, 0x7C, 0x1D, 0x30, 0x01, 0x87, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xF1, 0x07, 0x03, 0x01, 0x00, 0x15, 0x00, 0xEF, 0x00, 0x03, 0x42, 0x31, 0x8B, 0xCD, 0x42, 0xDC, 0xCE, 0x60, 0x64, 0x01, 0xF1, 0x00, 0x00, 0x59, 0xD9, 0x00, 0x00, 0x00, 0x01, 0x6C, 0x32, 0xB4, 0x8C, 0x88, 0x00, 0x0A, 0x7A, 0x36, 0x7C, 0x1D, 0x30, 0x01, 0x87, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x15, 0x07, 0x03, 0x01, 0x00, 0x15, 0x01, 0xEF, 0x00, 0x03, 0x42, 0x31, 0x95, 0xCD, 0x42, 0xDC, 0xCE, 0x60, 0x64, 0x01, 0xF1, 0x00, 0x00, 0x59, 0xD9, 0x00, 0x02}

	// initialize human decoder
	humanDecoder := HumanDecoder{}

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

			fmt.Printf("%v : %v\n", (*decoded).AvlIO.PropertyName, (*decoded).IOElement.Value)

		}
	}

	fmt.Printf("%+v", decoded)

	// Output:
	// Digital Input 1 : [0]
	// GSM Signal : [0]
	// Ignition : [0]
	// External Voltage : [49 139]
	// GSM Cell ID : [66 220]
	// GSM Area Code : [96 100]
	// Active GSM Operator : [0 0 89 217]
	// Digital Input 1 : [0]
	// GSM Signal : [1]
	// Ignition : [0]
	// External Voltage : [49 149]
	// GSM Cell ID : [66 220]
	// GSM Area Code : [96 100]
	// Active GSM Operator : [0 0 89 217]
	// {IMEI:352094081672179 CodecID:8 NoOfData:2 Data:[{UtimeMs:1564218788000 Utime:1564218788 Priority:0 Lat:175781500 Lng:489685383 Altitude:0 Angle:0 VisSat:0 Speed:0 EventID:241 IOElements:[{Length:1 IOID:1 Value:[0]} {Length:1 IOID:21 Value:[0]} {Length:1 IOID:239 Value:[0]} {Length:2 IOID:66 Value:[49 139]} {Length:2 IOID:205 Value:[66 220]} {Length:2 IOID:206 Value:[96 100]} {Length:4 IOID:241 Value:[0 0 89 217]}]} {UtimeMs:1564218789000 Utime:1564218789 Priority:0 Lat:175781500 Lng:489685383 Altitude:0 Angle:0 VisSat:0 Speed:0 EventID:21 IOElements:[{Length:1 IOID:1 Value:[0]} {Length:1 IOID:21 Value:[1]} {Length:1 IOID:239 Value:[0]} {Length:2 IOID:66 Value:[49 149]} {Length:2 IOID:205 Value:[66 220]} {Length:2 IOID:206 Value:[96 100]} {Length:4 IOID:241 Value:[0 0 89 217]}]}]}

}

func ExampleHAvlDataGetFinalValue() {
	// Example packet Teltonika UDP Codec 8 007CCAFE0133000F33353230393430383136373231373908020000016C32B488A0000A7A367C1D30018700000000000000F1070301001500EF000342318BCD42DCCE606401F1000059D9000000016C32B48C88000A7A367C1D3001870000000000000015070301001501EF0003423195CD42DCCE606401F1000059D90002

	var bs = []byte{00, 0x7C, 0xCA, 0xFE, 0x01, 0x33, 0x00, 0x0F, 0x33, 0x35, 0x32, 0x30, 0x39, 0x34, 0x30, 0x38, 0x31, 0x36, 0x37, 0x32, 0x31, 0x37, 0x39, 0x08, 0x02, 0x00, 0x00, 0x01, 0x6C, 0x32, 0xB4, 0x88, 0xA0, 0x00, 0x0A, 0x7A, 0x36, 0x7C, 0x1D, 0x30, 0x01, 0x87, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xF1, 0x07, 0x03, 0x01, 0x00, 0x15, 0x00, 0xEF, 0x00, 0x03, 0x42, 0x31, 0x8B, 0xCD, 0x42, 0xDC, 0xCE, 0x60, 0x64, 0x01, 0xF1, 0x00, 0x00, 0x59, 0xD9, 0x00, 0x00, 0x00, 0x01, 0x6C, 0x32, 0xB4, 0x8C, 0x88, 0x00, 0x0A, 0x7A, 0x36, 0x7C, 0x1D, 0x30, 0x01, 0x87, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x15, 0x07, 0x03, 0x01, 0x00, 0x15, 0x01, 0xEF, 0x00, 0x03, 0x42, 0x31, 0x95, 0xCD, 0x42, 0xDC, 0xCE, 0x60, 0x64, 0x01, 0xF1, 0x00, 0x00, 0x59, 0xD9, 0x00, 0x02}

	// initialize human decoder
	humanDecoder := HumanDecoder{}

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

	// Output:
	// GSM Signal : 0 - multiplier -
	// Ignition : 0 - multiplier -
	// External Voltage : 12683 mV multiplier -
	// GSM Cell ID : 17116 - multiplier -
	// GSM Area Code : 24676 - multiplier -
	// Active GSM Operator : 23001 - multiplier -
	// GSM Signal : 1 - multiplier -
	// Ignition : 0 - multiplier -
	// External Voltage : 12693 mV multiplier -
	// GSM Cell ID : 17116 - multiplier -
	// GSM Area Code : 24676 - multiplier -
	// Active GSM Operator : 23001 - multiplier -
}

func BenchmarkHAvlDataGetFinalValue(b *testing.B) {
	// Example packet Teltonika UDP Codec 8 007CCAFE0133000F33353230393430383136373231373908020000016C32B488A0000A7A367C1D30018700000000000000F1070301001500EF000342318BCD42DCCE606401F1000059D9000000016C32B48C88000A7A367C1D3001870000000000000015070301001501EF0003423195CD42DCCE606401F1000059D90002

	var bs = []byte{00, 0x7C, 0xCA, 0xFE, 0x01, 0x33, 0x00, 0x0F, 0x33, 0x35, 0x32, 0x30, 0x39, 0x34, 0x30, 0x38, 0x31, 0x36, 0x37, 0x32, 0x31, 0x37, 0x39, 0x08, 0x02, 0x00, 0x00, 0x01, 0x6C, 0x32, 0xB4, 0x88, 0xA0, 0x00, 0x0A, 0x7A, 0x36, 0x7C, 0x1D, 0x30, 0x01, 0x87, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xF1, 0x07, 0x03, 0x01, 0x00, 0x15, 0x00, 0xEF, 0x00, 0x03, 0x42, 0x31, 0x8B, 0xCD, 0x42, 0xDC, 0xCE, 0x60, 0x64, 0x01, 0xF1, 0x00, 0x00, 0x59, 0xD9, 0x00, 0x00, 0x00, 0x01, 0x6C, 0x32, 0xB4, 0x8C, 0x88, 0x00, 0x0A, 0x7A, 0x36, 0x7C, 0x1D, 0x30, 0x01, 0x87, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x15, 0x07, 0x03, 0x01, 0x00, 0x15, 0x01, 0xEF, 0x00, 0x03, 0x42, 0x31, 0x95, 0xCD, 0x42, 0xDC, 0xCE, 0x60, 0x64, 0x01, 0xF1, 0x00, 0x00, 0x59, 0xD9, 0x00, 0x02}

	// initialize human decoder
	humanDecoder := HumanDecoder{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
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
				_, err := humanDecoder.Human(&ioel, "FMBXY")
				if err != nil {
					return
				}

			}
		}
	}

}

func BenchmarkDecode(b *testing.B) {
	// Example packet Teltonika UDP Codec 8 007CCAFE0133000F33353230393430383136373231373908020000016C32B488A0000A7A367C1D30018700000000000000F1070301001500EF000342318BCD42DCCE606401F1000059D9000000016C32B48C88000A7A367C1D3001870000000000000015070301001501EF0003423195CD42DCCE606401F1000059D90002

	var bs = []byte{00, 0x7C, 0xCA, 0xFE, 0x01, 0x33, 0x00, 0x0F, 0x33, 0x35, 0x32, 0x30, 0x39, 0x34, 0x30, 0x38, 0x31, 0x36, 0x37, 0x32, 0x31, 0x37, 0x39, 0x08, 0x02, 0x00, 0x00, 0x01, 0x6C, 0x32, 0xB4, 0x88, 0xA0, 0x00, 0x0A, 0x7A, 0x36, 0x7C, 0x1D, 0x30, 0x01, 0x87, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xF1, 0x07, 0x03, 0x01, 0x00, 0x15, 0x00, 0xEF, 0x00, 0x03, 0x42, 0x31, 0x8B, 0xCD, 0x42, 0xDC, 0xCE, 0x60, 0x64, 0x01, 0xF1, 0x00, 0x00, 0x59, 0xD9, 0x00, 0x00, 0x00, 0x01, 0x6C, 0x32, 0xB4, 0x8C, 0x88, 0x00, 0x0A, 0x7A, 0x36, 0x7C, 0x1D, 0x30, 0x01, 0x87, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x15, 0x07, 0x03, 0x01, 0x00, 0x15, 0x01, 0xEF, 0x00, 0x03, 0x42, 0x31, 0x95, 0xCD, 0x42, 0xDC, 0xCE, 0x60, 0x64, 0x01, 0xF1, 0x00, 0x00, 0x59, 0xD9, 0x00, 0x02}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// decode raw data
		decoded, err := Decode(&bs)
		if err != nil {
			fmt.Println("ExampleDecode error", err)
		}
		_ = decoded
	}

}
