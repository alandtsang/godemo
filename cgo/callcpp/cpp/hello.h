#ifndef HELLO_H_
#define HELLO_H_

#ifdef __cplusplus
extern "C" {
#endif

extern void *Create();
extern void Destroy(void *p);
extern const char *GetName(void *p);
extern int GetAge(void *p);

#ifdef __cplusplus
}
#endif

#endif // HELLO_H_
