#include "api.h"
#include <cstdio>

extern char ip_server[16];
extern String hardWareName;

// Serial.println("connect to server");
// String httpserver = "http:/" + ip_server + "/";
void connectApiTest() {
  Serial.println("Start to test connect server");
  if (WiFi.status() != WL_CONNECTED) {
    Serial.println("\t-Wifi not connect");
    return;
  }
  HTTPClient http;
  Serial.println(ip_server);
  // String httpserver = "http://" + String(ip_server) + ":3000" + "/";
  char url[50];
  snprintf(url, sizeof(url), "http://%s:3000/", ip_server);
  Serial.printf("url : %s\n", url);
  http.begin(url);
  // http.begin("http://192.168.1.176:3000/");
  int httpCode = http.GET();

  if (httpCode > 0) {
    String payload = http.getString();
    Serial.print("\tResponse:\n\t\t- ");
    Serial.println(payload);
  } else {
    Serial.printf("\tError: %s\n", http.errorToString(httpCode).c_str());
  }
  http.end();
  return;
}

void putDataToServer(String pm) {
  Serial.println("Put data to server");
  if (WiFi.status() != WL_CONNECTED) {
    Serial.println("\t-Wifi not connect");
    return;
  }

  HTTPClient http;
  // char url[] = "http://192.168.1.176:3000/hwdUp?Name=Hardware1";
  char url[100];
  snprintf(url, sizeof(url), "http://%s:3000/hwdUp?Name=%s&Pm=%s&Batter=0",
           ip_server, hardWareName, pm);
  // String jsonData = "{\"pm\":\"100\",\"batter\":\"0\" }";
  // String jsonData = "{\"pm\":" + pm + "\"batter\":\"0\"}";
  // char jsonData[35];
  // snprintf(jsonData, sizeof(jsonData), "{\"pm\":
  // \"%s\",\"batter\":\"NONE\"}",
  //          pm);
  // Serial.println(pm);
  // Serial.println(jsonData);
  // Serial.println(url);
  http.begin(url);
  // int httpResponseCode = http.POST(String(jsonData));
  String jsonData;
  int httpResponseCode = http.POST(jsonData);

  if (httpResponseCode > 0) {
    Serial.print("\t- Response Code: ");
    Serial.println(httpResponseCode);
    String response = http.getString();
    Serial.println("\t\t- Response: " + response);
  } else {
    Serial.printf("\t- Error: %s\n",
                  http.errorToString(httpResponseCode).c_str());
  }
  http.end();
}
