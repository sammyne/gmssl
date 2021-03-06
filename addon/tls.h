#pragma once

#ifdef __cplusplus
extern "C"
{
#endif

#include "types.h"

  void destroyCert(const void *cert);

  // initTLS initializes the SSL library
  int initTLS();

  //Cert *loadX509KeyPair(const char *certFile, const char *keyFile);
  Response loadX509KeyPair(const char *certFile, const char *keyFile);

#ifdef __cplusplus
}
#endif