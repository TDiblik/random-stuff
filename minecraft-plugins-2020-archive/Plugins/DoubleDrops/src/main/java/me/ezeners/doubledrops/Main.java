package me.ezeners.doubledrops;

import org.bukkit.plugin.java.JavaPlugin;

public final class Main extends JavaPlugin {

    @Override
    public void onEnable() {
        // Plugin startup logic
        System.out.println("Alright, Alright, I am starting up, you piece of shit!");
        System.out.println("You just have to wait!");

        getServer().getPluginManager().registerEvents(new BreakBlock(), this);
    }

    @Override
    public void onDisable() {
        // Plugin shutdown logic
    }
}
