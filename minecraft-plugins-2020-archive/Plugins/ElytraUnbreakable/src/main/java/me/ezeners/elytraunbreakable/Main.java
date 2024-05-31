package me.ezeners.elytraunbreakable;

import org.bukkit.plugin.java.JavaPlugin;

public final class Main extends JavaPlugin {

    @Override
    public void onEnable() {
        // Plugin startup logic
        System.out.println("Starting elytra unbreakable mode");

        getServer().getPluginManager().registerEvents(new ElytraListener(), this);
    }

    @Override
    public void onDisable() {
        // Plugin shutdown logic
    }
}
