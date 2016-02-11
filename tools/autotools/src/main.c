#include <config.h>
#include <stdio.h>

int main (void) {
  puts("Hello World!");
  puts ("This is " PACKAGE_STRING ".");
  puts ("Some var: " SOMEVAR ".");

#if HAVE_SDL2_SDL_H
  puts ("Has SDL!");
#endif

  return 0;
}
