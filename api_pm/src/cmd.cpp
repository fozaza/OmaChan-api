#include "cmd.h"
#include "cmdCtx.h"
#include "cmdNoneCtx.h"

extern char ctx[16];
String commandNoneCtx[] = {"-o", "-c", "-u", "-d", "-r"};
String commandCtx[] = {"-s", "-p", "-v", "-l", "-h"};

void input_cmd() {
  if (!Serial.available() > 0)
    return;

  String input = Serial.readString();
  input.trim();
  if (input.length() == 0)
    return;

  // Serial.println(input);

  // map cmd
  String typeCmd;
  for (int i; i < 2; i++)
    typeCmd += input[i];

  // ctx
  String ctxString;
  // ctx = "";
  for (int i = 2; i < input.length(); i++)
    ctxString += input[i];
  ctxString.trim();
  ctxString.toCharArray(ctx, 16);
  // Serial.println(ctxString);
  // Serial.printf("ctx %s\n", ctx);

  for (int i; i < sizeof(commandNoneCtx); i++)
    if (typeCmd == commandNoneCtx[i]) {
      cmdNoneCtx(typeCmd);
      return;
    }

  for (int i; i < sizeof(commandCtx); i++)
    if (typeCmd == commandCtx[i]) {
      cmdCtx(typeCmd);
      return;
    }

  Serial.printf("not found command %s\n", typeCmd);
  return;
}
