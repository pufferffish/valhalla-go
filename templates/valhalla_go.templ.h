#ifdef __cplusplus
extern "C" {
#endif

typedef void* Actor;
Actor actor_init(const char*, char *);
{{ range $k, $v := .Functions }}
const char * actor_{{$k}}(Actor, const char *, char *);
{{ end }}

#ifdef __cplusplus
}
#endif
