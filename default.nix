{ stdenv, makeWrapper, lib, fetchFromGitHub, glibcLocales,
coreutils, curl, jshon, bc, gnused, perl, datamash, git }:

stdenv.mkDerivation rec {
  pname = "setzer";
  version = "0.8.0";
  src = ./.;

  nativeBuildInputs = [makeWrapper];
  buildPhase = "true";
  makeFlags = ["prefix=$(out)"];
  postInstall = let path = lib.makeBinPath [
    coreutils curl jshon bc gnused perl datamash git
  ]; in ''
    wrapProgram "$out/bin/setzer" --set PATH "${path}" \
      ${if glibcLocales != null then
        "--set LOCALE_ARCHIVE \"${glibcLocales}\"/lib/locale/locale-archive"
        else ""}
  '';

  meta = with lib; {
    description = "Asset Price feed tool for on-chain oracles";
    homepage = https://github.com/chronicleprotocol/setzer;
    license = licenses.gpl3;
    inherit version;
  };
}
