#ifndef API_H

#include <Arduino.h>
#include <HTTPClient.h>
#include <Preferences.h>
#include <WiFi.h>
#include <cstring>

void connectApiTest();
void putDataToServer(String pm);

#endif // !API_H
