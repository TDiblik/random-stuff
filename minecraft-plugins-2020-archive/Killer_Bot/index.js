const mineflayer = require('mineflayer');
const autoeat = require('mineflayer-auto-eat');
const armorManager = require('mineflayer-armor-manager');
const { pathfinder, Movements } = require('mineflayer-pathfinder');
const { GoalNear, GoalBlock, GoalXZ, GoalY, GoalInvert, GoalFollow } = require('mineflayer-pathfinder').goals;

const options = {
    host: 'gotcha.mc.hostify.cz',
    username: 'Killer_3000'
}

const bot = mineflayer.createBot(options);

bot.loadPlugin(pathfinder);
bot.loadPlugin(autoeat);
bot.loadPlugin(armorManager);

bot.once('spawn', () => {
    
	bot.autoEat.options = {
		priority: 'foodPoints',
		startAt: 14,
		bannedFood: ['rotten_flesh', 'spider_eye']
    };
    bot.autoEat.enable();
    
    const mcData = require('minecraft-data')(bot.version);

    const defaultMove = new Movements(bot, mcData);

    /*bot.on('path_update', (r) => {
        const nodesPerTick = (r.visitedNodes * 50 / r.time).toFixed(2)
        console.log(`I can get there in ${r.path.length} moves. Computation took ${r.time.toFixed(2)} ms (${nodesPerTick} nodes/tick).`)
    });
    
    bot.on('goal_reached', (goal) => {
        console.log('Here I am !')
    });*/
    bot.on('chat', (username, message) => {
        if (username === bot.username) return
    
        const target = bot.players[username] ? bot.players[username].entity : null

        if (message === 'killer-bot-start') {
            bot.chat("I am going to kill you! MUHAHAHAHAHA");
            setInterval(() => {
                const playerFilter = (entity) => entity.type == 'player';
                const playerEntity = bot.nearestEntity(playerFilter);
                if (playerEntity == null) {
                    return;
                }
                bot.pathfinder.setMovements(new Movements(bot, mcData));
                bot.pathfinder.setGoal(new GoalNear(playerEntity.position.x, playerEntity.position.y, playerEntity.position.z, 0), true);
            }, 1000)
        }
        else if (message.startsWith('killer-bot-goto')) {
          const cmd = message.split(' ')
    
          if (cmd.length === 4) { // goto x y z
            const x = parseInt(cmd[1], 10)
            const y = parseInt(cmd[2], 10)
            const z = parseInt(cmd[3], 10)
    
            bot.pathfinder.setMovements(defaultMove)
            bot.pathfinder.setGoal(new GoalBlock(x, y, z))
          } 
          else if (cmd.length === 3) { // goto x z
            const x = parseInt(cmd[1], 10)
            const z = parseInt(cmd[2], 10)
    
            bot.pathfinder.setMovements(defaultMove)
            bot.pathfinder.setGoal(new GoalXZ(x, z))
          } 
          else if (cmd.length === 2) { // goto y
            const y = parseInt(cmd[1], 10)
    
            bot.pathfinder.setMovements(defaultMove)
            bot.pathfinder.setGoal(new GoalY(y))
          }
        } 
        else if (message === 'killer-bot follow me') {
          bot.pathfinder.setMovements(defaultMove)
          bot.pathfinder.setGoal(new GoalFollow(target, 5), true)
        } 
        else if (message === 'killer-bot avoid me') {
          bot.pathfinder.setMovements(defaultMove)
          bot.pathfinder.setGoal(new GoalInvert(new GoalFollow(target, 100)), true)
        } 
        else if (message === 'killer-bot stop') {
          bot.pathfinder.setGoal(null)
        }
    });
});

bot.once('login', () => {
    bot.chat('Murderer loged in.');
    //1.10 a vyšší pvp
    /*setInterval(() => {
        const playerFilter = (entity) => entity.type == 'player';
        const playerEntity = bot.nearestEntity(playerFilter);
        if (playerEntity == null || bot.entity.position == null) {
            return;
        }
        var x_conf = (playerEntity.position.x - bot.entity.position.x <= 5 && playerEntity.position.x - bot.entity.position.x >= 0 || playerEntity.position.x + bot.entity.position.x <= 5 && playerEntity.position.x + bot.entity.position.x >= 0 || playerEntity.position.x - bot.entity.position.x >= -5 && playerEntity.position.x - bot.entity.position.x <= 0 || playerEntity.position.x + bot.entity.position.x >= -5 && playerEntity.position.x + bot.entity.position.x <= 0);
        var y_conf = (playerEntity.position.y - bot.entity.position.y <= 5 && playerEntity.position.y - bot.entity.position.y >= 0 || playerEntity.position.y + bot.entity.position.y <= 5 && playerEntity.position.y + bot.entity.position.y >= 0 || playerEntity.position.y - bot.entity.position.y >= -5 && playerEntity.position.y - bot.entity.position.y <= 0 || playerEntity.position.y + bot.entity.position.y >= -5 && playerEntity.position.y + bot.entity.position.y <= 0);
        var z_conf = (playerEntity.position.z - bot.entity.position.z <= 5 && playerEntity.position.z - bot.entity.position.z >= 0 || playerEntity.position.z + bot.entity.position.z <= 5 && playerEntity.position.z + bot.entity.position.z >= 0 || playerEntity.position.z - bot.entity.position.z >= -5 && playerEntity.position.z - bot.entity.position.z <= 0 || playerEntity.position.z + bot.entity.position.z >= -5 && playerEntity.position.z + bot.entity.position.z <= 0);
        if (x_conf && z_conf || x_conf && y_conf && z_conf) {
            bot.attack(playerEntity);
        }
    }, 590) // 600 -- OK, 590 -- hranice, ale OK*/
});


//!!!!!!!!!!!!!NEMAZAZT!!!!!!!!!!!!!
bot.on('physicTick', () => {
    //1.8 PVP
    const playerFilter = (entity) => entity.type == 'player';
    const playerEntity = bot.nearestEntity(playerFilter);
    if (playerEntity == null || bot.entity.position == null) {
        return;
    }
    //target.position.distanceTo(this.bot.entity.position);
    var x_conf = (playerEntity.position.x - bot.entity.position.x <= 5 && playerEntity.position.x - bot.entity.position.x >= 0 || playerEntity.position.x + bot.entity.position.x <= 5 && playerEntity.position.x + bot.entity.position.x >= 0 || playerEntity.position.x - bot.entity.position.x >= -5 && playerEntity.position.x - bot.entity.position.x <= 0 || playerEntity.position.x + bot.entity.position.x >= -5 && playerEntity.position.x + bot.entity.position.x <= 0);
    var y_conf = (playerEntity.position.y - bot.entity.position.y <= 5 && playerEntity.position.y - bot.entity.position.y >= 0 || playerEntity.position.y + bot.entity.position.y <= 5 && playerEntity.position.y + bot.entity.position.y >= 0 || playerEntity.position.y - bot.entity.position.y >= -5 && playerEntity.position.y - bot.entity.position.y <= 0 || playerEntity.position.y + bot.entity.position.y >= -5 && playerEntity.position.y + bot.entity.position.y <= 0);
    var z_conf = (playerEntity.position.z - bot.entity.position.z <= 5 && playerEntity.position.z - bot.entity.position.z >= 0 || playerEntity.position.z + bot.entity.position.z <= 5 && playerEntity.position.z + bot.entity.position.z >= 0 || playerEntity.position.z - bot.entity.position.z >= -5 && playerEntity.position.z - bot.entity.position.z <= 0 || playerEntity.position.z + bot.entity.position.z >= -5 && playerEntity.position.z + bot.entity.position.z <= 0);
    if (x_conf && z_conf || x_conf && y_conf && z_conf) {
        bot.attack(playerEntity);
    }
})

bot.on('respawn', () => {
    bot.chat('You were lucky once, but you wont be lucky twice!');
});

bot.on('death', () => {
    bot.chat('You were lucky this time!');
    bot.pathfinder.setMovements(null);
    bot.pathfinder.setGoal(null);
});

//preventivní
bot.on('kicked', (reason, loggedIn) => {
    console.log(reason, loggedIn);
});
bot.on('error', err => console.log(err));

//NEODEBÍRAT -- MUSÍ TADY BÝT, ABY AUTOEAT FUNGOVAL SPRÁVNĚ
bot.on('autoeat_started', () => {
	console.log('Auto Eat started!');
});

bot.on('autoeat_stopped', () => {
	console.log('Auto Eat stopped!');
});