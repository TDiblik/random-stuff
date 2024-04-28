alias pacman-install="sudo pacman -S --refresh --needed"
alias paru-install="paru -S --refresh --needed"
alias pacman-help="echo \" \
# General
System upgrade:
  - sudo pacman -Syy && sudo pacman -Syu && paru -Syy && paru -Syua && flatpak update
  OR
  1. sudo pacman -Syy && sudo pacman -Syu
  2. paru -Syy && paru -Syua
  3. flatpak update
  - If something goes wrong: https://wiki.archlinux.org/title/System_maintenance#Upgrading_the_system

# Pacman
Search for package: sudo pacman -Ss package_name_1 package_name_2
Install package (aliased as pacman-install): sudo pacman -S --needed package_name_1 package_name_2
Remove package: sudo pacman -Rs package_name_1 package_name_2

# Paru
Search for package: paru -Ss package_name_1 package_name_2
Install package (aliased as paru-install): paru -S --needed package_name_1 package_name_2
Remove package: paru -Rs package_name_1 package_name_2
\""