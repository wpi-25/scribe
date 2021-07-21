import discord

from discord.ext import commands
import sys
import datetime
import logging

from . import config

logger = logging.getLogger("discord")
logger.setLevel(logger.warn)
stdoutHandler = logging.StreamHandler(stream=sys.stdout)
stdoutHandler.setFormatter(
    logging.Formatter("%(asctime)s:%(levelname)s:%(name)s: %(message)s")
)
logger.addHandler(stdoutHandler)

botLogger = logging.getLogger("scribe")
botLogger.setLevel(logger.info)
logger.addHandler(stdoutHandler)

class Scribe(commands.Bot):
    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)

        self.config = config

        self.startTime = datetime.datetime.now()

        self.logger = botLogger

        self.embed_colour = discord.Colour.dark_red()

# Enable members intent
intents = discord.Intents.default()
intents.members = True

bot = Scribe(command_prefix=config.discord_prefix, intents=intents)

@bot.event
async def on_ready():
    bot.logger.info(f"Logged in as {bot.user}")
    game = discord.Game(f"{bot.command_prefix}help")
    await bot.change_presence(status=discord.Status.online, activity=game)

def run_bot():
    bot.run(config.discord_token)