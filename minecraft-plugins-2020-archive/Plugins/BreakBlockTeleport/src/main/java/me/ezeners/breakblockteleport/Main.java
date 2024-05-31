package me.ezeners.breakblockteleport;

import org.bukkit.plugin.java.JavaPlugin;

public final class Main extends JavaPlugin {

    @Override
    public void onEnable() {
        // Plugin startup logic
        System.out.println("Starting plugin that teleports you when you break block");

        getServer().getPluginManager().registerEvents(new BreakBlockListener(), this);
    }

    @Override
    public void onDisable() {
        // Plugin shutdown logic
    }
}
