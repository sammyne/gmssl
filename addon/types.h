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

  typedef struct
  {
    /* data */
    unsigned int remoteIP;
    short unsigned int remotePort;
    const void *ssl;
  } Conn;

  void CloseSSL(const void *ssl);
  void Disconnect(const void *conn);
  const void *ReleaseSSL(const Conn *conn);

  // TO BE REMOVE
  void Hello(const void *conn);

#ifdef __cplusplus
}
#endif