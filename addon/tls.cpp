#include "tls.h"

#include <memory>

#include <openssl/gmtls.h>
#include <openssl/ssl.h>

#include <openssl/err.h>

using namespace std;

/** dependencies */
SSL_CTX *newCtx()
{
  auto method = TLS_server_method(); /* create new server-method instance */
  auto ctx = SSL_CTX_new(method);    /* create new context from method */
  if (ctx == NULL)
  {
    ERR_print_errors_fp(stderr);
    return nullptr;
  }

  //if (!SSL_CTX_set_cipher_list(ctx, GMTLS_TXT_ECDHE_SM2_WITH_SMS4_SM3))
  if (!SSL_CTX_set_cipher_list(ctx, GMTLS_TXT_SM2DHE_WITH_SMS4_SM3))
  {
    ERR_print_errors_fp(stderr);
    return nullptr;
  }

  return ctx;
}

/** end dependencies */

/** API implementations */
//void destroyCert(const Cert *cert)
void destroyCert(const void *cert)
{
  SSL_CTX_free((SSL_CTX *)(cert));
  /*
  if (cert)
  {
    SSL_CTX_free((SSL_CTX *)(cert));
  }*/
}

//Cert *loadX509KeyPair(const char *certFile, const char *keyFile)
Response loadX509KeyPair(const char *certFile, const char *keyFile)
{
  //auto cert = new Cert;
  auto delCtx = [](SSL_CTX *ctx) {
    SSL_CTX_free(ctx);
  };

  std::unique_ptr<SSL_CTX, decltype(delCtx)> ctx(newCtx(), delCtx);

  /* set the local certificate from CertFile */
  if (SSL_CTX_use_certificate_file(ctx.get(), certFile, SSL_FILETYPE_PEM) <= 0)
  {
    //ERR_print_errors_fp(stderr);
    //abort();
    return {nullptr, 1};
  }

  /* set the private key from KeyFile (may be the same as CertFile) */
  if (SSL_CTX_use_PrivateKey_file(ctx.get(), keyFile, SSL_FILETYPE_PEM) <= 0)
  {
    //ERR_print_errors_fp(stderr);
    //abort();
    return {nullptr, 2};
  }
  /* verify private key */
  if (!SSL_CTX_check_private_key(ctx.get()))
  {
    //cerr << "private key does not match the public certificate" << endl;
    //abort();
    return {nullptr, 3};
  }

  return {ctx.release(), 0};
}

int initTLS()
{
  SSL_library_init();
  OpenSSL_add_all_algorithms(); /* load & register all cryptos, etc. */
  SSL_load_error_strings();     /* load all error messages */

  return 1;
}

/** end API implementations */