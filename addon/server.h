#pragma once

#ifdef __cplusplus
extern "C"
{
#endif

  int Close(const int server);
  int Listen(const int port);

  //int ListenAndServe();
  int ListenAndServe(const int server);

#ifdef __cplusplus
}
#endif