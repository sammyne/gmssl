#include "types.h"

#include <iostream>

#include <unistd.h>

#include <openssl/ssl.h>

void CloseSSL(const void *ssl)
{
  auto _ssl = (SSL *)(ssl);
  if (!_ssl)
  {
    return;
  }

  auto sd = SSL_get_fd(_ssl); /* get socket connection */
  SSL_free(_ssl);             /* release SSL state */
  close(sd);                  /* close connection */
}

void Disconnect(const void *conn)
{
  auto c = (Conn *)(conn);
  if (!c)
  {
    return;
  }

  if (auto ssl = (SSL *)(c->ssl); ssl)
  {
    auto sd = SSL_get_fd(ssl); /* get socket connection */
    SSL_free(ssl);             /* release SSL state */
    close(sd);                 /* close connection */
  }
}

const void *ReleaseSSL(const Conn *conn)
{
  if (!conn)
  {
    return nullptr;
  }

  auto ssl = conn->ssl;
  delete conn;

  return ssl;
}

void Hello(const void *conn)
{
  auto c = (const Conn *)(conn);
  if (!c)
  {
    std::cout << "c is nil" << std::endl;
    return;
  }

  auto ip = c->remoteIP;

  std::cout << (ip & 0xff) << "." << ((ip >> 8) & 0xff) << "."
            << ((ip >> 16) & 0xff) << "." << ((ip >> 24) & 0xff)

            << ":" << c->remotePort << std::endl;
  //std::cout << c->remoteIP << ":" << c->remotePort << std::endl;
}