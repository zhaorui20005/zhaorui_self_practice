package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"time"

	"github.com/linkedin/goavro/v2"
)

func main() {
	// Read schema from file
	schemaJSON, err := ioutil.ReadFile("schema2.json")
	if err != nil {
		fmt.Printf("Failed to read schema file: %v\n", err)
		return
	}

	// Parse schema
	codec, err := goavro.NewCodec(string(schemaJSON))
	if err != nil {
		fmt.Printf("Failed to parse schema: %v\n", err)
		return
	}

	// Construct sample data matching the schema
	sampleData := map[string]interface{}{
		"empid":          "E12345",
		"lastname":       "Doe",
		"firstname":      "John",
		"commshonly":     int32(500),
		"salary":         75000,
		"rate":           big.NewRat(12345, 100),                                 // Decimal with scale 2 (123.45)
		"hsalary":        big.NewRat(1234567890, 10000000000),                    // Decimal with scale 10 (0.1234567890)
		"__op":           map[string]interface{}{"string": "c"},                  // Union: ["null", "string"]
		"__ts_ms":        map[string]interface{}{"long": time.Now().UnixMilli()}, // Union: ["null", "long"]
		"__source_db":    map[string]interface{}{"string": "dbonline"},           // Union: ["null", "string"]
		"__source_name":  map[string]interface{}{"string": "employee_table"},     // Union: ["null", "string"]
		"__source_ts_ms": map[string]interface{}{"long": time.Now().UnixMilli()}, // Union: ["null", "long"]
		"__deleted":      map[string]interface{}{"string": "false"},              // Union: ["null", "string"]
		"__dbport":       map[string]interface{}{"string": "3306"},               // Union: ["null", "string"]
		"salaryArray": []interface{}{
			big.NewRat(10000, 100), // 100.00
			big.NewRat(20000, 100), // 200.00
		},
		"salaryMap": map[string]interface{}{
			"foo": big.NewRat(30000, 100), // 300.00
			"feb": big.NewRat(31000, 100), // 310.00
		},
	}

	// Encode to Avro binary
	binary, err := codec.BinaryFromNative(nil, sampleData)
	if err != nil {
		fmt.Printf("Failed to encode data: %v\n", err)
		return
	}
	fmt.Printf("Encoded binary data length: %d bytes\n", len(binary))
	err = ioutil.WriteFile("output.avro", binary, 0644)
	if err != nil {
		fmt.Printf("Failed to write binary data to file: %v\n", err)
		return
	}
	fmt.Println("Binary data written to output.avro")

	// Decode back to native Go data
	native, _, err := codec.NativeFromBinary(binary)
	if err != nil {
		fmt.Printf("Failed to decode data: %v\n", err)
		return
	}

	// Pretty print the decoded data
	nativeJSON, err := json.MarshalIndent(native, "", "  ")
	if err != nil {
		fmt.Printf("Failed to marshal decoded data: %v\n", err)
		return
	}
	fmt.Println("Decoded data:")
	fmt.Println(string(nativeJSON))

	// Read schema from file
	schemaJSON2, err := ioutil.ReadFile("schema.json")
	if err != nil {
		fmt.Printf("Failed to read schema file: %v\n", err)
		return
	}
	// Parse schema
	codec2, err := goavro.NewCodec(string(schemaJSON2))
	if err != nil {
		fmt.Printf("Failed to parse schema: %v\n", err)
		return
	}

	// Decode back to native Go data
	native2, _, err := codec2.NativeFromBinary(binary)
	if err != nil {
		fmt.Printf("Failed to decode data: %v\n", err)
		return
	}

	// Pretty print the decoded data
	nativeJSON2, err := json.MarshalIndent(native2, "", "  ")
	if err != nil {
		fmt.Printf("Failed to marshal decoded data: %v\n", err)
		return
	}
	fmt.Println("Decoded data:")
	fmt.Println(string(nativeJSON2))

}
