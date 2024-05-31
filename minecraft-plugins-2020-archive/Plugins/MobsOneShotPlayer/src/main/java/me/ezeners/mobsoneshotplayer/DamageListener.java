package me.ezeners.mobsoneshotplayer;

import org.bukkit.Bukkit;
import org.bukkit.entity.*;
import org.bukkit.event.EventHandler;
import org.bukkit.event.Listener;
import org.bukkit.event.entity.EntityDamageByEntityEvent;
import org.bukkit.event.entity.EntityDamageEvent;

import java.util.Map;

public class DamageListener implements Listener {
    @EventHandler
    public void onHit(EntityDamageByEntityEvent event)
    {
        if (!(event.getEntity() instanceof Player)) return;
        Boolean pass = (event.getDamager() instanceof Arrow || event.getDamager() instanceof Fireball);
        if (!(event.getDamager() instanceof Mob) && !pass) return;

        Player player = (Player) event.getEntity();

        if (event.getDamager() instanceof Arrow) {
            Arrow arrow = (Arrow) event.getDamager();
            Entity entity = (Entity) arrow.getShooter();
            if (!(entity instanceof Player)) {
                player.damage(50000, entity);
            }
            event.setCancelled(true);
            return;
        }


        if (event.getDamager() instanceof Fireball) {
            Fireball fireball = (Fireball) event.getDamager();
            Entity entity = (Entity) fireball.getShooter();
            if (!(entity instanceof Player)) {
                player.damage(50000, entity);
            }
            event.setCancelled(true);
            return;
        }

        Monster monster = (Monster) event.getDamager();
        player.damage(50000, monster);
        event.setCancelled(true);

        /*System.out.println("kontrola 1");
        if (!(event.getDamager() instanceof Monster)) return;
        System.out.println("kontrola 2");
        if (!(event.getEntity() instanceof Player)) return;
        System.out.println("kontrola 3");
        Monster monster = (Monster) event.getDamager();
        Player player = (Player) event.getEntity();

        player.damage(1000, monster);
            event.setCancelled(true);
        System.out.println("zavolal event");*/
    }
}
