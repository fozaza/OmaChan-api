#include "i2cPm.h"

HardwareSerial pmsSerial(2);
LiquidCrystal_I2C lcd(0x27, 16, 2);
extern unsigned int pm2_5;

void setupI2c(uint32_t bandBit) {
  pmsSerial.begin(bandBit, SERIAL_8N1, 16, 17);
  lcd.init();
  lcd.backlight();
  int loadingValues[] = {15, 30, 55, 70, 90, 100};

  for (int i = 0; i < 6; i++) {
    lcd.setCursor(6, 0);
    lcd.print("ADCS");

    String percentText = String(loadingValues[i]) + "%";
    int startCol = (16 - percentText.length()) / 2;

    lcd.setCursor(startCol, 1);
    lcd.print(percentText);
    delay(550);

    lcd.setCursor(6, 0);
    lcd.print("    ");
    lcd.setCursor(startCol, 1);
    for (int s = 0; s < percentText.length(); s++)
      lcd.print(" ");
    delay(250);
  }

  lcd.clear();
  lcd.setCursor(6, 0);
  lcd.print("ADCS");
  return;
}

void updateScreen() {
  String line2 = "PM2.5: " + String(pm2_5) + " ug";

  int startCursor = (16 - line2.length()) / 2;

  if (startCursor < 0)
    startCursor = 0;

  lcd.setCursor(0, 1);
  lcd.print("                ");

  lcd.setCursor(startCursor, 1);
  lcd.print(line2);

  Serial.print("PM2.5: ");
  Serial.println(pm2_5);
}

void readpm() {
  if (pmsSerial.available() >= 32) {
    if (pmsSerial.read() == 0x42) {
      if (pmsSerial.read() == 0x4D) {

        uint8_t dataBuffer[30];
        pmsSerial.readBytes(dataBuffer, 30);
        unsigned int new_pm2_5 = (dataBuffer[4] << 8) + dataBuffer[5];

        pm2_5 = new_pm2_5;
        updateScreen();
      }
    }
  }
}
