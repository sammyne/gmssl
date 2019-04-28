include(ExternalProject)

### GmSSL
SET(GMSSL_PREFIX ${CMAKE_CURRENT_SOURCE_DIR}/3rd_party/gmssl) 
# INSTALL_DIR ${CMAKE_CURRENT_SOURCE_DIR}/3rd_party/gmssl  
ExternalProject_Add(GmSSL 
        PREFIX ${GMSSL_PREFIX} 
        GIT_PROGRESS 1 
        GIT_REPOSITORY https://github.com/guanzhi/GmSSL.git 
        GIT_TAG 4a20b5f54c0a313ce998d8ecc5dd8f34c5c4c1b4 
        CONFIGURE_COMMAND ./config --prefix=${GMSSL_PREFIX} 
        no-weak-ssl-ciphers enable-ec_nistp_64_gcc_128 
        BUILD_IN_SOURCE 1)

ExternalProject_Get_Property(GmSSL PREFIX)
#MESSAGE("--- ${PREFIX}")

# set global env to referenced by others
set(GMSSL_INCLUDE_DIRECTORIES ${PREFIX}/include)
set(GMSSL_LINK_DIRECTORIES ${PREFIX}/lib)

# create the ${PREFIX}/include directory
file(MAKE_DIRECTORY ${PREFIX}/include)