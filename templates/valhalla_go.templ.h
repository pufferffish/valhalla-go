#ifdef __cplusplus
extern "C" {
#endif

typedef void* Actor;
Actor actor_init(const char*);
{{ range .Functions }}
const char * actor_{{.}}(Actor, const char *, char *);
{{ end }}

#ifdef __cplusplus
}
#endif
