{
  description = "postgres devshell";

  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixos-25.11";

  outputs = { self, nixpkgs }:
  let
    system = "x86_64-linux";
    pkgs = import nixpkgs { inherit system; };
  in {
    devShells.${system}.default = pkgs.mkShell {
      packages = with pkgs; [
        postgresql_18
        pgcli
      ];

      shellHook = ''
        export PROJECT_ROOT=$(git rev-parse --show-toplevel 2>/dev/null || pwd)
        export PGDATA="$PROJECT_ROOT/package/db/.pgdata"
        export PGHOST="$PROJECT_ROOT/package/db/.pgsocket"
        export PGPORT=11011

        export APP_ADDRESS=":8080"
        export DATABASE_URL="host=$PGHOST port=$PGPORT dbname=booking"
        export JWT_SECRET="veryveryveryveryveryveryveryveryveryverylongsecret"

        mkdir -p "$PGDATA" "$PGHOST"

        if [ ! -f "$PGDATA/PG_VERSION" ]; then
          initdb --auth=trust -D "$PGDATA"
        fi

        pg_ctl stop -D "$PGDATA" -m fast
        pg_ctl start -D "$PGDATA" -l "$PROJECT_ROOT/package/db/log.txt" -o "-k $PGHOST -p $PGPORT"
        echo "postgres started at $PGHOST:$PGPORT"

        psql -h $PGHOST -p $PGPORT -d postgres -f $PROJECT_ROOT/package/db/src/default.sql
      '';
    };
  };
}
