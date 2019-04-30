#pragma once

#include <stdio.h>
#include <string.h>

#include <openssl/err.h>
#include <openssl/gmtls.h>
#include <openssl/ssl.h>

const int FAIL = -1;

int initSSL()
{
  return SSL_library_init();
}

int loadCert(SSL_CTX *ctx, const char *cert)
{
  return SSL_CTX_use_certificate_file(ctx, cert, SSL_FILETYPE_PEM);
}

int loadKey(SSL_CTX *ctx, const char *prv)
{
  int err;
  /* set the private key from KeyFile (may be the same as CertFile) */
  err = SSL_CTX_use_PrivateKey_file(ctx, prv, SSL_FILETYPE_PEM);
  if (err <= 0)
  {
    return err;
  }

  /* verify private key */
  return SSL_CTX_check_private_key(ctx);
}

SSL_CTX *newCtx()
{
  SSL_CTX *ctx;

  OpenSSL_add_all_algorithms(); /* load & register all cryptos, etc. */
  SSL_load_error_strings();     /* load all error messages */

  ctx = SSL_CTX_new(TLS_server_method());
  if (!ctx)
  {
    return ctx;
  }

  if (!SSL_CTX_set_cipher_list(ctx, GMTLS_TXT_SM2DHE_WITH_SMS4_SM3))
  {
    return NULL;
  }

  return ctx;
}

void showCert(const SSL *ssl)
{
  X509 *cert;
  char *line;

  /* Get certificates (if available) */
  cert = SSL_get_peer_certificate(ssl);
  if (!cert)
  {
    printf("no cert\n");
    return;
  }

  printf("client cert\n");

  line = X509_NAME_oneline(X509_get_subject_name(cert), 0, 0);
  printf("Subject: %s\n", line);
  free(line);

  line = X509_NAME_oneline(X509_get_issuer_name(cert), 0, 0);
  //cout << "Issuer: " << line << endl;
  printf("Issuer: %s\n", line);
  free(line);

  //finalizing:
  X509_free(cert);
}

/**
 * following are placed due to dependency ordering
*/
void decode(SSL *ssl) /* Serve the connection -- threadable */
{
  const char *EXPECT_TOKEN = "I'm sammy";
  char *response = "hello world";
  char *reply;
  char req[2014] = {0};
  int ell;

  if (SSL_accept(ssl) == FAIL) /* do SSL-protocol accept */
  {
    ERR_print_errors_fp(stderr);
    return;
  }

  /* get any certificates */
  showCert(ssl);

  ell = SSL_read(ssl, req, sizeof(req)); /* get request */
  req[ell] = '\0';

  printf("request msg: %s\n", req);

  if (ell <= 0)
  {
    ERR_print_errors_fp(stderr);
    return;
  }

  reply = (strcmp(EXPECT_TOKEN, req) == 0 ? response : "invalid msg");

  SSL_write(ssl, reply, strlen(reply));
}