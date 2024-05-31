package me.ezeners.elytraunbreakable;

import org.bukkit.Material;
import org.bukkit.event.EventHandler;
import org.bukkit.event.Listener;
import org.bukkit.event.inventory.InventoryMoveItemEvent;
import org.bukkit.event.inventory.InventoryType;
import org.bukkit.event.player.PlayerPickupItemEvent;
import org.bukkit.inventory.meta.ItemMeta;

public class ElytraListener implements Listener {
    @EventHandler
    public void PickupItem(PlayerPickupItemEvent e) {
        if (e.getItem().getItemStack().getType() == Material.ELYTRA) {
            ItemMeta meta = e.getItem().getItemStack().getItemMeta();
            meta.setUnbreakable(true);
            e.getItem().getItemStack().setItemMeta(meta);
        };
    }
}
