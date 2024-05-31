package me.ezeners.breakblockteleport;

import com.sun.org.apache.xpath.internal.operations.Bool;
import org.bukkit.Location;
import org.bukkit.Material;
import org.bukkit.block.Block;
import org.bukkit.entity.Player;
import org.bukkit.event.EventHandler;
import org.bukkit.event.Listener;
import org.bukkit.event.block.BlockBreakEvent;

import java.util.Random;

public class BreakBlockListener implements Listener {
    @EventHandler
    public void blockBreak(BlockBreakEvent event) {
        Player player = event.getPlayer();
        Location b_location = event.getBlock().getLocation();
        int b_x = b_location.getBlockX() - 100;
        int b_y = b_location.getBlockY() - 10;
        int b_z = b_location.getBlockZ() - 100;
        int b_x1 = b_location.getBlockX() + 100;
        int b_y1 = b_location.getBlockY() + 10;
        int b_z1 = b_location.getBlockZ() + 100;

        int random_x = 0;
        int random_y = 0;
        int random_z = 0;

        Location port_location = new Location(event.getBlock().getWorld(), random_x, random_y, random_z);

        Random random = new Random();

        int runs = 0;
        while (true) {
            random_x = random.nextInt(b_x1 - b_x + 1) + b_x;
            random_y = random.nextInt(b_y1 - b_y + 1) + b_y;
            random_z = random.nextInt(b_z1 - b_z + 1) + b_z;
            port_location = new Location(event.getBlock().getWorld(), random_x, random_y, random_z);
            Block b = port_location.getBlock();
            if (b.getType() == Material.AIR) {
                break;
            }
            runs++;
            if (runs > 50) {
                b_y1 += 5;
            }
        }

        player.teleport(port_location);
    }
}
