#include <stdio.h>
#include "wrapper.h"
#include "hello.h"

void *Create() {
  return call_Person_Create();
}

void Destroy(void *p) {
  call_Person_Destroy(p);
}

const char *GetName(void *p) {
  return call_Person_GetName(p);
}

int GetAge(void *p) {
  return call_Person_GetAge(p);
}
