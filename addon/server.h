#pragma once

#ifdef __cplusplus
extern "C"
{
#endif

#include "types.h"

  Response Accept(const int server, const void *cert);
  int Close(const int server);
  int Listen(const int port);
  int Read(const void *ssl, char *buf, const int bufLen);
  int Write(const void *ssl, const char *msg, const int msgLen);

#ifdef __cplusplus
}
#endif