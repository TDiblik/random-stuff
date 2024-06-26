package me.ezeners.doubledrops;

import org.bukkit.Material;
import org.bukkit.block.Block;
import org.bukkit.event.EventHandler;
import org.bukkit.event.Listener;
import org.bukkit.event.block.BlockBreakEvent;
import org.bukkit.inventory.ItemStack;

import java.util.Random;

public class BreakBlock implements Listener {
    @EventHandler
    public void onPlayerBreakBlock(BlockBreakEvent event) {
        Block blockBroken = event.getBlock();

        if (blockBroken.getType() == Material.DIAMOND_ORE) {
            event.setCancelled(true);
            blockBroken.setType(Material.AIR);
            ItemStack diamonds = new ItemStack(Material.DIAMOND, new Random().nextInt(100000));
            blockBroken.getWorld().dropItemNaturally(blockBroken.getLocation(), diamonds);
        }
    }
}