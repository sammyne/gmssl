#include "server.h"

#include <iostream>
#include <memory>

#include <errno.h>
#include <unistd.h>
#include <malloc.h>
#include <string.h>
#include <arpa/inet.h>
#include <sys/socket.h>
#include <sys/types.h>
#include <netinet/in.h>
#include <resolv.h>
#include "openssl/ssl.h"
#include "openssl/err.h"

#include <openssl/gmtls.h>

using namespace std;

// --- old stuff
void showCert(const SSL *ssl)
{
  /* Get certificates (if available) */
  auto delCert = [](X509 *cert) {
    X509_free(cert);
  };

  //auto cert = SSL_get_peer_certificate(ssl);
  unique_ptr<X509, decltype(delCert)> cert(SSL_get_peer_certificate(ssl),
                                           delCert);
  //if (nullptr == cert)
  if (nullptr == cert.get())
  {
    cout << "no cert" << endl;
    return;
  }

  cout << "client cert:" << endl;

  auto line = X509_NAME_oneline(X509_get_subject_name(cert.get()), 0, 0);
  cout << "Subject: " << line << endl;
  free(line);

  line = X509_NAME_oneline(X509_get_issuer_name(cert.get()), 0, 0);
  cout << "Issuer: " << line << endl;
  free(line);

  //X509_free(cert);
}
// new world

//int Accept(const int server, const void *cert)
Response Accept(const int server, const void *cert)
{
  auto ctx = (SSL_CTX *)(cert);

  auto conn = new Conn;

  struct sockaddr_in addr;
  socklen_t len = sizeof(addr);

  /* accept connection as usual */
  auto socket = accept(server, (struct sockaddr *)&addr, &len);
  conn->remoteIP = addr.sin_addr.s_addr;
  conn->remotePort = ntohs(addr.sin_port);

  //auto delSSL = Disconnect;
  /* get new SSL state with context */

  //unique_ptr<SSL, decltype(delSSL)> ssl(SSL_new(ctx), delSSL);
  auto ssl = SSL_new(ctx);
  /* set connection socket to SSL state */
  //SSL_set_fd(ssl.get(), conn);
  SSL_set_fd(ssl, socket);
  /* service connection */
  //if (SSL_accept(ssl.get()) == FAIL) /* do SSL-protocol accept */
  if (-1 == SSL_accept(ssl)) /* do SSL-protocol accept */
  {
    //ERR_print_errors_fp(stderr);
    return {nullptr, 1};
  }

  conn->ssl = ssl;

  return {conn, 0};
}

int Close(const int server)
{
  return close(server); /* close server socket */
}

// Create the SSL socket and intialize the socket address structure
int Listen(const int port)
{
  struct sockaddr_in addr;

  bzero(&addr, sizeof(addr));
  addr.sin_family = AF_INET;
  addr.sin_port = htons(port);
  addr.sin_addr.s_addr = INADDR_ANY;

  auto sd = socket(PF_INET, SOCK_STREAM, 0);
  if (bind(sd, (struct sockaddr *)&addr, sizeof(addr)) != 0)
  {
    perror("can't bind port");
    abort();
  }

  if (listen(sd, 10) != 0)
  {
    perror("Can't configure listening port");
    abort();
  }

  return sd;
}

int Read(const void *ssl, char *buf, const int bufLen)
{
  auto _ssl = (SSL *)(ssl);
  if (!_ssl)
  {
    return -2;
  }

  //decode(_ssl);
  //return 0;

  // read out request message
  return SSL_read(_ssl, buf, bufLen); /* get request */
}

int Write(const void *ssl, const char *msg, const int msgLen)
{
  auto _ssl = (SSL *)(ssl);
  if (!_ssl)
  {
    return -2;
  }

  return SSL_write(_ssl, msg, msgLen);
}