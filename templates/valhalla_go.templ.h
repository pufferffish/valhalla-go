#ifdef __cplusplus
extern "C" {
#endif

typedef void* Actor;
Actor actor_init_from_file(const char*, char *);
Actor actor_init_from_config(const char*, char *);
{{ range $k, $v := .Functions }}
const char * actor_{{$k}}(Actor, const char *, char *);
{{ end }}

#ifdef __cplusplus
}
#endif
