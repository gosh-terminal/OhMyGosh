function install_me() {
  if [ -f $HOME/.ohmygosh ]; then
  		mkdir $HOME/.ohmygosh
  fi
  echo 'export PATH="$PATH:$HOME/.ohmygosh"' >> ~/.bashrc;
  go get -v -t -d ./...;
  go build -o ohmygosh src/*.go;
  mv ohmygosh "$HOME/.ohmygosh/ohmygosh"
  source ~/.bashrc;
}
install_me;