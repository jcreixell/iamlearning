AC_INIT([helloautotools], [1.0], [jcreixell@gmail.com])
AC_CONFIG_AUX_DIR([build-aux])
AM_INIT_AUTOMAKE([foreign -Wall -Werror])
AC_PROG_CC
AC_PROG_CXX
AC_CHECK_PROGS([TAR], [tar gtar], [:])
if test "$TAR" = :; then
  AC_MSG_ERROR([This package needs tar.])
fi
AC_DEFINE([SOMEVAR], ["jorge"], [Set some var, just for fun])
AC_SUBST([MAKEFILEVAR], [something])
AC_CHECK_HEADERS([SDL2/SDL.h])
AC_CONFIG_HEADERS([config.h])
AC_CONFIG_FILES([Makefile src/Makefile])
AC_OUTPUT
