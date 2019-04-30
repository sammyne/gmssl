cmake_minimum_required(VERSION 3.12)

project(addon)

set(ADDON_PREFIX ${CMAKE_SOURCE_DIR}/3rd_party/addon)

include_directories(${GMSSL_PREFIX}/include)

#add_library(cert SHARED cert.cpp)
#add_dependencies(cert GmSSL)
#install(TARGETS cert DESTINATION "${ADDON_PREFIX}/lib")

aux_source_directory(${CMAKE_CURRENT_SOURCE_DIR} ADDON_SRCS)
#message("-------- ${ADDON_SRCS}")
add_library(addon SHARED ${ADDON_SRCS})
add_dependencies(addon GmSSL)
install(TARGETS addon DESTINATION "${ADDON_PREFIX}/lib")

file(GLOB add_headers "${CMAKE_CURRENT_SOURCE_DIR}/*.h")
install(FILES ${add_headers} DESTINATION "${ADDON_PREFIX}/include")