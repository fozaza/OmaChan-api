#include "cmdCtx.h"
#include "api.h"

extern Preferences storage;
extern String ssid;
extern String password;
extern String hardWareName;
extern char ip_server[16];
extern char ctx[16];

void setdata(String &old, String ssh) {
  // Serial.println(ctx);
  storage.begin("wifi-config", false);
  Serial.printf("Set new %s %s >>> %s\n\n", ssh, old, ctx);
  old = ctx;
  storage.putString(ssh.c_str(), old);
  storage.end();
  return;
};

void cmdCtx(String &typeCmd) {
  storage.begin("wifi-config", false);
  if (typeCmd == "-s")
    setdata(ssid, "ssid");

  if (typeCmd == "-p")
    setdata(password, "pass");

  if (typeCmd == "-h") {
    setdata(hardWareName, "hardwarename");
  }

  if (typeCmd == "-l")
    putDataToServer(ctx);

  if (typeCmd == "-v") {
    storage.begin("wifi-config", false);
    Serial.printf("Set new %s %s >>> %s\n\n", "password", ip_server, ctx);
    strlcpy(ip_server, ctx, 16);
    Serial.println(ip_server);
    storage.putString("server", ip_server);
    storage.end();
  }
  return;
}
