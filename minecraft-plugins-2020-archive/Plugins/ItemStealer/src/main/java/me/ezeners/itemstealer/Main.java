package me.ezeners.itemstealer;

import org.bukkit.Bukkit;
import org.bukkit.entity.Player;
import org.bukkit.inventory.ItemStack;
import org.bukkit.plugin.java.JavaPlugin;
import org.bukkit.scheduler.BukkitRunnable;

import java.util.ArrayList;
import java.util.List;
import java.util.Random;

public final class Main extends JavaPlugin {

    @Override
    public void onEnable() {
        // Plugin startup logic
        System.out.println("Plugin for stealing items starting.");

        new BukkitRunnable() {

            @Override
            public void run() {
                System.out.println("Stealing items.");
                //get all player that have items and steal from them
                for(Player p : Bukkit.getOnlinePlayers()){
                    if (p.getInventory().isEmpty()) {
                        System.out.println("Inventory is empty.");
                        continue;
                    }
                    Integer index = new Random().nextInt(p.getInventory().getContents().length);
                    ItemStack item = p.getInventory().getContents()[index];
                    while (item == null) {
                        index = new Random().nextInt(p.getInventory().getContents().length);
                        item = p.getInventory().getContents()[index];
                    }
                    p.getInventory().remove(item);
                }
            }

        }.runTaskTimer(this, 0L, 20*20);
    }

    @Override
    public void onDisable() {
        // Plugin shutdown logic
    }
}
