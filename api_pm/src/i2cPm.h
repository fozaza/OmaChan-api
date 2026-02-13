#ifndef I2CPM_H

#include <Arduino.h>
#include <HardwareSerial.h>
#include <LiquidCrystal_I2C.h>
#include <Wire.h>
#include <cstdint>

void setupI2c(uint32_t bandBit);
void readpm();

#endif // !I2CPM_H
