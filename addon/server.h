#pragma once

#include <arpa/inet.h>
#include <stdio.h>
#include <string.h>
#include <sys/socket.h>
#include <unistd.h>

#include "ssl.h"

int Accept(int socketDescriptor)
{
  struct sockaddr_in addr;
  socklen_t len = sizeof(addr);
  int conn;
  SSL_CTX *ctx;
  SSL *ssl;
  int socketConn;

  /* accept connection as usual */
  conn = accept(socketDescriptor, (struct sockaddr *)&addr, &len);
  //cout << "Connection: " << inet_ntoa(addr.sin_addr) << ":"
  //     << ntohs(addr.sin_port) << endl;
  printf("Connection: %s:%d\n", inet_ntoa(addr.sin_addr), ntohs(addr.sin_port));

  ctx = newCtx();
  /* get new SSL state with context */
  ssl = SSL_new(ctx);

  /* set connection socket to SSL state */
  SSL_set_fd(ssl, conn);
  /* service connection */
  decode(ssl);

  //finalizing:
  socketConn = SSL_get_fd(ssl); /* get socket connection */
  SSL_free(ssl);                /* release SSL state */
  close(socketConn);            /* close connection */

  SSL_CTX_free(ctx);
}

// Create the SSL socket and intialize the socket address structure
int ListenAndServe(const int port)
{
  struct sockaddr_in addr;
  int sd;

  // bzero in string.h
  bzero(&addr, sizeof(addr));
  addr.sin_family = AF_INET;
  addr.sin_port = htons(port);
  addr.sin_addr.s_addr = INADDR_ANY;

  sd = socket(PF_INET, SOCK_STREAM, 0);
  if (bind(sd, (struct sockaddr *)&addr, sizeof(addr)) != 0)
  {
    //perror("can't bind port");
    return -2;
  }

  if (listen(sd, 10) != 0)
  {
    //perror("Can't configure listening port");
    return -3;
  }

  return sd;
}

void Shutdown(int socketDescriptor)
{
  close(socketDescriptor);
}