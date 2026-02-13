all:
	arduino-cli compile --fqbn esp32:esp32:esp32 --only-compilation-database
	arduino-cli compile --fqbn esp32:esp32:esp32
upload:
	arduino-cli upload -p /dev/ttyUSB0 --fqbn esp32:esp32:esp32  ./
monitor:
	arduino-cli monitor -p /dev/ttyUSB0 --config baudrate=9600
AU:
	arduino-cli compile --fqbn esp32:esp32:esp32 --only-compilation-database
	arduino-cli compile --fqbn esp32:esp32:esp32
	arduino-cli upload -p /dev/ttyUSB0 --fqbn esp32:esp32:esp32  ./
