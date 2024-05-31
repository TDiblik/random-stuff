package me.ezeners.spidermanplugin;

import org.bukkit.Material;
import org.bukkit.entity.Fireball;
import org.bukkit.entity.Player;
import org.bukkit.event.EventHandler;
import org.bukkit.event.Listener;
import org.bukkit.event.entity.EntityDamageEvent;
import org.bukkit.event.player.PlayerInteractEvent;

public class SpiderShooter implements Listener {
    @EventHandler
    public void onPlayerUse(PlayerInteractEvent event){
        Player p = event.getPlayer();

        if(p.getItemInHand().getType() == Material.BLAZE_ROD){
            Fireball fire = p.getWorld().spawn(event.getPlayer().getLocation().add(0, 2, 0), Fireball.class);
            fire.setShooter(p);
        }
        else if (p.getItemInHand().getType() == Material.STICK) {
            p.setVelocity(p.getLocation().getDirection().multiply(100));
        }
    }

    @EventHandler
    public void fallDamageEvent(EntityDamageEvent event){
        if(event.getEntity() instanceof Player && event.getCause() == EntityDamageEvent.DamageCause.FALL) {
            event.setDamage(event.getDamage() / 8);
        }
    }
}
