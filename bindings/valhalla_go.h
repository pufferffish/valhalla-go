#ifdef __cplusplus
extern "C" {
#endif

typedef void* Actor;
Actor actor_init(const char*, char *);

const char * actor_route(Actor, const char *, char *);

const char * actor_locate(Actor, const char *, char *);

const char * actor_optimized_route(Actor, const char *, char *);

const char * actor_matrix(Actor, const char *, char *);

const char * actor_isochrone(Actor, const char *, char *);

const char * actor_trace_route(Actor, const char *, char *);

const char * actor_trace_attributes(Actor, const char *, char *);

const char * actor_height(Actor, const char *, char *);

const char * actor_transit_available(Actor, const char *, char *);

const char * actor_expansion(Actor, const char *, char *);

const char * actor_centroid(Actor, const char *, char *);

const char * actor_status(Actor, const char *, char *);


#ifdef __cplusplus
}
#endif
