#include <config.h>
#include <stdio.h>

int main (void) {
  puts("Hello World!");
  puts ("This is " PACKAGE_STRING ".");
  puts ("Some var: " SOMEVAR ".");
  return 0;
}
