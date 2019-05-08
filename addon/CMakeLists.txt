cmake_minimum_required(VERSION 3.12)

project(addon)

set(ADDON_PREFIX ${CMAKE_SOURCE_DIR}/3rd_party/addon)

# below is to set GMSSL_PREFIX to avoid re-compilation GmSSL every time
if(NOT GMSSL_PREFIX)
  set(GMSSL_PREFIX ${CMAKE_SOURCE_DIR}/3rd_party/gmssl) 
#  message(${GMSSL_PREFIX})
endif(NOT GMSSL_PREFIX)

include_directories(${GMSSL_PREFIX}/include)

aux_source_directory(${CMAKE_CURRENT_SOURCE_DIR} ADDON_SRCS)
#message("-------- ${ADDON_SRCS}")
add_library(addon SHARED ${ADDON_SRCS})

if(Lib3rd)
  add_dependencies(addon ${Lib3rd})
endif(Lib3rd)

install(TARGETS addon DESTINATION "${ADDON_PREFIX}/lib")

file(GLOB add_headers "${CMAKE_CURRENT_SOURCE_DIR}/*.h")
install(FILES ${add_headers} DESTINATION "${ADDON_PREFIX}/include")