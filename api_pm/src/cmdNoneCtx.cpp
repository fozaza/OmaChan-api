#include "cmdNoneCtx.h";
#include "api.h";

extern Preferences storage;
extern String ssid;
extern String password;
extern String hardWareName;
extern char ip_server[16];
extern boolean startPM;
int maxWifiConnect = 30;

void wifi_connect() {
  Serial.printf("connect wifi %s\n", ssid);
  ssid.trim();
  password.trim();
  WiFi.mode(WIFI_STA);
  WiFi.begin(ssid, password);

  int i = 0;
  while (WiFi.status() != WL_CONNECTED) {
    if (i == maxWifiConnect)
      break;
    Serial.printf("\t\t- Filed to connect wifi %d\n", i + 1);
    i++;
    delay(1000);
  }

  if (WiFi.status() == WL_CONNECTED) {
    Serial.printf("\tConnected Success\n\t\t- IP: ");
    Serial.println(WiFi.localIP());
  } else {
    Serial.println("\tConnected Filed");
  }
  return;
}

void cmdNoneCtx(String &typeCmd) {
  if (typeCmd == "-o") {
    Serial.printf("config\n\t- ssid >>> %s\n\t- password >>> %s\n\t- "
                  "url_server >>> %s\n\t- HardWareName >>> %s\n",
                  ssid, password, ip_server, hardWareName);
    return;
  }

  if (typeCmd == "-u") {
    connectApiTest();
    return;
  }

  if (typeCmd == "-c") {
    wifi_connect();
    return;
  }

  if (typeCmd == "-r") {
    if (startPM) {
      Serial.println("HardWareName Status >>> Enbale");
      startPM = 0;
    } else {
      Serial.println("HardWareName Status >>> Disbale");
      startPM = 1;
    }
    return;
  }
  // if (typeCmd == "-l") {
  //   putDataToServer();
  // }

  Serial.printf("Command not suppoet %s", typeCmd);
  return;
}
