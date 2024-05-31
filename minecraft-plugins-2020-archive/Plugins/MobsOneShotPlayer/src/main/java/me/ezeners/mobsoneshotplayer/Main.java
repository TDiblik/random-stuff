package me.ezeners.mobsoneshotplayer;

import org.bukkit.plugin.java.JavaPlugin;

public final class Main extends JavaPlugin {

    @Override
    public void onEnable() {
        // Plugin startup logic
        System.out.println("I am starting up plugin which makes that every mobs one shot player");

        getServer().getPluginManager().registerEvents(new DamageListener(), this);
    }

    @Override
    public void onDisable() {
        // Plugin shutdown logic
    }
}
