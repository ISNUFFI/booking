{
  flake,
  self,
  inputs,
  system,
  ...
}: let
  name = "general";
in self.lib.nixpkgs-lib.nixosSystem {
  inherit (self.legacyPackages."${system}") pkgs;
  modules = [
    { networking.hostName = name; }
    (import ./${name}.nix { inherit flake self inputs; })
  ];
}
