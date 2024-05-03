# Types

Each packet carries different types of data rather than having one packet which contains everything. The header in each packet describes the packet type and versioning info so it will be easier for applications to check they are interpreting the incoming data in the correct way. Please note that all values are encoded using `Little Endian` format. All data is packed.

The following data types are used in the structures:


| Type   | Description                              | Go type |
| ------ | ---------------------------------------- | ------- |
| uint8  | Unsigned 8-bit integer                   | uint8   |
| int8   | Signed 8-bit integer                     | int8    |
| uint16 | Unsigned 16-bit integer                  | uint16  |
| int16  | Signed 16-bit integer                    | int16   |
| uint32 | Unsigned 32-bit integer                  | uint32  |
| float  | Floating point (32-bit)                  | float32 |
| Double | Double-precision floating point (64-bit) | float64 |
| uint64 | Unsigned 64-bit integer                  | uint64  |
| char   | Character                                | char    |
