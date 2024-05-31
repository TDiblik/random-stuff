package me.ezeners.spidermanplugin;

import org.bukkit.plugin.java.JavaPlugin;

public final class Main extends JavaPlugin {

    @Override
    public void onEnable() {
        // Plugin startup logic
        System.out.println("I am starting spiderman plugin");

        getServer().getPluginManager().registerEvents(new SpiderShooter(), this);
    }

    @Override
    public void onDisable() {
        // Plugin shutdown logic
    }
}
