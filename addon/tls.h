#pragma once

#ifdef __cplusplus
extern "C"
{
#endif

  typedef struct
  {
    void *value;
    int error;
  } Response;

  struct _cert;
  typedef struct _cert Cert;

  //void destroyCert(const Cert *cert);
  void destroyCert(const void *cert);

  // initTLS initializes the SSL library
  int initTLS();

  //Cert *loadX509KeyPair(const char *certFile, const char *keyFile);
  Response loadX509KeyPair(const char *certFile, const char *keyFile);

#ifdef __cplusplus
}
#endif