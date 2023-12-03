alias pacman-install="sudo pacman -S --needed"
alias paru-install="paru -S --needed"
alias pacman-help="echo \" \

# Pacman
System upgrade: sudo pacman -Syu
System upgrade docs (if something goes wrong): https://wiki.archlinux.org/title/System_maintenance#Upgrading_the_system

Search for package: sudo pacman -Ss package_name_1 package_name_2

Install package: sudo pacman -S --needed package_name_1 package_name_2

Install package (custom alias): pacman-install package_name_1 package_name_2

Remove package: sudo pacman -Rs package_name


# Paru
Aur packages upgrade: paru -Syua

Search for package: paru -Ss package_name_1 package_name_2

Install package: paru -S --needed package_name_1 package_name_2

Install package (custom alias): paru-install package_name_1 package_name_2

Remove package: paru -Rs package_name
\""