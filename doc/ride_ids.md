1. some value is stored in ride_error_string

2. 0x100 is subtracted from the value

3. the value is multiplied by 2

4. this is an index into ride_configuration_string_ids

example: 159h

- subtract 0x100 to get 0x59

- multiply by 2 to get 0xb2

get pointer in ride_configuration_string_ids

## reverse

start with 0xb2

- divide by 2 to get 0x59 (?)

- add 0x100
