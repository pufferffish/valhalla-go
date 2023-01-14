{ pkgs ? import <nixpkgs> {} }:
pkgs.callPackage ./bindings {}
