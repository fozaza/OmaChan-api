#include "src/cmd.h"
#include <WiFi.h>
#include <Preferences.h>
#include <Arduino.h>
#include <cstring>
#include "src/i2cPm.h"
#include "src/api.h"
// #include <HardwareSerial.h>
// #include <Wire.h>
// #include <LiquidCrystal_I2C.h>

Preferences storage;
String password;
String ssid;
String hardWareName;
char ip_server[16];
char ctx[16];
unsigned int pm2_5;
boolean startPM;

void setup() {
  Serial.begin(115200);
  setupI2c(9600);
  WiFi.mode(WIFI_STA);

  Serial.println("--- Start up ---");
  Serial.println("Load data");

  storage.begin("wifi-config", false);
  ssid = storage.getString("ssid", "NONE");
  password = storage.getString("pass", "NONE");
  hardWareName = storage.getString("hardwarename", "NONE");
  // ip_server = storage.getString("server", "NONE");
  strcpy(ip_server, storage.getString("server", "NONE").c_str());
  storage.end();

  Serial.println("--- System Started ---");
  Serial.printf("\t- Last WIFI NAME >>> %s\n", ssid.c_str());
  Serial.printf("\t- Last Password >>> %s\n", password.c_str());
  Serial.printf("\t- Last Ip server >>> %s\n", ip_server);
  Serial.printf("\t- Last HardWareName >>> %s\n", hardWareName.c_str());
}

void loop() {
  input_cmd();

  static unsigned long lastout = 0;
  if (millis() - lastout > 2000) {
    if (startPM) {
      readpm();
      putDataToServer(String(pm2_5));
    }
    lastout = millis();
  }
}
