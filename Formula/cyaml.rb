# This file was generated by GoReleaser. DO NOT EDIT.
class Cyaml < Formula
  desc ""
  homepage ""
  version "0.1.1"
  bottle :unneeded

  if OS.mac?
    url "https://github.com/jamhed/cyaml/releases/download/v0.1.1/cyaml_0.1.1_darwin_amd64.tar.gz"
    sha256 "ea698b7e303522174ce21229fa022fc306d0c2ff5d73fd741f40d1b46184bf74"
  elsif OS.linux?
    if Hardware::CPU.intel?
      url "https://github.com/jamhed/cyaml/releases/download/v0.1.1/cyaml_0.1.1_linux_amd64.tar.gz"
      sha256 "adf28071bfc9a5fb4cc55baad59656495999ddededd6bbc31114499a717b0907"
    end
  end

  def install
    bin.install "cyaml"
  end
end