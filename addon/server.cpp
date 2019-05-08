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

const int FAIL = -1;

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

void decode(SSL *ssl) /* Serve the connection -- threadable */
{
  const string response = "hello world";
  const string EXPECT_TOKEN = "I'm sammy";

  if (SSL_accept(ssl) == FAIL) /* do SSL-protocol accept */
  {
    ERR_print_errors_fp(stderr);
    return;
  }

  /* get any certificates */
  showCert(ssl);

  char req[1024] = {0};
  auto ell = SSL_read(ssl, req, sizeof(req)); /* get request */
  req[ell] = '\0';

  cout << "request msg: " << req << endl;

  if (ell > 0)
  {
    string reply = (strcmp(EXPECT_TOKEN.c_str(), req) == 0 ? response : "invalid msg");

    SSL_write(ssl, reply.c_str(), reply.size());
  }
  else
  {
    ERR_print_errors_fp(stderr);
  }
}

// new world

int Accept(const int server, const void *cert)
{
  auto ctx = (SSL_CTX *)(cert);

  //while (1)
  //{
  struct sockaddr_in addr;
  socklen_t len = sizeof(addr);

  /* accept connection as usual */
  auto conn = accept(server, (struct sockaddr *)&addr, &len);
  cout << "Connection: " << inet_ntoa(addr.sin_addr) << ":"
       << ntohs(addr.sin_port) << endl;

  auto delSSL = [](SSL *ssl) {
    auto sd = SSL_get_fd(ssl); /* get socket connection */
    SSL_free(ssl);             /* release SSL state */
    close(sd);                 /* close connection */
  };
  /* get new SSL state with context */
  //unique_ptr<SSL, decltype(delSSL)> ssl(SSL_new(ctx.get()), delSSL);
  unique_ptr<SSL, decltype(delSSL)> ssl(SSL_new(ctx), delSSL);
  /* set connection socket to SSL state */
  SSL_set_fd(ssl.get(), conn);
  /* service connection */
  decode(ssl.get());

  return 0;
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