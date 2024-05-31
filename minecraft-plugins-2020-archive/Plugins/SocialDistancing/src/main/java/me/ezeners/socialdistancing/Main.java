package me.ezeners.socialdistancing;

import org.bukkit.Bukkit;
import org.bukkit.entity.Entity;
import org.bukkit.entity.Player;
import org.bukkit.entity.Zombie;
import org.bukkit.plugin.java.JavaPlugin;
import org.bukkit.scheduler.BukkitRunnable;

public final class Main extends JavaPlugin {

    @Override
    public void onEnable() {
        // Plugin startup logic
        System.out.println("Plugin for social distancing is starting.");

        new BukkitRunnable() {

            @Override
            public void run() {
                //get all player that have items and steal from them
                for(Player p : Bukkit.getOnlinePlayers()){
                    for (Entity e : p.getNearbyEntities(4, 4, 4)) {
                        if (e instanceof Player) {
                            p.damage(10, e);
                            ((Player) e).damage(10, p);
                        }
                        /*if (e instanceof Zombie) {
                            p.damage(1, e);
                        }*/
                    }
                }
            }

        }.runTaskTimer(this, 0L, 20);
    }

    @Override
    public void onDisable() {
        // Plugin shutdown logic
    }
}
