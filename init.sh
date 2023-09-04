#!/bin/bash

apt-get update && apt-get install -y tmux vim zsh autojump wget ssh less iputils-ping tig graphviz tree

rm -r ~/.oh-my-zsh
sh -c "$(curl -fsSL https://raw.github.com/ohmyzsh/ohmyzsh/master/tools/install.sh) --unattended"
chsh -s $(which zsh)

git clone https://github.com/zsh-users/zsh-autosuggestions ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-autosuggestions
git clone https://github.com/zsh-users/zsh-syntax-highlighting.git ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-syntax-highlighting

sed -i 's/plugins=(git)/plugins=(git autojump zsh-autosuggestions zsh-syntax-highlighting)/g' ~/.zshrc
sed -i 's/ZSH_THEME=".*$/ZSH_THEME="cloud"/g' ~/.zshrc


# >>> git >>>
# git config --global url.https://git-pd.megvii-inc.com/.insteadof ssh://git@git-pd.megvii-inc.com:
# git config --global url.https://git-pd.megvii-inc.com/.insteadof git@git-pd.megvii-inc.com:
# git config --global credential.https://git-pd.megvii-inc.com.helper store

git config --global core.editor vim
git config --global gui.encoding utf-8
git config --global i18n.commitencoding utf-8
git config --global i18n.logoutputencoding utf-8

# >>> locale >>>
echo "export LANG=C.UTF-8" >> ~/.zshrc
echo "export LC_ALL=C.UTF-8" >> ~/.zshrc
