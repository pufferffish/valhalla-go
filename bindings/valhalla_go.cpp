#include <valhalla/tyr/actor.h>
#include <valhalla/baldr/rapidjson_utils.h>
#include <valhalla/midgard/logging.h>
#include <valhalla/midgard/util.h>

#include <boost/make_shared.hpp>
#include <boost/noncopyable.hpp>
#include <boost/optional.hpp>
#include <boost/property_tree/ptree.hpp>

#include "valhalla_go.h"

const boost::property_tree::ptree configure(const std::string& config) {
  boost::property_tree::ptree pt;
  try {
    // parse the config and configure logging
    rapidjson::read_json(config, pt);

    boost::optional<boost::property_tree::ptree&> logging_subtree =
        pt.get_child_optional("mjolnir.logging");
    if (logging_subtree) {
      auto logging_config = valhalla::midgard::ToMap<const boost::property_tree::ptree&,
                                                     std::unordered_map<std::string, std::string>>(
          logging_subtree.get());
      valhalla::midgard::logging::Configure(logging_config);
    }
  } catch (...) { throw std::runtime_error("Failed to load config from: " + config); }

  return pt;
}

Actor actor_init(const char* config) {
  auto actor = new valhalla::tyr::actor_t(configure(config), true);
  return (Actor) actor;
}


const char * actor_route(Actor actor, const char * req, char * is_error) {
  try {
    std::string resp = ((valhalla::tyr::actor_t*) actor)->route(req);
    char *cstr = (char *) malloc(resp.length() + 1);
    strcpy(cstr, resp.c_str());
    *is_error = 0;
    return cstr;
  } catch (std::exception& ex) {
    char *cstr = (char *) malloc(strlen(ex.what()) + 1);
    strcpy(cstr, ex.what());
    *is_error = 1;
    return cstr;
  }
}

const char * actor_locate(Actor actor, const char * req, char * is_error) {
  try {
    std::string resp = ((valhalla::tyr::actor_t*) actor)->locate(req);
    char *cstr = (char *) malloc(resp.length() + 1);
    strcpy(cstr, resp.c_str());
    *is_error = 0;
    return cstr;
  } catch (std::exception& ex) {
    char *cstr = (char *) malloc(strlen(ex.what()) + 1);
    strcpy(cstr, ex.what());
    *is_error = 1;
    return cstr;
  }
}

const char * actor_optimized_route(Actor actor, const char * req, char * is_error) {
  try {
    std::string resp = ((valhalla::tyr::actor_t*) actor)->optimized_route(req);
    char *cstr = (char *) malloc(resp.length() + 1);
    strcpy(cstr, resp.c_str());
    *is_error = 0;
    return cstr;
  } catch (std::exception& ex) {
    char *cstr = (char *) malloc(strlen(ex.what()) + 1);
    strcpy(cstr, ex.what());
    *is_error = 1;
    return cstr;
  }
}

const char * actor_matrix(Actor actor, const char * req, char * is_error) {
  try {
    std::string resp = ((valhalla::tyr::actor_t*) actor)->matrix(req);
    char *cstr = (char *) malloc(resp.length() + 1);
    strcpy(cstr, resp.c_str());
    *is_error = 0;
    return cstr;
  } catch (std::exception& ex) {
    char *cstr = (char *) malloc(strlen(ex.what()) + 1);
    strcpy(cstr, ex.what());
    *is_error = 1;
    return cstr;
  }
}

const char * actor_isochrone(Actor actor, const char * req, char * is_error) {
  try {
    std::string resp = ((valhalla::tyr::actor_t*) actor)->isochrone(req);
    char *cstr = (char *) malloc(resp.length() + 1);
    strcpy(cstr, resp.c_str());
    *is_error = 0;
    return cstr;
  } catch (std::exception& ex) {
    char *cstr = (char *) malloc(strlen(ex.what()) + 1);
    strcpy(cstr, ex.what());
    *is_error = 1;
    return cstr;
  }
}

const char * actor_trace_route(Actor actor, const char * req, char * is_error) {
  try {
    std::string resp = ((valhalla::tyr::actor_t*) actor)->trace_route(req);
    char *cstr = (char *) malloc(resp.length() + 1);
    strcpy(cstr, resp.c_str());
    *is_error = 0;
    return cstr;
  } catch (std::exception& ex) {
    char *cstr = (char *) malloc(strlen(ex.what()) + 1);
    strcpy(cstr, ex.what());
    *is_error = 1;
    return cstr;
  }
}

const char * actor_trace_attributes(Actor actor, const char * req, char * is_error) {
  try {
    std::string resp = ((valhalla::tyr::actor_t*) actor)->trace_attributes(req);
    char *cstr = (char *) malloc(resp.length() + 1);
    strcpy(cstr, resp.c_str());
    *is_error = 0;
    return cstr;
  } catch (std::exception& ex) {
    char *cstr = (char *) malloc(strlen(ex.what()) + 1);
    strcpy(cstr, ex.what());
    *is_error = 1;
    return cstr;
  }
}

const char * actor_height(Actor actor, const char * req, char * is_error) {
  try {
    std::string resp = ((valhalla::tyr::actor_t*) actor)->height(req);
    char *cstr = (char *) malloc(resp.length() + 1);
    strcpy(cstr, resp.c_str());
    *is_error = 0;
    return cstr;
  } catch (std::exception& ex) {
    char *cstr = (char *) malloc(strlen(ex.what()) + 1);
    strcpy(cstr, ex.what());
    *is_error = 1;
    return cstr;
  }
}

const char * actor_transit_available(Actor actor, const char * req, char * is_error) {
  try {
    std::string resp = ((valhalla::tyr::actor_t*) actor)->transit_available(req);
    char *cstr = (char *) malloc(resp.length() + 1);
    strcpy(cstr, resp.c_str());
    *is_error = 0;
    return cstr;
  } catch (std::exception& ex) {
    char *cstr = (char *) malloc(strlen(ex.what()) + 1);
    strcpy(cstr, ex.what());
    *is_error = 1;
    return cstr;
  }
}

const char * actor_expansion(Actor actor, const char * req, char * is_error) {
  try {
    std::string resp = ((valhalla::tyr::actor_t*) actor)->expansion(req);
    char *cstr = (char *) malloc(resp.length() + 1);
    strcpy(cstr, resp.c_str());
    *is_error = 0;
    return cstr;
  } catch (std::exception& ex) {
    char *cstr = (char *) malloc(strlen(ex.what()) + 1);
    strcpy(cstr, ex.what());
    *is_error = 1;
    return cstr;
  }
}

const char * actor_centroid(Actor actor, const char * req, char * is_error) {
  try {
    std::string resp = ((valhalla::tyr::actor_t*) actor)->centroid(req);
    char *cstr = (char *) malloc(resp.length() + 1);
    strcpy(cstr, resp.c_str());
    *is_error = 0;
    return cstr;
  } catch (std::exception& ex) {
    char *cstr = (char *) malloc(strlen(ex.what()) + 1);
    strcpy(cstr, ex.what());
    *is_error = 1;
    return cstr;
  }
}

const char * actor_status(Actor actor, const char * req, char * is_error) {
  try {
    std::string resp = ((valhalla::tyr::actor_t*) actor)->status(req);
    char *cstr = (char *) malloc(resp.length() + 1);
    strcpy(cstr, resp.c_str());
    *is_error = 0;
    return cstr;
  } catch (std::exception& ex) {
    char *cstr = (char *) malloc(strlen(ex.what()) + 1);
    strcpy(cstr, ex.what());
    *is_error = 1;
    return cstr;
  }
}
